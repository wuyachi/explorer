package models

import (
	"explorer/basedef"
	"explorer/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/native"
	"github.com/ethereum/go-ethereum/contracts/native/utils/decimal"
)

type ExplorerResp struct {
	Version string
	URL     string
}

type ChainResp struct {
	Id           uint64
	Name         string
	Height       uint64
	StakeAmount  string
	LastReward   uint64
	MintPrice    string
	RewardPeriod uint64
	GasFee       string
	TransactionNum uint64
	BlockTime    string
}

func MakeChainResponse(chain *Chain, transactionNum uint64, totalTime uint64, totalBlock uint64) *ChainResp {
	decimal.DivisionPrecision = 2
	times := decimal.NewFromInt(int64(totalTime))
	blocks := decimal.NewFromInt(int64(totalBlock))
	avgTime := times.Div(blocks)
	chainResp := &ChainResp{
		Id:           chain.Id,
		Name:         chain.Name,
		Height:       chain.Height,
		StakeAmount:  utils.AmountWithoutPrecision(chain.StakeAmount),
		LastReward:   chain.LastReward,
		MintPrice:    utils.AmountWithoutPrecision(chain.MintPrice),
		RewardPeriod: chain.RewardPeriod,
		GasFee:       utils.AmountWithoutPrecision(chain.GasFee),
		TransactionNum: transactionNum,
		BlockTime: avgTime.String(),
	}
	return chainResp
}

type BlockResp struct {
	Hash         string
	GasLimit     uint64
	Size         uint64
	Validators   uint64
	GasUsed      uint64
	Difficulty   uint64
	Number       uint64
	ParentHash   string
	ReceiptHash  string
	TxHash       string
	TxNumber     uint64
	Time         uint64
	Transactions []string
}

type BlockByNumberReq struct {
	Number uint64
}

type BlockByHashReq struct {
	Hash string
}

func MakeBlockResponse(block *Block) *BlockResp {
	blockOut := &BlockResp{
		Hash:        block.Hash,
		Size:        block.Size,
		Validators:  block.Validators,
		GasLimit:    block.GasLimit,
		GasUsed:     block.GasUsed,
		Difficulty:  block.Difficulty,
		Number:      block.Number,
		ParentHash:  block.ParentHash,
		ReceiptHash: block.ReceiptHash,
		TxHash:      block.TxHash,
		TxNumber:    block.TxNumber,
		Time:        block.Time,
	}
	for _, transaction := range block.Transactions {
		blockOut.Transactions = append(blockOut.Transactions, transaction.Hash)
	}
	return blockOut
}

type BlocksReq struct {
	PageSize int
	PageNo   int
}

type BlocksResp struct {
	PageSize   int
	PageNo     int
	TotalPage  int
	TotalCount int
	Blocks     []*BlockResp
}

func MakeBlocksResponse(pageSize int, pageNo int, totalPage int, totalCount int, blocks []*Block) *BlocksResp {
	blocksResp := &BlocksResp{
		PageSize:   pageSize,
		PageNo:     pageNo,
		TotalPage:  totalPage,
		TotalCount: totalCount,
	}
	for _, block := range blocks {
		blocksResp.Blocks = append(blocksResp.Blocks, MakeBlockResponse(block))
	}
	return blocksResp
}

type TransactionResp struct {
	Hash               string
	From               string
	Cost               uint64
	Data               string
	Gas                uint64
	GasPrice           string
	To                 string
	Value              string
	Time               uint64
	BlockNumber        uint64
	Type               uint64
	Status             uint64
	BlockHash          string
	Events             []*EventResp             `json:",omitempty"`
	TransactionDetails []*TransactionDetailResp `json:",omitempty"`
}

type TransactionByHashReq struct {
	Hash string
}

