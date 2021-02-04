package controllers

import (
	"encoding/json"
	"explorer/models"
	"github.com/astaxie/beego"
)

type PLTContractController struct {
	beego.Controller
}

func (c *PLTContractController) PLTHolderInfo() {
	var pltHolderInfoReq models.PLTHolderInfoReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &pltHolderInfoReq); err != nil {
		panic(err)
	}
	pltContract := new(models.PLTContract)
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
	pltContracts := make([]*models.PLTContract, 0)
	db.Limit(pltHoldersReq.PageSize).Offset(pltHoldersReq.PageSize * pltHoldersReq.PageNo).Order("amount desc").Find(&pltContracts)
	var pltContractNum int64
	db.Model(&models.PLTContract{}).Count(&pltContractNum)
	c.Data["json"] = models.MakePLTHoldersResponse(pltHoldersReq.PageSize, pltHoldersReq.PageNo,
		(int(pltContractNum) + pltHoldersReq.PageSize - 1) / pltHoldersReq.PageSize, int(pltContractNum), pltContracts)
	c.ServeJSON()
}
