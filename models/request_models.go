package models

type ExplorerResp struct {
	Version   string
	URL       string
}

type ChainResp struct {
	Id     uint64
	Name   string
	Height uint64
	StakeAmount uint64
	LastReward uint64
	MintPrice uint64
	RewardPeriod uint64
	GasFee uint64
}

func MakeChainResponse(chain *Chain) *ChainResp {
	chainResp := &ChainResp{
		Id:           chain.Id,
		Name:         chain.Name,
		Height:       chain.Height,
		StakeAmount:  chain.StakeAmount,
		LastReward:   chain.LastReward,
		MintPrice:    chain.MintPrice,
		RewardPeriod: chain.RewardPeriod,
		GasFee:       chain.GasFee,
	}
	return chainResp
}

type BlockResp struct {
	Hash    string
	GasLimit  uint64
	GasUsed  uint64
	Difficulty uint64
	Number     uint64
	ParentHash string
	ReceiptHash  string
	TxHash  string
	TxNumber uint64
	Time  uint64
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
		Hash:         block.Hash,
		GasLimit:     block.GasLimit,
		GasUsed:      block.GasUsed,
		Difficulty:   block.Difficulty,
		Number:       block.Number,
		ParentHash:   block.ParentHash,
		ReceiptHash:  block.ReceiptHash,
		TxHash:       block.TxHash,
		TxNumber:     block.TxNumber,
		Time:         block.Time,
	}
	for _, transaction := range block.Transactions {
		blockOut.Transactions = append(blockOut.Transactions, transaction.Hash)
	}
	return blockOut
}

type BlocksReq struct {
	PageSize int
	PageNo int
}

type BlocksResp struct {
	PageSize  int
	PageNo int
	TotalPage int
	TotalCount int
	Blocks []*BlockResp
}

func MakeBlocksResponse(pageSize int, pageNo int, totalPage int, totalCount int, blocks []*Block) *BlocksResp {
	blocksResp := &BlocksResp{
		PageSize: pageSize,
		PageNo: pageNo,
		TotalPage: totalPage,
		TotalCount: totalCount,
	}
	for _, block := range blocks {
		blocksResp.Blocks = append(blocksResp.Blocks, MakeBlockResponse(block))
	}
	return blocksResp
}

type TransactionResp struct {
	Hash  string
	From  string
	Cost  uint64
	Data  string
	Gas   uint64
	GasPrice uint64
	To    string
	Value uint64
	Time  uint64
	BlockNumber  uint64
	BlockHash  string
	Events []*EventResp
	TransactionDetails []*TransactionDetailResp
}

type TransactionByHashReq struct {
	Hash string
}

func MakeTransactionResponse(transaction *Transaction) *TransactionResp {
	transactionResp := &TransactionResp{
		Hash:               transaction.Hash,
		From:               transaction.From,
		Cost:               transaction.Cost,
		Data:               transaction.Data,
		Gas:                transaction.Gas,
		GasPrice:           transaction.GasPrice,
		To:                 transaction.To,
		Value:              transaction.Value,
		BlockNumber:        transaction.BlockNumber,
		Time:               transaction.Time,
		BlockHash:          transaction.BlockHash,
	}
	for _, event := range transaction.Events {
		transactionResp.Events = append(transactionResp.Events, MakeEventResponse(event))
	}
	for _, transactionDetail := range transaction.TransactionDetails {
		transactionResp.TransactionDetails = append(transactionResp.TransactionDetails, MakeTransactionDetailResponse(transactionDetail))
	}
	return transactionResp
}


type EventResp struct {
	Number   uint64
	Contract string
	EventId  string
	Topic1 string
	Topic2 string
	Topic3 string
	Topic4 string
	Data string
	TransactionHash  string
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
		TransactionHash: event.TransactionHash,
	}
	return eventResp
}

type TransactionDetailResp struct {
	Contract string
	From  string
	To string
	Value string
	TransactionHash  string
}

func MakeTransactionDetailResponse(transactionDetail *TransactionDetail) *TransactionDetailResp {
	transactionDetailResp := &TransactionDetailResp{
		Contract:        transactionDetail.Contract,
		From:            transactionDetail.From,
		To:              transactionDetail.To,
		Value:           transactionDetail.Value,
		TransactionHash: transactionDetail.TransactionHash,
	}
	return transactionDetailResp
}

type TransactionsReq struct {
	PageSize int
	PageNo int
}

type TransactionsOfContractReq struct {
	Contract string
	PageSize int
	PageNo int
}

type TransactionsOfUserReq struct {
	User string
	PageSize int
	PageNo int
}

type TransactionsResp struct {
	PageSize  int
	PageNo int
	TotalPage int
	TotalCount int
	Transactions []*TransactionResp
}

func MakeTransactionsResponse(pageSize int, pageNo int, totalPage int, totalCount int, transactions []*Transaction) *TransactionsResp {
	transactionsResp := &TransactionsResp{
		PageSize: pageSize,
		PageNo: pageNo,
		TotalPage: totalPage,
		TotalCount: totalCount,
	}
	for _, transaction := range transactions {
		transactionsResp.Transactions = append(transactionsResp.Transactions, MakeTransactionResponse(transaction))
	}
	return transactionsResp
}