func MakeTransactionResponse(transaction *Transaction) *TransactionResp {
	transactionResp := &TransactionResp{
		Hash:        transaction.Hash,
		From:        transaction.From,
		Cost:        transaction.Cost,
		Data:        transaction.Data,
		Gas:         transaction.Gas,
		GasPrice:    utils.AmountWithoutPrecision(transaction.GasPrice),
		To:          transaction.To,
		Value:       transaction.Value.String(),
		BlockNumber: transaction.BlockNumber,
		Time:        transaction.Time,
		BlockHash:   transaction.BlockHash,
		Type:        transaction.Type,
		Status:      transaction.Status,
	}
	for _, event := range transaction.Events {
		transactionResp.Events = append(transactionResp.Events, MakeEventResponse(event))
	}
	for _, transactionDetail := range transaction.TransactionDetails {
		transactionResp.TransactionDetails = append(transactionResp.TransactionDetails, MakeTransactionDetailResponse1(transactionDetail))
	}
	return transactionResp
}

func MakeTransaction1Response(transaction *Transaction1) *TransactionResp {
	transactionResp := &TransactionResp{
		Hash:        transaction.Hash,
		From:        transaction.From,
		Cost:        transaction.Cost,
		Data:        transaction.Data,
		Gas:         transaction.Gas,
		GasPrice:    utils.AmountWithoutPrecision(transaction.GasPrice),
		To:          transaction.To,
		Value:       transaction.Value.String(),
		BlockNumber: transaction.BlockNumber,
		Time:        transaction.Time,
		BlockHash:   transaction.BlockHash,
		Type:        transaction.Type,
		Status:      transaction.Status,
	}
	for _, event := range transaction.Events {
		transactionResp.Events = append(transactionResp.Events, MakeEventResponse(event))
	}
	for _, transactionDetail := range transaction.TransactionDetails {
		transactionResp.TransactionDetails = append(transactionResp.TransactionDetails, MakeTransactionDetailResponse(transactionDetail))
	}
	return transactionResp
}

type TransactionsReq struct {
	PageSize int
	PageNo   int
}

type TransactionsOfContractReq struct {
	Contract string
	PageSize int
	PageNo   int
}

type TransactionsOfUserReq struct {
	User     string
	PageSize int
	PageNo   int
}

type TransactionsResp struct {
	PageSize     int
	PageNo       int
	TotalPage    int
	TotalCount   int
	Transactions []*TransactionResp
}

func MakeTransactionsResponse(pageSize int, pageNo int, totalPage int, totalCount int, transactions []*Transaction) *TransactionsResp {
	transactionsResp := &TransactionsResp{
		PageSize:   pageSize,
		PageNo:     pageNo,
		TotalPage:  totalPage,
		TotalCount: totalCount,
	}
	for _, transaction := range transactions {
		transactionsResp.Transactions = append(transactionsResp.Transactions, MakeTransactionResponse(transaction))
	}
	return transactionsResp
}

type EventResp struct {
	Number          uint64
	Contract        string
	EventId         string
	Topic1          string
	Topic2          string
	Topic3          string
	Topic4          string
	Data            string
	Time            uint64
	TransactionHash string
}

func MakeEventResponse(event *Event) *EventResp {
	eventResp := &EventResp{
		Number:          event.Number,
		Contract:        event.Contract,
		EventId:         event.EventId,
		Topic1:          event.Topic1,
		Topic2:          event.Topic2,
		Topic3:          event.Topic3,
		Topic4:          event.Topic4,
		Data:            event.Data,
		Time:            event.Time,
		TransactionHash: event.TransactionHash,
	}
	return eventResp
}

type TransactionDetailResp struct {
	Contract        string
	From            string
	To              string
	Value           string
	Time            uint64
	Status          uint64
	IsNft           bool
	TransactionHash string
	ContractInfo    *ContractInfoResp `json:",omitempty"`
	TokenInfo       *NFTHolderResp    `json:",omitempty"`
}

