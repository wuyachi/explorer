package controllers

import (
	"explorer/models"
	"github.com/astaxie/beego"
)

type ChainController struct {
	beego.Controller
}

func (c *ChainController) Chain() {
	chain := new(models.Chain)
	db.First(chain)
	var transactionNum int64
	db.Model(&models.Transaction{}).Count(&transactionNum)
	var firstBlock models.Block
	db.Where("number = ?", 1).First(&firstBlock)
	var latestBlock models.Block
	db.Where("number = ?", chain.Height).First(&latestBlock)
	c.Data["json"] = models.MakeChainResponse(chain, uint64(transactionNum), latestBlock.Time - firstBlock.Time, latestBlock.Number - firstBlock.Number)
	c.ServeJSON()
}
