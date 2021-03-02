package controllers

import (
	"encoding/json"
	"explorer/basedef"
	"explorer/models"
	"github.com/astaxie/beego"
)

type NFTContractController struct {
	beego.Controller
}

func (c *NFTContractController) NFTs() {
	var nftsReq models.ContractInfosReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &nftsReq); err != nil {
		panic(err)
	}
	nftContracts := make([]*models.ContractInfo, 0)
	db.Limit(nftsReq.PageSize).Offset(nftsReq.PageSize * nftsReq.PageNo).Order("time desc").Where("`type` = ?", basedef.CONTRACT_TYPE_NFT).Find(&nftContracts)
	var nftContractNum int64
	db.Model(&models.ContractInfo{}).Where("`type` = ?", basedef.CONTRACT_TYPE_NFT).Count(&nftContractNum)
	c.Data["json"] = models.MakeContractInfosResponse(nftsReq.PageSize, nftsReq.PageNo,
		(int(nftContractNum)+nftsReq.PageSize-1)/nftsReq.PageSize, int(nftContractNum), nftContracts)
	c.ServeJSON()
}

func (c *NFTContractController) NFT() {
	var nftReq models.ContractInfoReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &nftReq); err != nil {
		panic(err)
	}
	var nftContract models.ContractInfo
	db.Where("contract = ?", nftReq.Contract).Find(&nftContract)
	c.Data["json"] = models.MakeContractInfoResponse(&nftContract)
	c.ServeJSON()
}

func (c *NFTContractController) NFTTokenInfo() {
	var nftTokenInfoReq models.NFTHolderReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &nftTokenInfoReq); err != nil {
		panic(err)
	}
	nftContract := new(models.NFTHolder)
	db.Where("nft = ? and token = ?", nftTokenInfoReq.Contract, nftTokenInfoReq.Token).First(nftContract)
	c.Data["json"] = models.MakeNFTHolderResponse(nftContract)
	c.ServeJSON()
}

func (c *NFTContractController) NFTTokenTransactions() {
	var transactionDetailsReq models.TransactionDetailsOfNFTTokenReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &transactionDetailsReq); err != nil {
		panic(err)
	}
	transactionDetails := make([]*models.TransactionDetailWithInfo, 0)
	db.Where("contract = ? and value = ?", transactionDetailsReq.Contract, transactionDetailsReq.Token).Preload("NFTHolder").Limit(transactionDetailsReq.PageSize).Offset(transactionDetailsReq.PageSize * transactionDetailsReq.PageNo).Order("time desc").Find(&transactionDetails)
	var transactionDetailsNum int64
	db.Model(&models.TransactionDetailWithInfo{}).Where("contract = ? and value = ?", transactionDetailsReq.Contract, transactionDetailsReq.Token).Count(&transactionDetailsNum)
	c.Data["json"] = models.MakeTransactionDetailsResponse(transactionDetailsReq.PageSize, transactionDetailsReq.PageNo,
		(int(transactionDetailsNum)+transactionDetailsReq.PageSize-1)/transactionDetailsReq.PageSize, int(transactionDetailsNum), transactionDetails)
	c.ServeJSON()
}

func (c *NFTContractController) NFTHoldersOfUser() {
	var nftHoldersOfUserReq models.NFTHoldersOfUserReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &nftHoldersOfUserReq); err != nil {
		panic(err)
	}
	nftContracts := make([]*models.NFTHolder, 0)
	db.Where("nft = ? and owner = ?", nftHoldersOfUserReq.Contract, nftHoldersOfUserReq.Address).Limit(nftHoldersOfUserReq.PageSize).Offset(nftHoldersOfUserReq.PageSize * nftHoldersOfUserReq.PageNo).Find(&nftContracts)
	var nftContractNum int64
	db.Model(&models.NFTHolder{}).Where("nft = ? and owner = ?", nftHoldersOfUserReq.Contract, nftHoldersOfUserReq.Address).Count(&nftContractNum)
	c.Data["json"] = models.MakeNFTHoldersResponse(nftHoldersOfUserReq.PageSize, nftHoldersOfUserReq.PageNo,
		(int(nftContractNum)+nftHoldersOfUserReq.PageSize-1)/nftHoldersOfUserReq.PageSize, int(nftContractNum), nftContracts)
	c.ServeJSON()
}

func (c *NFTContractController) NFTHolders() {
	var nftHoldersReq models.NFTHoldersReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &nftHoldersReq); err != nil {
		panic(err)
	}
	nftContracts := make([]*models.NFTHolder, 0)
	db.Where("nft = ?", nftHoldersReq.Contract).Limit(nftHoldersReq.PageSize).Offset(nftHoldersReq.PageSize * nftHoldersReq.PageNo).Find(&nftContracts)
	var nftTokenNum int64
	db.Model(&models.NFTHolder{}).Where("nft = ?", nftHoldersReq.Contract).Count(&nftTokenNum)
	c.Data["json"] = models.MakeNFTHoldersResponse(nftHoldersReq.PageSize, nftHoldersReq.PageNo,
		(int(nftTokenNum)+nftHoldersReq.PageSize-1)/nftHoldersReq.PageSize, int(nftTokenNum), nftContracts)
	c.ServeJSON()
}

func (c *NFTContractController) NFTUsers() {
	var nftUsersReq models.NFTUsersReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &nftUsersReq); err != nil {
		panic(err)
	}
	nftContracts := make([]*models.NFTUser, 0)
	db.Table("(?) as u, nft_holders", db.Model(&models.NFTUser{}).Select("count(*) as total").Where("nft = ?", nftUsersReq.Contract)).
		Select("nft, owner, count(*) as token_number, count(*)/total as percent").Where("nft = ?", nftUsersReq.Contract).Group("owner").Order("count(*)/total desc").
		Limit(nftUsersReq.PageSize).Offset(nftUsersReq.PageSize * nftUsersReq.PageNo).Find(&nftContracts)

	var nftTokenNum int64
	db.Table("(?) as u", db.Model(&models.NFTUser{}).Select("nft, owner").Where("nft = ?", nftUsersReq.Contract).Group("owner")).Count(&nftTokenNum)
	c.Data["json"] = models.MakeNFTUsersResponse(nftUsersReq.PageSize, nftUsersReq.PageNo,
		(int(nftTokenNum)+nftUsersReq.PageSize-1)/nftUsersReq.PageSize, int(nftTokenNum), nftContracts)
	c.ServeJSON()
}

func (c *NFTContractController) NFTTransactions() {
	var transactionDetailsReq models.TransactionDetailsOfContractReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &transactionDetailsReq); err != nil {
		panic(err)
	}
	transactionDetails := make([]*models.TransactionDetailWithInfo, 0)
	db.Where("contract = ?", transactionDetailsReq.Contract).Preload("NFTHolder").Limit(transactionDetailsReq.PageSize).Offset(transactionDetailsReq.PageSize * transactionDetailsReq.PageNo).Order("time desc").Find(&transactionDetails)
	var transactionDetailsNum int64
	db.Model(&models.TransactionDetailWithInfo{}).Where("contract = ?", transactionDetailsReq.Contract).Count(&transactionDetailsNum)
	c.Data["json"] = models.MakeTransactionDetailsResponse(transactionDetailsReq.PageSize, transactionDetailsReq.PageNo,
		(int(transactionDetailsNum)+transactionDetailsReq.PageSize-1)/transactionDetailsReq.PageSize, int(transactionDetailsNum), transactionDetails)
	c.ServeJSON()
}