func MakeTransactionDetailResponse(transactionDetail *TransactionDetailWithInfo) *TransactionDetailResp {
	transactionDetailResp := &TransactionDetailResp{
		Contract:        transactionDetail.Contract,
		From:            transactionDetail.From,
		To:              transactionDetail.To,
		Value:           transactionDetail.Value.String(),
		Time:            transactionDetail.Time,
		Status:          basedef.TRANSACTION_STATUS_SUCCESS,
		TransactionHash: transactionDetail.TransactionHash,
	}

	thisAddr := common.HexToAddress(transactionDetailResp.Contract)
	expectAddr := common.HexToAddress(native.PLTContractAddress)
	if thisAddr == expectAddr {
		var pltAmount uint64
		fmt.Sscanf(transactionDetailResp.Value, "%d", &pltAmount)
		transactionDetailResp.Value = utils.AmountWithoutPrecision(pltAmount)
	} else {
		transactionDetailResp.IsNft = true
	}

	if transactionDetail.ContractInfo != nil {
		transactionDetailResp.ContractInfo = MakeContractInfoResponse(transactionDetail.ContractInfo)
	}
	if transactionDetail.NFTHolder != nil {
		transactionDetailResp.TokenInfo = MakeNFTHolderResponse(transactionDetail.NFTHolder)
		if transactionDetail.ContractInfo != nil {
			transactionDetailResp.TokenInfo.Uri = transactionDetail.ContractInfo.BaseUri + transactionDetailResp.TokenInfo.Uri
		}
	}
	return transactionDetailResp
}

func MakeTransactionDetailResponse1(transactionDetail *TransactionDetail) *TransactionDetailResp {
	transactionDetailResp := &TransactionDetailResp{
		Contract:        transactionDetail.Contract,
		From:            transactionDetail.From,
		To:              transactionDetail.To,
		Value:           transactionDetail.Value.String(),
		Time:            transactionDetail.Time,
		Status:          basedef.TRANSACTION_STATUS_SUCCESS,
		TransactionHash: transactionDetail.TransactionHash,
	}
	thisAddr := common.HexToAddress(transactionDetailResp.Contract)
	expectAddr := common.HexToAddress(native.PLTContractAddress)
	if thisAddr == expectAddr {
		var pltAmount uint64
		fmt.Sscanf(transactionDetailResp.Value, "%d", &pltAmount)
		transactionDetailResp.Value = utils.AmountWithoutPrecision(pltAmount)
	} else {
		transactionDetailResp.IsNft = true
	}
	return transactionDetailResp
}

type TransactionDetailsOfContractReq struct {
	Contract string
	PageSize int
	PageNo   int
}

type TransactionDetailsOfNFTTokenReq struct {
	Contract string
	Token    string
	PageSize int
	PageNo   int
}

type TransactionDetailsOfTransactionReq struct {
	Hash     string
	Token    string
	PageSize int
	PageNo   int
}

type TransactionDetailsResp struct {
	PageSize           int
	PageNo             int
	TotalPage          int
	TotalCount         int
	TransactionDetails []*TransactionDetailResp
}

func MakeTransactionDetailsResponse(pageSize int, pageNo int, totalPage int, totalCount int, transactionDetails []*TransactionDetailWithInfo) *TransactionDetailsResp {
	transactionDetailsResp := &TransactionDetailsResp{
		PageSize:   pageSize,
		PageNo:     pageNo,
		TotalPage:  totalPage,
		TotalCount: totalCount,
	}
	for _, transactionDetail := range transactionDetails {
		transactionDetailsResp.TransactionDetails = append(transactionDetailsResp.TransactionDetails, MakeTransactionDetailResponse(transactionDetail))
	}
	return transactionDetailsResp
}

type PLTHolderInfoReq struct {
	Address string
}

type PLTHolderInfoResp struct {
	Address string
	Amount  string
	Percent string `json:",omitempty"`
}

func MakePLTHolderInfoResponse(pltContract *PLTHolderWithPercent) *PLTHolderInfoResp {
	pltHolderInfoResp := &PLTHolderInfoResp{
		Address: pltContract.Address,
		Amount:  utils.AmountWithoutPrecision(pltContract.Amount),
		Percent: basedef.FromatPercent(pltContract.Percent),
	}
	return pltHolderInfoResp
}

type PLTHoldersReq struct {
	PageSize int
	PageNo   int
}

type PLTHoldersResp struct {
	PageSize       int
	PageNo         int
	TotalPage      int
	TotalCount     int
	PLTHolderInfos []*PLTHolderInfoResp
}

