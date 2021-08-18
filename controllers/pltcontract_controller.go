package controllers

import (
	"encoding/json"
	"explorer/models"
	"github.com/astaxie/beego"
	"github.com/ethereum/go-ethereum/contracts/native"
)

type PLTContractController struct {
	beego.Controller
}

func (c *PLTContractController) PLT() {
	var pltReq models.ContractInfoReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &pltReq); err != nil {
		panic(err)
	}
	var pltContract models.ContractInfo
	db.Where("contract = ?", native.PLTContractAddress).Find(&pltContract)
	c.Data["json"] = models.MakeContractInfoResponse(&pltContract)
	c.ServeJSON()
}

func (c *PLTContractController) PLTHolderInfo() {
	var pltHolderInfoReq models.PLTHolderInfoReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &pltHolderInfoReq); err != nil {
		panic(err)
	}
	pltContract := new(models.PLTHolderWithPercent)
	db.Where("address = ?", pltHolderInfoReq.Address).First(pltContract)
	c.Data["json"] = models.MakePLTHolderInfoResponse(pltContract)
	c.ServeJSON()
}

func (c *PLTContractController) PLTHolders() {
	var pltHoldersReq models.PLTHoldersReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &pltHoldersReq); err != nil {
		panic(err)
	}
	pltContracts := make([]*models.PLTHolderWithPercent, 0)
	db.Table("(?) as u, plt_holders", db.Model(&models.PLTHolderWithPercent{}).Select("sum(amount) as total").Where("amount > 0")).
		Select("address, amount, amount/total as percent").Where("amount >= 0").Limit(pltHoldersReq.PageSize).Offset(pltHoldersReq.PageSize * pltHoldersReq.PageNo).Order("amount desc").Find(&pltContracts)
	var pltContractNum int64
	db.Model(&models.PLTHolderWithPercent{}).Count(&pltContractNum)
	c.Data["json"] = models.MakePLTHoldersResponse(pltHoldersReq.PageSize, pltHoldersReq.PageNo,
		(int(pltContractNum)+pltHoldersReq.PageSize-1)/pltHoldersReq.PageSize, int(pltContractNum), pltContracts)
	c.ServeJSON()
}

func (c *PLTContractController) PLTTransactions() {
	var transactionDetailsReq models.TransactionDetailsOfContractReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &transactionDetailsReq); err != nil {
		panic(err)
	}
	transactionDetails := make([]*models.TransactionDetailWithInfo, 0)
	db.Where("contract = ?", native.PLTContractAddress).Limit(transactionDetailsReq.PageSize).Offset(transactionDetailsReq.PageSize * transactionDetailsReq.PageNo).Order("time desc").Find(&transactionDetails)
	var transactionDetailsNum int64
	db.Model(&models.TransactionDetailWithInfo{}).Where("contract = ?", native.PLTContractAddress).Count(&transactionDetailsNum)
	c.Data["json"] = models.MakeTransactionDetailsResponse(transactionDetailsReq.PageSize, transactionDetailsReq.PageNo,
		(int(transactionDetailsNum)+transactionDetailsReq.PageSize-1)/transactionDetailsReq.PageSize, int(transactionDetailsNum), transactionDetails)
	c.ServeJSON()
}

func (c *PLTContractController) PLTTransactionsOfUser() {
	var transactionDetailsOfUserReq models.TransactionDetailsOfUserReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &transactionDetailsOfUserReq); err != nil {
		panic(err)
	}
	transactionDetails := make([]*models.TransactionDetailWithInfo, 0)
	db.Where("contract = ?", native.PLTContractAddress).
		Where("`from` = ? or `to` = ?", transactionDetailsOfUserReq.User, transactionDetailsOfUserReq.User).
		Limit(transactionDetailsOfUserReq.PageSize).Offset(transactionDetailsOfUserReq.PageSize * transactionDetailsOfUserReq.PageNo).Order("time desc").Find(&transactionDetails)
	var transactionDetailsNum int64
	db.Model(&models.TransactionDetailWithInfo{}).Where("contract = ?", native.PLTContractAddress).Where("`from` = ? or `to` = ?", transactionDetailsOfUserReq.User, transactionDetailsOfUserReq.User).Count(&transactionDetailsNum)
	c.Data["json"] = models.MakeTransactionDetailsResponse(transactionDetailsOfUserReq.PageSize, transactionDetailsOfUserReq.PageNo,
		(int(transactionDetailsNum)+transactionDetailsOfUserReq.PageSize-1)/transactionDetailsOfUserReq.PageSize, int(transactionDetailsNum), transactionDetails)
	c.ServeJSON()
}
