package controllers

import (
	"explorer/models"
	"github.com/astaxie/beego"
)

type ChainController struct {
	beego.Controller
}

func (c *ChainController) Chain() {
	db := newDB()
	chain := new(models.Chain)
	db.First(chain)
	c.Data["json"] = models.MakeChainResponse(chain)
	c.ServeJSON()
}
