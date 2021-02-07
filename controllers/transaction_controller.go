package controllers

import (
	"encoding/json"
	"explorer/models"
	"github.com/astaxie/beego"
)

type TransactionController struct {
	beego.Controller
}

func (c *TransactionController) TransactionByHash() {
	var transactionByHashReq models.TransactionByHashReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &transactionByHashReq); err != nil {
		panic(err)
	}
	transactionInfo := new(models.Transaction)
	db.Where("hash = ?", transactionByHashReq.Hash).Preload("Events").Preload("TransactionDetails").First(transactionInfo)
	c.Data["json"] = models.MakeTransactionResponse(transactionInfo)
	c.ServeJSON()
}

func (c *TransactionController) Transactions() {
	var transactionsReq models.TransactionsReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &transactionsReq); err != nil {
		panic(err)
	}
	transactions := make([]*models.Transaction, 0)
	db.Limit(transactionsReq.PageSize).Offset(transactionsReq.PageSize * transactionsReq.PageNo).Order("time desc").Find(&transactions)
	var transactionNum int64
	db.Model(&models.Transaction{}).Count(&transactionNum)
	c.Data["json"] = models.MakeTransactionsResponse(transactionsReq.PageSize, transactionsReq.PageNo,
		(int(transactionNum)+transactionsReq.PageSize-1)/transactionsReq.PageSize, int(transactionNum), transactions)
	c.ServeJSON()
}

func (c *TransactionController) TransactionsOfContract() {
	var transactionsOfContractReq models.TransactionsOfContractReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &transactionsOfContractReq); err != nil {
		panic(err)
	}
	transactions := make([]*models.Transaction, 0)
	db.Where("`to` = ?", transactionsOfContractReq.Contract).Limit(transactionsOfContractReq.PageSize).Offset(transactionsOfContractReq.PageSize * transactionsOfContractReq.PageNo).Order("time desc").Find(&transactions)
	var transactionNum int64
	db.Model(&models.Transaction{}).Where("`to` = ?", transactionsOfContractReq.Contract).Count(&transactionNum)
	c.Data["json"] = models.MakeTransactionsResponse(transactionsOfContractReq.PageSize, transactionsOfContractReq.PageNo,
		(int(transactionNum)+transactionsOfContractReq.PageSize-1)/transactionsOfContractReq.PageSize, int(transactionNum), transactions)
	c.ServeJSON()
}

func (c *TransactionController) TransactionsOfUser() {
	var transactionsOfUserReq models.TransactionsOfUserReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &transactionsOfUserReq); err != nil {
		panic(err)
	}
	transactions := make([]*models.Transaction, 0)
	db.Where("`from` = ?", transactionsOfUserReq.User).Limit(transactionsOfUserReq.PageSize).Offset(transactionsOfUserReq.PageSize * transactionsOfUserReq.PageNo).Order("time desc").
		Preload("Events").Preload("TransactionDetails").Find(&transactions)
	var transactionNum int64
	db.Model(&models.Transaction{}).Where("`from` = ?", transactionsOfUserReq.User).Count(&transactionNum)
	c.Data["json"] = models.MakeTransactionsResponse(transactionsOfUserReq.PageSize, transactionsOfUserReq.PageNo,
		(int(transactionNum)+transactionsOfUserReq.PageSize-1)/transactionsOfUserReq.PageSize, int(transactionNum), transactions)
	c.ServeJSON()
}

func (c *TransactionController) TransactionsDetails() {
	var transactionsReq models.TransactionDetailsOfTransactionReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &transactionsReq); err != nil {
		panic(err)
	}
	transactions := make([]*models.TransactionDetailWithInfo, 0)
	db.Where("transaction_hash = ?", transactionsReq.Hash).Preload("ContractInfo").Preload("NFTHolder").Limit(transactionsReq.PageSize).Offset(transactionsReq.PageSize * transactionsReq.PageNo).Order("time desc").Find(&transactions)
	var transactionNum int64
	db.Model(&models.TransactionDetailWithInfo{}).Where("transaction_hash = ?", transactionsReq.Hash).Count(&transactionNum)
	c.Data["json"] = models.MakeTransactionDetailsResponse(transactionsReq.PageSize, transactionsReq.PageNo,
		(int(transactionNum)+transactionsReq.PageSize-1)/transactionsReq.PageSize, int(transactionNum), transactions)
	c.ServeJSON()
}
