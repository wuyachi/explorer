package chain

import (
	"context"
	"encoding/hex"
	"explorer/conf"
	"explorer/models"
	"explorer/utils"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/native"
	"github.com/palettechain/palette-go-sdk"
	"github.com/palettechain/palette-go-sdk/client"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/big"
	"strings"
	"time"
)

type Chain struct {
	cfg   *conf.Config
	sdk   *palette_go_sdk.PaletteSdk
	db    *gorm.DB
	chain *models.Chain
}

func NewChain(cfg *conf.Config) *Chain {
	chain := &Chain{}
	chain.cfg = cfg
	sdk := palette_go_sdk.NewPaletteSdk(cfg.NodeConfig.RestURL)
	accounts := sdk.ListAccounts()
	if len(accounts) == 0 {
		sdk.NewAccount()
	}
	chain.sdk = sdk
	//db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/palette?charset=utf8"), &gorm.Config{})
	db, err := gorm.Open(mysql.Open(cfg.DBConfig.User + ":" +
		cfg.DBConfig.Password + "@tcp(" + cfg.DBConfig.URL + ")/" +
		cfg.DBConfig.Scheme + "?charset=utf8"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&models.Chain{}, &models.Block{}, &models.Transaction{},&models.Event{},
		&models.PLTContract{}, &models.NFTContract{}, &models.Validator{}, &models.Propose{},
		&models.Stake{}, &models.TransactionDetail{})
	if err != nil {
		panic(err)
	}
	chain.db = db
	//
	paletteChain := new(models.Chain)
	paletteChain.Id = 1
	paletteChain.Name = "palette"
	db.Create(paletteChain)
	adminAccount := new(models.PLTContract)
	adminAccount.Address = strings.ToLower(common.HexToAddress(cfg.NodeConfig.Admin).String())
	adminAccount.Amount = 600000000000000000
	db.Create(adminAccount)
	return chain
}

func (this *Chain) Start() {
	this.chain = new(models.Chain)
	result := this.db.First(this.chain)
	if result.Error != nil {
		panic(result.Error)
	}
	if result.RowsAffected != 1 {
		panic("There is no chains!")
	}
	go this.ListenChain()
}

func (this *Chain) ListenChain() {
	logs.Debug("listen chain......")
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			height, err := this.sdk.BlockNumber()
			if err != nil || height.Uint64() == 0 {
				logs.Error("ListenChain - cannot get height, err: %s", err)
				continue
			}
			if this.chain.Height >= height.Uint64() - 1 {
				continue
			}
			logs.Info("ListenChain - chain latest height is %d, parser height: %d", height.Uint64(), this.chain.Height)
			for this.chain.Height < height.Uint64() - 1 {
				err := this.HandleNewBlock(this.chain.Height + 1)
				if err != nil {
					logs.Error("HandleNewBlock err: %v", err)
					break
				}
			}
		}
	}
}

