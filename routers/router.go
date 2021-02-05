package routers

import (
	"explorer/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/", &controllers.InfoController{}, "*:Get"),
		beego.NSRouter("/chain", &controllers.ChainController{}, "get:Chain"),
		beego.NSRouter("/blockbynumber/", &controllers.BlockController{}, "post:BlockByNumber"),
		beego.NSRouter("/blockbyhash/", &controllers.BlockController{}, "post:BlockByHash"),
		beego.NSRouter("/blocks/", &controllers.BlockController{}, "post:Blocks"),
		beego.NSRouter("/transactionbyhash/", &controllers.TransactionController{}, "post:TransactionByHash"),
		beego.NSRouter("/transactions/", &controllers.TransactionController{}, "post:Transactions"),
		beego.NSRouter("/transactionsofcontract/", &controllers.TransactionController{}, "post:TransactionsOfContract"),
		beego.NSRouter("/transactionsofuser/", &controllers.TransactionController{}, "post:TransactionsOfUser"),
		beego.NSRouter("/pltholderinfo/", &controllers.PLTContractController{}, "post:PLTHolderInfo"),
		beego.NSRouter("/pltholders/", &controllers.PLTContractController{}, "post:PLTHolders"),
		beego.NSRouter("/nftholdersofuser/", &controllers.NFTContractController{}, "post:NFTHoldersOfUser"),
		beego.NSRouter("/nfttokeninfo/", &controllers.NFTContractController{}, "post:NFTTokenInfo"),
		beego.NSRouter("/nfttokens/", &controllers.NFTContractController{}, "post:NFTTokens"),
		beego.NSRouter("/nfts/", &controllers.NFTContractController{}, "post:NFTs"),
		beego.NSRouter("/nftholders/", &controllers.NFTContractController{}, "post:NFTHolders"),
		beego.NSRouter("/stakesofowner/", &controllers.StakeController{}, "post:StakesOfOwner"),
		beego.NSRouter("/stakesofvalidator/", &controllers.StakeController{}, "post:StakesOfValidator"),
		beego.NSRouter("/stakeinfo/", &controllers.StakeController{}, "post:StakeInfo"),
		beego.NSRouter("/validators/", &controllers.ValidatorController{}, "post:Validators"),
		beego.NSRouter("/validatorinfo/", &controllers.ValidatorController{}, "post:ValidatorInfo"),
		beego.NSRouter("/proposes/", &controllers.ProposeController{}, "post:Proposes"),
		beego.NSRouter("/proposeinfo/", &controllers.ProposeController{}, "post:ProposeInfo"),
		beego.NSRouter("/proposesofuser/", &controllers.ProposeController{}, "post:ProposesOfUser"),
	)
	beego.AddNamespace(ns)
	beego.Router("/", &controllers.InfoController{}, "*:Get")
}
