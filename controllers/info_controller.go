package controllers

import (
	"explorer/models"
	"github.com/astaxie/beego"
)

type InfoController struct {
	beego.Controller
}

func (c *InfoController) Get() {
	explorer := &models.ExplorerResp{
		Version: "v1",
		URL:     "http://localhost:8080/v1",
	}
	c.Data["json"] = explorer
	c.ServeJSON()
}