func MakePLTHoldersResponse(pageSize int, pageNo int, totalPage int, totalCount int, pltContracts []*PLTHolderWithPercent) *PLTHoldersResp {
	pltHoldersResp := &PLTHoldersResp{
		PageSize:   pageSize,
		PageNo:     pageNo,
		TotalPage:  totalPage,
		TotalCount: totalCount,
	}
	for _, pltContract := range pltContracts {
		pltHoldersResp.PLTHolderInfos = append(pltHoldersResp.PLTHolderInfos, MakePLTHolderInfoResponse(pltContract))
	}
	return pltHoldersResp
}

type ContractInfoReq struct {
	Contract string
}

type ContractInfoResp struct {
	Contract    string
	Type        uint64
	Name        string
	Symbol      string
	Owner       string
	Uri         string
	Site        string
	Description string
	BaseUri string
	Time        uint64
	TotalSupply uint64
	AddressNum  uint64
	TransferNum uint64
}

func MakeContractInfoResponse(nftContract *ContractInfo) *ContractInfoResp {
	nftResp := &ContractInfoResp{
		Contract:    nftContract.Contract,
		Type:        nftContract.Type,
		Name:        nftContract.Name,
		Symbol:      nftContract.Symbol,
		Owner:       nftContract.Owner,
		Uri:         nftContract.Uri,
		Site:        nftContract.Site,
		Time:        nftContract.Time,
		BaseUri: nftContract.BaseUri,
		Description: nftContract.Description,
		TotalSupply: nftContract.TotalSupply,
		AddressNum:  nftContract.AddressNum,
		TransferNum: nftContract.TransferNum,
	}
	return nftResp
}

type ContractInfosReq struct {
	PageSize int
	PageNo   int
}

type ContractInfosResp struct {
	PageSize   int
	PageNo     int
	TotalPage  int
	TotalCount int
	Contracts  []*ContractInfoResp
}

func MakeContractInfosResponse(pageSize int, pageNo int, totalPage int, totalCount int, nftContracts []*ContractInfo) *ContractInfosResp {
	nftsResp := &ContractInfosResp{
		PageSize:   pageSize,
		PageNo:     pageNo,
		TotalPage:  totalPage,
		TotalCount: totalCount,
	}
	for _, nftContract := range nftContracts {
		nftsResp.Contracts = append(nftsResp.Contracts, MakeContractInfoResponse(nftContract))
	}
	return nftsResp
}

type NFTHolderReq struct {
	Contract   string
	Token string
}

type NFTHolderResp struct {
	Contract   string
	Token string
	Owner string
	Uri   string
}

func MakeNFTHolderResponse(nftContract *NFTHolder) *NFTHolderResp {
	nftTokenInfoResp := &NFTHolderResp{
		Contract:   nftContract.NFT,
		Token: nftContract.Token,
		Owner: nftContract.Owner,
		Uri:   nftContract.Uri,
	}
	return nftTokenInfoResp
}

func MakeNFTHolderWithUriResponse(nftContract *NFTHolderWithUri) *NFTHolderResp {
	nftTokenInfoResp := &NFTHolderResp{
		Contract:   nftContract.NFT,
		Token: nftContract.Token,
		Owner: nftContract.Owner,
		Uri:   nftContract.Uri,
	}
	if nftContract.ContractInfo != nil {
		nftTokenInfoResp.Uri = nftContract.ContractInfo.BaseUri + nftContract.Uri
	}
	return nftTokenInfoResp
}

type NFTHoldersReq struct {
	Contract      string
	PageSize int
	PageNo   int
}

type NFTHoldersOfUserReq struct {
	Contract      string
	Address  string
	PageSize int
	PageNo   int
}

type NFTHoldersResp struct {
	PageSize      int
	PageNo        int
	TotalPage     int
	TotalCount    int
	NFTTokenInfos []*NFTHolderResp
}