type PLTHolderInfoReq struct {
	Address string
}

type PLTHolderInfoResp struct {
	Address string
	Amount uint64
}

func MakePLTHolderInfoResponse(pltContract *PLTContract) *PLTHolderInfoResp {
	pltHolderInfoResp := &PLTHolderInfoResp {
		Address: pltContract.Address,
		Amount: pltContract.Amount,
	}
	return pltHolderInfoResp
}

type PLTHoldersReq struct {
	PageSize int
	PageNo int
}

type PLTHoldersResp struct {
	PageSize  int
	PageNo int
	TotalPage int
	TotalCount int
	PLTHolderInfos []*PLTHolderInfoResp
}

func MakePLTHoldersResponse(pageSize int, pageNo int, totalPage int, totalCount int, pltContracts []*PLTContract) *PLTHoldersResp {
	pltHoldersResp := &PLTHoldersResp{
		PageSize: pageSize,
		PageNo: pageNo,
		TotalPage: totalPage,
		TotalCount: totalCount,
	}
	for _, pltContract := range pltContracts {
		pltHoldersResp.PLTHolderInfos = append(pltHoldersResp.PLTHolderInfos, MakePLTHolderInfoResponse(pltContract))
	}
	return pltHoldersResp
}

type NFTResp struct {
	NFT string
}

func MakeNFTResponse(nftContract *NFTContract) *NFTResp {
	nftResp := &NFTResp{NFT:nftContract.NFT}
	return nftResp
}


type NFTsReq struct {
	PageSize int
	PageNo int
}

type NFTsResp struct {
	PageSize  int
	PageNo int
	TotalPage int
	TotalCount int
	NFTs []*NFTResp
}

func MakeNFTsResponse(pageSize int, pageNo int, totalPage int, totalCount int, nftContracts []*NFTContract) *NFTsResp {
	nftsResp := &NFTsResp{
		PageSize: pageSize,
		PageNo: pageNo,
		TotalPage: totalPage,
		TotalCount: totalCount,
	}
	for _, nftContract := range nftContracts {
		nftsResp.NFTs = append(nftsResp.NFTs, MakeNFTResponse(nftContract))
	}
	return nftsResp
}


type NFTTokenResp struct {
	NFT string
	Token string
}

func MakeNFTTokenResponse(nftContract *NFTContract) *NFTTokenResp {
	nftResp := &NFTTokenResp{
		NFT: nftContract.NFT,
		Token: nftContract.Token,
	}
	return nftResp
}

type NFTTokensReq struct {
	NFT string
	PageSize int
	PageNo int
}

type NFTTokensResp struct {
	PageSize  int
	PageNo int
	TotalPage int
	TotalCount int
	NFTTokens []*NFTTokenResp
}

func MakeNFTTokensResponse(pageSize int, pageNo int, totalPage int, totalCount int, nftContracts []*NFTContract) *NFTTokensResp {
	nftTokensResp := &NFTTokensResp{
		PageSize: pageSize,
		PageNo: pageNo,
		TotalPage: totalPage,
		TotalCount: totalCount,
	}
	for _, nftContract := range nftContracts {
		nftTokensResp.NFTTokens = append(nftTokensResp.NFTTokens, MakeNFTTokenResponse(nftContract))
	}
	return nftTokensResp
}


type NFTTokenInfoReq struct {
	NFT string
	Token string
}

type NFTTokenInfoResp struct {
	NFT string
	Token string
	Owner string
	URI string
}


func MakeNFTTokenInfoResponse(nftContract *NFTContract) *NFTTokenInfoResp {
	nftTokenInfoResp := &NFTTokenInfoResp{
		NFT: nftContract.NFT,
		Token: nftContract.Token,
		Owner: nftContract.Owner,
		URI: nftContract.Uri,
	}
	return nftTokenInfoResp
}

type NFTHoldersOfUserReq struct {
	NFT string
	Address string
	PageSize int
	PageNo int
}

type NFTHoldersOfUserResp struct {
	PageSize  int
	PageNo int
	TotalPage int
	TotalCount int
	NFTTokenInfos []*NFTTokenInfoResp
}

func MakeNFTHoldersOfUserResponse(pageSize int, pageNo int, totalPage int, totalCount int, nftContracts []*NFTContract) *NFTHoldersOfUserResp {
	nftHolderByAddressResp := &NFTHoldersOfUserResp{
		PageSize: pageSize,
		PageNo: pageNo,
		TotalPage: totalPage,
		TotalCount: totalCount,
	}
	for _, nftContract := range nftContracts {
		nftHolderByAddressResp.NFTTokenInfos = append(nftHolderByAddressResp.NFTTokenInfos, MakeNFTTokenInfoResponse(nftContract))
	}
	return nftHolderByAddressResp
}


type NFTHoldersReq struct {
	NFT string
	PageSize int
	PageNo int
}

type NFTHoldersResp struct {
	PageSize  int
	PageNo int
	TotalPage int
	TotalCount int
	NFTTokenInfos []*NFTTokenInfoResp
}

