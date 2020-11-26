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
	db := newDB()
	validators := make([]*models.Validator, 0)
	db.Limit(validatorsReq.PageSize).Offset(validatorsReq.PageSize * validatorsReq.PageNo).Find(&validators)
	var validatorNum int64
	db.Model(&models.Validator{}).Count(&validatorNum)
	c.Data["json"] = models.MakeValidatorsResponse(validatorsReq.PageSize, validatorsReq.PageNo,
		(int(validatorNum) + validatorsReq.PageSize - 1) / validatorsReq.PageSize, int(validatorNum), validators)
	c.ServeJSON()
}


func (c *ValidatorController) ValidatorInfo() {
	var validatorInfoReq models.ValidatorInfoReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &validatorInfoReq); err != nil {
		panic(err)
	}
	db := newDB()
	validator := new(models.Validator)
	db.Where("address = ?", validatorInfoReq.Address).First(validator)
	c.Data["json"] = models.MakeValidatorInfoResponse(validator)
	c.ServeJSON()
}
