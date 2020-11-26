package controllers

import (
	"encoding/json"
	"explorer/models"
	"github.com/astaxie/beego"
)

type NFTContractController struct {
	beego.Controller
}

func (c *NFTContractController) NFTs() {
	var nftsReq models.NFTsReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &nftsReq); err != nil {
		panic(err)
	}
	db := newDB()
	nftContracts := make([]*models.NFTContract, 0)
	db.Limit(nftsReq.PageSize).Offset(nftsReq.PageSize * nftsReq.PageNo).Distinct("nft").Find(&nftContracts)
	var nftContractNum int64
	db.Model(&models.NFTContract{}).Distinct("nft").Count(&nftContractNum)
	c.Data["json"] = models.MakeNFTsResponse(nftsReq.PageSize, nftsReq.PageNo,
		(int(nftContractNum) + nftsReq.PageSize - 1) / nftsReq.PageSize, int(nftContractNum), nftContracts)
	c.ServeJSON()
}

func (c *NFTContractController) NFTTokens() {
	var nftTokensReq models.NFTTokensReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &nftTokensReq); err != nil {
		panic(err)
	}
	db := newDB()
	nftContracts := make([]*models.NFTContract, 0)
	db.Where("nft = ?", nftTokensReq.NFT).Limit(nftTokensReq.PageSize).Offset(nftTokensReq.PageSize * nftTokensReq.PageNo).Distinct("nft", "token").Find(&nftContracts)
	var nftContractNum int64
	db.Model(&models.NFTContract{}).Where("nft = ?", nftTokensReq.NFT).Distinct("nft", "token").Count(&nftContractNum)
	c.Data["json"] = models.MakeNFTTokensResponse(nftTokensReq.PageSize, nftTokensReq.PageNo,
		(int(nftContractNum) + nftTokensReq.PageSize - 1) / nftTokensReq.PageSize, int(nftContractNum), nftContracts)
	c.ServeJSON()
}

func (c *NFTContractController) NFTTokenInfo() {
	var nftTokenInfoReq models.NFTTokenInfoReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &nftTokenInfoReq); err != nil {
		panic(err)
	}
	db := newDB()
	nftContract := new(models.NFTContract)
	db.Where("nft = ? and token = ?", nftTokenInfoReq.NFT, nftTokenInfoReq.Token).First(nftContract)
	c.Data["json"] = models.MakeNFTTokenInfoResponse(nftContract)
	c.ServeJSON()
}

func (c *NFTContractController) NFTHoldersOfUser() {
	var nftHoldersOfUserReq models.NFTHoldersOfUserReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &nftHoldersOfUserReq); err != nil {
		panic(err)
	}
	db := newDB()
	nftContracts := make([]*models.NFTContract, 0)
	db.Where("nft = ? and owner = ?", nftHoldersOfUserReq.NFT, nftHoldersOfUserReq.Address).Limit(nftHoldersOfUserReq.PageSize).Offset(nftHoldersOfUserReq.PageSize * nftHoldersOfUserReq.PageNo).Find(&nftContracts)
	var nftContractNum int64
	db.Model(&models.NFTContract{}).Where("nft = ? and owner = ?", nftHoldersOfUserReq.NFT, nftHoldersOfUserReq.Address).Count(&nftContractNum)
	c.Data["json"] = models.MakeNFTHoldersOfUserResponse(nftHoldersOfUserReq.PageSize, nftHoldersOfUserReq.PageNo,
		(int(nftContractNum) + nftHoldersOfUserReq.PageSize - 1) / nftHoldersOfUserReq.PageSize, int(nftContractNum), nftContracts)
	c.ServeJSON()
}

func (c *NFTContractController) NFTHolders() {
	var nftHoldersReq models.NFTHoldersReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &nftHoldersReq); err != nil {
		panic(err)
	}
	db := newDB()
	nftContracts := make([]*models.NFTContract, 0)
	db.Where("nft = ?", nftHoldersReq.NFT).Limit(nftHoldersReq.PageSize).Offset(nftHoldersReq.PageSize * nftHoldersReq.PageNo).Find(&nftContracts)
	var nftContractNum int64
	db.Model(&models.NFTContract{}).Where("nft = ?", nftHoldersReq.NFT).Count(&nftContractNum)
	c.Data["json"] = models.MakeNFTHoldersResponse(nftHoldersReq.PageSize, nftHoldersReq.PageNo,
		(int(nftContractNum) + nftHoldersReq.PageSize - 1) / nftHoldersReq.PageSize, int(nftContractNum), nftContracts)
	c.ServeJSON()
}
