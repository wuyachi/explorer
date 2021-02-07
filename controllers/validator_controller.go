package controllers

import (
	"encoding/json"
	"explorer/models"
	"github.com/astaxie/beego"
)

type ValidatorController struct {
	beego.Controller
}

func (c *ValidatorController) Validators() {
	var validatorsReq models.ValidatorsReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &validatorsReq); err != nil {
		panic(err)
	}
	validators := make([]*models.ValidatorWithPercent, 0)
	db.Table("(?) as u, validators", db.Model(&models.ValidatorWithPercent{}).Select("sum(stake_amount) as total")).
		Select("address, delegate_factor, stake_amount, stake_amount/total as percent, name, uri").
		Limit(validatorsReq.PageSize).Offset(validatorsReq.PageSize * validatorsReq.PageNo).Order("stake_amount desc").Find(&validators)

	var validatorNum int64
	db.Model(&models.ValidatorWithPercent{}).Count(&validatorNum)
	c.Data["json"] = models.MakeValidatorsResponse(validatorsReq.PageSize, validatorsReq.PageNo,
		(int(validatorNum)+validatorsReq.PageSize-1)/validatorsReq.PageSize, int(validatorNum), validators)
	c.ServeJSON()
}

func (c *ValidatorController) ValidatorInfo() {
	var validatorInfoReq models.ValidatorInfoReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &validatorInfoReq); err != nil {
		panic(err)
	}
	validator := new(models.ValidatorWithPercent)
	db.Where("address = ?", validatorInfoReq.Address).First(validator)
	c.Data["json"] = models.MakeValidatorInfoResponse(validator)
	c.ServeJSON()
}
