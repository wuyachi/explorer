package controllers

import (
	"encoding/json"
	"explorer/models"
	"github.com/astaxie/beego"
)

type ProposeController struct {
	beego.Controller
}

func (c *ProposeController) Proposes() {
	var proposesReq models.ProposesReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &proposesReq); err != nil {
		panic(err)
	}
	proposes := make([]*models.Propose, 0)
	db.Limit(proposesReq.PageSize).Offset(proposesReq.PageSize * proposesReq.PageNo).Find(&proposes)
	var proposeNum int64
	db.Model(&models.Propose{}).Count(&proposeNum)
	c.Data["json"] = models.MakeProposesResponse(proposesReq.PageSize, proposesReq.PageNo,
		(int(proposeNum) + proposesReq.PageSize - 1) / proposesReq.PageSize, int(proposeNum), proposes)
	c.ServeJSON()
}


func (c *ProposeController) ProposeInfo() {
	var proposeInfoReq models.ProposeInfoReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &proposeInfoReq); err != nil {
		panic(err)
	}
	propose := new(models.Propose)
	db.Where("propose_id = ?", proposeInfoReq.ProposeId).First(propose)
	c.Data["json"] = models.MakeProposeInfoResponse(propose)
	c.ServeJSON()
}

func (c *ProposeController) ProposesOfUser() {
	var proposesOfUserReq models.ProposesOfUserReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &proposesOfUserReq); err != nil {
		panic(err)
	}
	proposes := make([]*models.Propose, 0)
	db.Where("proposer = ?", proposesOfUserReq.Proposer).Limit(proposesOfUserReq.PageSize).Offset(proposesOfUserReq.PageSize * proposesOfUserReq.PageNo).Find(&proposes)
	var validatorNum int64
	db.Model(&models.Propose{}).Where("proposer = ?", proposesOfUserReq.Proposer).Count(&validatorNum)
	c.Data["json"] = models.MakeProposesOfUserResponse(proposesOfUserReq.PageSize, proposesOfUserReq.PageNo,
		(int(validatorNum) + proposesOfUserReq.PageSize - 1) / proposesOfUserReq.PageSize, int(validatorNum), proposes)
	c.ServeJSON()
}