func MakeNFTHoldersResponse(pageSize int, pageNo int, totalPage int, totalCount int, nftContracts []*NFTHolderWithUri) *NFTHoldersResp {
	nftHoldersResp := &NFTHoldersResp{
		PageSize:   pageSize,
		PageNo:     pageNo,
		TotalPage:  totalPage,
		TotalCount: totalCount,
	}
	for _, nftContract := range nftContracts {
		nftHoldersResp.NFTTokenInfos = append(nftHoldersResp.NFTTokenInfos, MakeNFTHolderWithUriResponse(nftContract))
	}
	return nftHoldersResp
}

type NFTUserReq struct {
	Contract   string
	Owner string
}

type NFTUserResp struct {
	Contract         string
	Owner       string
	TokenNumber uint64
	ContractInfo    *ContractInfoResp `json:",omitempty"`
	Percent     string `json:",omitempty"`
}

func MakeNFTUserResponse(nftUser *NFTUser) *NFTUserResp {
	nftUserResp := &NFTUserResp{
		Contract:         nftUser.NFT,
		Owner:       nftUser.Owner,
		TokenNumber: nftUser.TokenNumber,
		Percent:     basedef.FromatPercent(nftUser.Percent),
	}
	if nftUser.ContractInfo != nil {
		nftUserResp.ContractInfo = MakeContractInfoResponse(nftUser.ContractInfo)
	}
	return nftUserResp
}

type NFTUsersReq struct {
	Contract      string
	Owner       string
	PageSize int
	PageNo   int
}

type NFTUsersResp struct {
	PageSize   int
	PageNo     int
	TotalPage  int
	TotalCount int
	NFTUsers   []*NFTUserResp
}

func MakeNFTUsersResponse(pageSize int, pageNo int, totalPage int, totalCount int, nftContracts []*NFTUser) *NFTUsersResp {
	nftUsers := &NFTUsersResp{
		PageSize:   pageSize,
		PageNo:     pageNo,
		TotalPage:  totalPage,
		TotalCount: totalCount,
	}
	for _, nftContract := range nftContracts {
		nftUsers.NFTUsers = append(nftUsers.NFTUsers, MakeNFTUserResponse(nftContract))
	}
	return nftUsers
}

type StakeInfoReq struct {
	Owner     string
	Validator string
}

type StakeInfoResp struct {
	Owner     string
	Validator string
	Amount    string
}

func MakeStakeInfoResponse(stake *Stake) *StakeInfoResp {
	stakeInfoResp := &StakeInfoResp{
		Owner:     stake.Owner,
		Validator: stake.Validator,
		Amount:    utils.AmountWithoutPrecision(stake.StakeAmount),
	}
	return stakeInfoResp
}

type StakesOfOwnerReq struct {
	Owner    string
	PageSize int
	PageNo   int
}

type StakesOfOwnerResp struct {
	PageSize   int
	PageNo     int
	TotalPage  int
	TotalCount int
	StakeInfos []*StakeInfoResp
}

func MakeStakesOfOwnerResponse(pageSize int, pageNo int, totalPage int, totalCount int, stakes []*Stake) *StakesOfOwnerResp {
	stakeByOwnerResp := &StakesOfOwnerResp{
		PageSize:   pageSize,
		PageNo:     pageNo,
		TotalPage:  totalPage,
		TotalCount: totalCount,
	}
	for _, stake := range stakes {
		stakeByOwnerResp.StakeInfos = append(stakeByOwnerResp.StakeInfos, MakeStakeInfoResponse(stake))
	}
	return stakeByOwnerResp
}

type StakesOfValidatorReq struct {
	Validator string
	PageSize  int
	PageNo    int
}

type StakesOfValidatorResp struct {
	PageSize   int
	PageNo     int
	TotalPage  int
	TotalCount int
	StakeInfos []*StakeInfoResp
}

func MakeStakesOfValidatorResponse(pageSize int, pageNo int, totalPage int, totalCount int, stakes []*Stake) *StakesOfValidatorResp {
	stakesOfValidatorResp := &StakesOfValidatorResp{
		PageSize:   pageSize,
		PageNo:     pageNo,
		TotalPage:  totalPage,
		TotalCount: totalCount,
	}
	for _, stake := range stakes {
		stakesOfValidatorResp.StakeInfos = append(stakesOfValidatorResp.StakeInfos, MakeStakeInfoResponse(stake))
	}
	return stakesOfValidatorResp
}

