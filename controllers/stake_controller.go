package controllers

import (
	"encoding/json"
	"explorer/models"
	"github.com/astaxie/beego"
)

type StakeController struct {
	beego.Controller
}

func (c *StakeController) StakesOfOwner() {
	var stakesOfOwnerReq models.StakesOfOwnerReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &stakesOfOwnerReq); err != nil {
		panic(err)
	}
	stakes := make([]*models.Stake, 0)
	db.Where("owner = ?", stakesOfOwnerReq.Owner).Limit(stakesOfOwnerReq.PageSize).Offset(stakesOfOwnerReq.PageSize * stakesOfOwnerReq.PageNo).Find(&stakes)
	var stakeNum int64
	db.Model(&models.Stake{}).Where("owner = ?", stakesOfOwnerReq.Owner).Count(&stakeNum)
	c.Data["json"] = models.MakeStakesOfOwnerResponse(stakesOfOwnerReq.PageSize, stakesOfOwnerReq.PageNo,
		(int(stakeNum) + stakesOfOwnerReq.PageSize - 1) / stakesOfOwnerReq.PageSize, int(stakeNum), stakes)
	c.ServeJSON()
}

func (c *StakeController) StakesOfValidator() {
	var stakesOfValidatorReq models.StakesOfValidatorReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &stakesOfValidatorReq); err != nil {
		panic(err)
	}
	stakes := make([]*models.Stake, 0)
	db.Where("validator = ?", stakesOfValidatorReq.Validator).Limit(stakesOfValidatorReq.PageSize).Offset(stakesOfValidatorReq.PageSize * stakesOfValidatorReq.PageNo).Find(&stakes)
	var stakeNum int64
	db.Model(&models.Stake{}).Where("validator = ?", stakesOfValidatorReq.Validator).Count(&stakeNum)
	c.Data["json"] = models.MakeStakesOfValidatorResponse(stakesOfValidatorReq.PageSize, stakesOfValidatorReq.PageNo,
		(int(stakeNum) + stakesOfValidatorReq.PageSize - 1) / stakesOfValidatorReq.PageSize, int(stakeNum), stakes)
	c.ServeJSON()
}

func (c *StakeController) StakeInfo() {
	var stakeInfoReq models.StakeInfoReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &stakeInfoReq); err != nil {
		panic(err)
	}
	stake := new(models.Stake)
	db.Where("owner = ? and validator = ?", stakeInfoReq.Owner, stakeInfoReq.Validator).First(stake)
	c.Data["json"] = models.MakeStakeInfoResponse(stake)
	c.ServeJSON()
}