func (this *Chain) HandleNewBlock(height uint64) error {
	//
	pltContractMap := make(map[string]*models.PLTContract, 0)
	nftContractMap := make(map[string]*models.NFTContract, 0)
	proposeMap := make(map[string]*models.Propose, 0)
	stakeMap := make(map[string]*models.Stake, 0)
	// parse block
	block, err := this.sdk.BlockByNumber(context.Background(), big.NewInt(int64(height)))
	if err != nil {
		return err
	}
	blockInfo := new(models.Block)
	blockInfo.Hash = block.Root().String()
	blockInfo.GasLimit = block.GasLimit()
	blockInfo.GasUsed = block.GasUsed()
	blockInfo.Difficulty = block.Difficulty().Uint64()
	blockInfo.Number = block.NumberU64()
	blockInfo.ParentHash = block.ParentHash().String()
	blockInfo.ReceiptHash = block.ReceiptHash().String()
	blockInfo.TxHash = block.TxHash().String()
	blockInfo.TxNumber = uint64(block.Transactions().Len())
	blockInfo.Time = block.Time()
	// parse all transactions
	blockInfo.Transactions = make([]*models.Transaction, 0)
	for _, transaction := range block.Transactions() {
		transactionInfo := new(models.Transaction)
		transactionInfo.Hash = transaction.Hash().String()
		transactionInfo.From = strings.ToLower(transaction.From().String())
		transactionInfo.Cost = transaction.Cost().Uint64()
		data := hex.EncodeToString(transaction.Data())
		dataLen := len(data)
		if dataLen > 4096 {
			dataLen = 4096
		}
		transactionInfo.Data = data[0:dataLen]
		transactionInfo.Gas = transaction.Gas()
		transactionInfo.GasPrice = utils.AbandonPrecision(transaction.GasPrice())
		if transaction.To() == nil {
			transactionInfo.To = ""
		} else {
			transactionInfo.To = strings.ToLower(transaction.To().String())
		}
		transactionInfo.Value = utils.AbandonPrecision(transaction.Value())
		transactionInfo.BlockNumber = blockInfo.Number
		transactionInfo.Time = blockInfo.Time
		blockInfo.Transactions = append(blockInfo.Transactions, transactionInfo)
		// parse all events
		events, err := this.sdk.GetEventLog(transaction.Hash())
		if err != nil {
			return err
		}
		for _, event :=  range events {
			eventInfo := new(models.Event)
			eventInfo.Number = event.BlockNumber
			eventInfo.Contract = strings.ToLower(event.Contract.String())
			eventInfo.EventId = event.EventId.String()
			if len(event.Topic) > 0 {
				eventInfo.Topic1 = event.Topic[0].String()
			}
			if len(event.Topic) > 1 {
				eventInfo.Topic2 = event.Topic[1].String()
			}
			if len(event.Topic) > 2 {
				eventInfo.Topic3 = event.Topic[2].String()
			}
			if len(event.Topic) > 3 {
				eventInfo.Topic4 = event.Topic[3].String()
			}
			data := hex.EncodeToString(event.Data)
			dataLen := len(data)
			if dataLen > 4096 {
				dataLen = 4096
			}
			eventInfo.Data = data[0:dataLen]
			transactionInfo.Events = append(transactionInfo.Events, eventInfo)
			//
			if event.Contract == common.HexToAddress(native.PLTContractAddress) {
				if event.EventId.String() == client.PLTEventID_Transfer {
					from := strings.ToLower(common.BytesToAddress(event.Topic[0].Bytes()).String())
					to := strings.ToLower(common.BytesToAddress(event.Topic[1].Bytes()).String())
					amount := utils.AbandonPrecision(new(big.Int).SetBytes(event.Data))

					txDetailInfo := new(models.TransactionDetail)
					txDetailInfo.Contract = strings.ToLower(event.Contract.String())
					txDetailInfo.From = from
					txDetailInfo.To = to
					txDetailInfo.Value = fmt.Sprintf("%d", amount)
					transactionInfo.TransactionDetails = append(transactionInfo.TransactionDetails, txDetailInfo)

					var fromUser *models.PLTContract
					fromUser, ok := pltContractMap[from]
					if !ok {
						fromUser = new(models.PLTContract)
						this.db.Where("address = ?", from).First(fromUser)
						fromUser.Address = from
						pltContractMap[from] = fromUser
					}

					var toUser *models.PLTContract
					toUser, ok = pltContractMap[to]
					if !ok {
						toUser = new(models.PLTContract)
						this.db.Where("address = ?", to).First(toUser)
						toUser.Address = to
						pltContractMap[to] = toUser
					}

					if fromUser.Amount < amount {
						logs.Error("from : %s amount is %d, but transfer amount is %d", fromUser.Address, fromUser.Amount, amount)
					}

					fromUser.Amount = fromUser.Amount - amount
					toUser.Amount = toUser.Amount + amount
				}
				continue
			}
			addr := new(big.Int).SetBytes(event.Contract.Bytes())
			if addr.Cmp(native.NFTContractAddressStart) >= 0 && addr.Cmp(native.NativeContractAddressEnd) < 0 {
				if event.EventId.String() == client.NFTEventID_Transfer {
					from := strings.ToLower(common.BytesToAddress(event.Topic[0].Bytes()).String())
					nft := strings.ToLower(event.Contract.String())
					to := strings.ToLower(common.BytesToAddress(event.Topic[1].Bytes()).String())
					tokenId := new(big.Int).SetBytes(event.Data)
					token := common.BigToHash(tokenId).String()

					txDetailInfo := new(models.TransactionDetail)
					txDetailInfo.Contract = strings.ToLower(event.Contract.String())
					txDetailInfo.From = from
					txDetailInfo.To = to
					txDetailInfo.Value = token
					transactionInfo.TransactionDetails = append(transactionInfo.TransactionDetails, txDetailInfo)

					var tokenInfo *models.NFTContract
					tokenInfo, ok := nftContractMap[nft + token]
					if !ok {
						tokenInfo = new(models.NFTContract)
						this.db.Where("nft = ? and token = ?", nft, token).First(tokenInfo)
						tokenInfo.NFT = nft
						tokenInfo.Token = token
						nftContractMap[nft + token] = tokenInfo
					}
					tokenInfo.Owner = to
					tokenInfo.Uri, _ = this.sdk.NFTTokenUri(event.Contract, tokenId)
				}
				continue
			}
			if event.Contract == common.HexToAddress(native.GovernanceContractAddress) {
				if event.EventId.String() == client.GovernanceEventID_Propose {
					proposer := strings.ToLower(common.BytesToAddress(event.Topic[0].Bytes()).String())
					proposeId := strings.ToLower(common.BytesToAddress(event.Topic[1].Bytes()).String())

					var proposeInfo *models.Propose
					proposeInfo, ok := proposeMap[proposeId]
					if !ok {
						proposeInfo = new(models.Propose)
						this.db.Where("propose_id = ?", proposeId).First(proposeInfo)
						proposeInfo.ProposeId = proposeId
						proposeMap[proposeId] = proposeInfo
					}
					proposeInfo.Proposer = proposer
					propose, _ := this.sdk.GovernanceGetProposal(common.BytesToAddress(event.Topic[1].Bytes()))
					proposeInfo.Value = propose.Value.String()
					proposeInfo.EndBlock = propose.EndBlock.Uint64()
					proposeInfo.Type = propose.ProposalType
					proposeInfo.State = 0
				} else if event.EventId.String() == client.GovernanceEventID_Vote {
					proposeId := strings.ToLower(common.BytesToAddress(event.Topic[1].Bytes()).String())
					var proposeInfo *models.Propose
					proposeInfo, ok := proposeMap[proposeId]
					if !ok {
						proposeInfo = new(models.Propose)
						this.db.Where("propose_id = ?", proposeId).First(proposeInfo)
						proposeInfo.ProposeId = proposeId
						proposeMap[proposeId] = proposeInfo
					}
					propose, _ := this.sdk.GovernanceGetProposal(common.BytesToAddress(event.Topic[1].Bytes()))
					proposeInfo.State = 0
					if propose.Passed == true {
						proposeInfo.State = 1
					} else if height > propose.EndBlock.Uint64() {
						proposeInfo.State = 2
					}
				} else if event.EventId.String() == client.GovernanceEventID_Stake {
					validator := strings.ToLower(common.BytesToAddress(event.Topic[1].Bytes()).String())
					owner := strings.ToLower(common.BytesToAddress(event.Topic[2].Bytes()).String())
					revoke := utils.Hash2Bool(event.Topic[3])
					amount := utils.AbandonPrecision(new(big.Int).SetBytes(event.Data))

					var stakeInfo *models.Stake
					stakeInfo, ok := stakeMap[owner + validator]
					if !ok {
						stakeInfo = new(models.Stake)
						this.db.Where("owner = ? and validator = ?", owner, validator).First(stakeInfo)
						stakeInfo.Owner = owner
						stakeInfo.Validator = validator
						stakeMap[owner + validator] = stakeInfo
					}

					if revoke == false {
						stakeInfo.StakeAmount = stakeInfo.StakeAmount + amount
					} else {
						if stakeInfo.StakeAmount < amount {
							logs.Error("unstake : %s amount is %d, but stake amount is %d", owner, amount, stakeInfo.StakeAmount)
						}
						stakeInfo.StakeAmount = stakeInfo.StakeAmount - amount
					}
				}
			}
		}
	}

	pltContractInfos := make([]*models.PLTContract, 0)
	for _, pltContract := range pltContractMap {
		pltContractInfos = append(pltContractInfos, pltContract)
	}

	nftContractInfos := make([]*models.NFTContract, 0)
	for _, nftContract := range nftContractMap {
		nftContractInfos = append(nftContractInfos, nftContract)
	}

	stakeInfos := make([]*models.Stake, 0)
	for _, stake := range stakeMap {
		stakeInfos = append(stakeInfos, stake)
	}

	proposeInfos := make([]*models.Propose, 0)
	for _, propose := range proposeMap {
		proposeInfos = append(proposeInfos, propose)
	}

	validatorInfos := make([]*models.Validator, 0)
	validators, _ := this.sdk.GovernanceGetValidators()
	for _, validator := range validators {
		validatorInfo := new(models.Validator)
		validatorInfo.Address = validator.String()
		factor, _ := this.sdk.GovernanceGetDelegateFactor(validator)
		amount, _ := this.sdk.GovernanceGetValidatorTotalStakeAmount(validator)
		validatorInfo.DelegateFactor = factor.Uint64()
		validatorInfo.StakeAmount = utils.AbandonPrecision(amount)
		validatorInfos = append(validatorInfos, validatorInfo)
	}

	stakeAmount, _ := this.sdk.GovernanceGetAllStakeAmount()
	lastReward, _ := this.sdk.GovernanceGetLastRewardBlockHeight()
	mintPrice, _ := this.sdk.GetMintPrice()
	rewardPeriod, _ := this.sdk.GetRewardPeriod()
	gasFee, _ := this.sdk.GetGasFee()
	this.chain.StakeAmount = utils.AbandonPrecision(stakeAmount)
	this.chain.LastReward = lastReward.Uint64()
	this.chain.MintPrice = utils.AbandonPrecision(mintPrice)
	this.chain.RewardPeriod = rewardPeriod.Uint64()
	this.chain.GasFee = utils.AbandonPrecision(gasFee)
	this.chain.Height = height
	// update db
	this.db.Create(blockInfo)
	for _, pltContract := range pltContractInfos {
		this.db.Save(pltContract)
	}
	for _, nftContract := range nftContractInfos {
		this.db.Save(nftContract)
	}
	for _, stake := range stakeInfos {
		this.db.Save(stake)
	}
	for _, propose := range proposeInfos {
		this.db.Save(propose)
	}
	this.db.Where("1 = 1").Delete(&models.Validator{})
	if len(validatorInfos) > 0 {
		this.db.Save(validatorInfos)
	}
	this.db.Save(this.chain)
	return nil
}
