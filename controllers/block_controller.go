package controllers

import (
	"encoding/json"
	"explorer/models"
	"github.com/astaxie/beego"
)

type BlockController struct {
	beego.Controller
}

func (c *BlockController) BlockByNumber() {
	var blockByNumberReq models.BlockByNumberReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &blockByNumberReq); err != nil {
		panic(err)
	}
	blockInfo := new(models.Block)
	db.Where("number = ?", blockByNumberReq.Number).Preload("Transactions").First(blockInfo)
	c.Data["json"] = models.MakeBlockResponse(blockInfo)
	c.ServeJSON()
}

func (c *BlockController) BlockByHash() {
	var blockByHashReq models.BlockByHashReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &blockByHashReq); err != nil {
		panic(err)
	}
	blockInfo := new(models.Block)
	db.Where("hash = ?", blockByHashReq.Hash).Preload("Transactions").First(blockInfo)
	c.Data["json"] = models.MakeBlockResponse(blockInfo)
	c.ServeJSON()
}

func (c *BlockController) Blocks() {
	var blocksReq models.BlocksReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &blocksReq); err != nil {
		panic(err)
	}
	blocks := make([]*models.Block, 0)
	db.Debug().Limit(blocksReq.PageSize).Offset(blocksReq.PageSize * blocksReq.PageNo).Order("number desc").Preload("Transactions").Find(&blocks)
	var blockNum int64
	db.Model(&models.Block{}).Count(&blockNum)
	c.Data["json"] = models.MakeBlocksResponse(blocksReq.PageSize, blocksReq.PageNo, (int(blockNum) + blocksReq.PageSize - 1) / blocksReq.PageSize, int(blockNum), blocks)
	c.ServeJSON()
}