func MakeNFTHoldersResponse(pageSize int, pageNo int, totalPage int, totalCount int, nftContracts []*NFTContract) *NFTHoldersResp {
	nftHoldersResp := &NFTHoldersResp{
		PageSize: pageSize,
		PageNo: pageNo,
		TotalPage: totalPage,
		TotalCount: totalCount,
	}
	for _, nftContract := range nftContracts {
		nftHoldersResp.NFTTokenInfos = append(nftHoldersResp.NFTTokenInfos, MakeNFTTokenInfoResponse(nftContract))
	}
	return nftHoldersResp
}

type StakeInfoReq struct {
	Owner string
	Validator string
}

type StakeInfoResp struct {
	Owner string
	Validator string
	Amount uint64
}

func MakeStakeInfoResponse(stake *Stake) *StakeInfoResp {
	stakeInfoResp := &StakeInfoResp{
		Owner:     stake.Owner,
		Validator: stake.Validator,
		Amount:    stake.StakeAmount,
	}
	return stakeInfoResp
}


type StakesOfOwnerReq struct {
	Owner string
	PageSize int
	PageNo int
}

type StakesOfOwnerResp struct {
	PageSize  int
	PageNo int
	TotalPage int
	TotalCount int
	StakeInfos []*StakeInfoResp
}

func MakeStakesOfOwnerResponse(pageSize int, pageNo int, totalPage int, totalCount int, stakes []*Stake) *StakesOfOwnerResp {
	stakeByOwnerResp := &StakesOfOwnerResp{
		PageSize: pageSize,
		PageNo: pageNo,
		TotalPage: totalPage,
		TotalCount: totalCount,
	}
	for _, stake := range stakes {
		stakeByOwnerResp.StakeInfos = append(stakeByOwnerResp.StakeInfos, MakeStakeInfoResponse(stake))
	}
	return stakeByOwnerResp
}

type StakesOfValidatorReq struct {
	Validator string
	PageSize int
	PageNo int
}

type StakesOfValidatorResp struct {
	PageSize  int
	PageNo int
	TotalPage int
	TotalCount int
	StakeInfos []*StakeInfoResp
}

func MakeStakesOfValidatorResponse(pageSize int, pageNo int, totalPage int, totalCount int, stakes []*Stake) *StakesOfValidatorResp {
	stakesOfValidatorResp := &StakesOfValidatorResp{
		PageSize: pageSize,
		PageNo: pageNo,
		TotalPage: totalPage,
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
	Address string
	DelegateFactor uint64
	StakeAmount uint64
}

func MakeValidatorInfoResponse(validator *Validator) *ValidatorInfoResp {
	validatorInfoResp := &ValidatorInfoResp{
		Address:     validator.Address,
		DelegateFactor: validator.DelegateFactor,
		StakeAmount:    validator.StakeAmount,
	}
	return validatorInfoResp
}

type ValidatorsReq struct {
	PageSize int
	PageNo int
}

type ValidatorsResp struct {
	PageSize  int
	PageNo int
	TotalPage int
	TotalCount int
	ValidatorInfos []*ValidatorInfoResp
}

func MakeValidatorsResponse(pageSize int, pageNo int, totalPage int, totalCount int, validators []*Validator) *ValidatorsResp {
	validatorsResp := &ValidatorsResp{
		PageSize: pageSize,
		PageNo: pageNo,
		TotalPage: totalPage,
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
	ProposeId    string
	Proposer     string
	Value        string
	EndBlock     uint64
	Type         uint8
	State        uint8
}

func MakeProposeInfoResponse(propose *Propose) *ProposeInfoResp {
	proposeInfoResp := &ProposeInfoResp{
		ProposeId:     propose.ProposeId,
		Proposer: propose.Proposer,
		Value:    propose.Value,
		EndBlock: propose.EndBlock,
		Type: propose.Type,
		State: propose.State,
	}
	return proposeInfoResp
}

type ProposesReq struct {
	PageSize int
	PageNo int
}

type ProposesResp struct {
	PageSize  int
	PageNo int
	TotalPage int
	TotalCount int
	ProposeInfos []*ProposeInfoResp
}

func MakeProposesResponse(pageSize int, pageNo int, totalPage int, totalCount int, proposes []*Propose) *ProposesResp {
	proposesResp := &ProposesResp{
		PageSize: pageSize,
		PageNo: pageNo,
		TotalPage: totalPage,
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
	PageNo int
}

type ProposesOfUserResp struct {
	PageSize  int
	PageNo int
	TotalPage int
	TotalCount int
	ProposeInfos []*ProposeInfoResp
}

func MakeProposesOfUserResponse(pageSize int, pageNo int, totalPage int, totalCount int, proposes []*Propose) *ProposesOfUserResp {
	proposesOfUserResp := &ProposesOfUserResp{
		PageSize: pageSize,
		PageNo: pageNo,
		TotalPage: totalPage,
		TotalCount: totalCount,
	}
	for _, propose := range proposes {
		proposesOfUserResp.ProposeInfos = append(proposesOfUserResp.ProposeInfos, MakeProposeInfoResponse(propose))
	}
	return proposesOfUserResp
}