type ValidatorInfoReq struct {
	Address string
}

type ValidatorInfoResp struct {
	Address        string
	DelegateFactor string
	StakeAmount    string
	Name           string
	Uri            string
	Percent        string `json:",omitempty"`
}

func MakeValidatorInfoResponse(validator *ValidatorWithPercent) *ValidatorInfoResp {
	bbb := decimal.NewFromInt(int64(validator.DelegateFactor))
	ccc := bbb.Div(decimal.NewFromInt(100))
	delegateFactor := fmt.Sprintf("%s%s", ccc.String(), "%")
	validatorInfoResp := &ValidatorInfoResp{
		Address:        validator.Address,
		DelegateFactor: delegateFactor,
		StakeAmount:    utils.AmountWithoutPrecision(validator.StakeAmount),
		Percent:        basedef.FromatPercent(validator.Percent),
	}
	return validatorInfoResp
}

type ValidatorsReq struct {
	PageSize int
	PageNo   int
}

type ValidatorsResp struct {
	PageSize       int
	PageNo         int
	TotalPage      int
	TotalCount     int
	ValidatorInfos []*ValidatorInfoResp
}

func MakeValidatorsResponse(pageSize int, pageNo int, totalPage int, totalCount int, validators []*ValidatorWithPercent) *ValidatorsResp {
	validatorsResp := &ValidatorsResp{
		PageSize:   pageSize,
		PageNo:     pageNo,
		TotalPage:  totalPage,
		TotalCount: totalCount,
	}
	for _, validator := range validators {
		validatorsResp.ValidatorInfos = append(validatorsResp.ValidatorInfos, MakeValidatorInfoResponse(validator))
	}
	return validatorsResp
}

type ProposeInfoReq struct {
	ProposeId string
}

type ProposeInfoResp struct {
	ProposeId string
	Proposer  string
	Value     string
	EndBlock  uint64
	Type      uint8
	State     uint8
}

func MakeProposeInfoResponse(propose *Propose) *ProposeInfoResp {
	proposeInfoResp := &ProposeInfoResp{
		ProposeId: propose.ProposeId,
		Proposer:  propose.Proposer,
		Value:     propose.Value,
		EndBlock:  propose.EndBlock,
		Type:      propose.Type,
		State:     propose.State,
	}
	return proposeInfoResp
}

type ProposesReq struct {
	PageSize int
	PageNo   int
}

type ProposesResp struct {
	PageSize     int
	PageNo       int
	TotalPage    int
	TotalCount   int
	ProposeInfos []*ProposeInfoResp
}

func MakeProposesResponse(pageSize int, pageNo int, totalPage int, totalCount int, proposes []*Propose) *ProposesResp {
	proposesResp := &ProposesResp{
		PageSize:   pageSize,
		PageNo:     pageNo,
		TotalPage:  totalPage,
		TotalCount: totalCount,
	}
	for _, propose := range proposes {
		proposesResp.ProposeInfos = append(proposesResp.ProposeInfos, MakeProposeInfoResponse(propose))
	}
	return proposesResp
}

type ProposesOfUserReq struct {
	Proposer string
	PageSize int
	PageNo   int
}

type ProposesOfUserResp struct {
	PageSize     int
	PageNo       int
	TotalPage    int
	TotalCount   int
	ProposeInfos []*ProposeInfoResp
}

func MakeProposesOfUserResponse(pageSize int, pageNo int, totalPage int, totalCount int, proposes []*Propose) *ProposesOfUserResp {
	proposesOfUserResp := &ProposesOfUserResp{
		PageSize:   pageSize,
		PageNo:     pageNo,
		TotalPage:  totalPage,
		TotalCount: totalCount,
	}
	for _, propose := range proposes {
		proposesOfUserResp.ProposeInfos = append(proposesOfUserResp.ProposeInfos, MakeProposeInfoResponse(propose))
	}
	return proposesOfUserResp
}


type TransactionDetailsOfUserReq struct {
	Contract      string
	User string
	PageSize int
	PageNo   int
}
