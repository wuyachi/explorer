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
		beego.NSRouter("/transactiondetails/", &controllers.TransactionController{}, "post:TransactionsDetails"),
		beego.NSRouter("/plt/", &controllers.PLTContractController{}, "post:PLT"),
		beego.NSRouter("/pltholderinfo/", &controllers.PLTContractController{}, "post:PLTHolderInfo"),
		beego.NSRouter("/pltholders/", &controllers.PLTContractController{}, "post:PLTHolders"),
		beego.NSRouter("/plttransactions/", &controllers.PLTContractController{}, "post:PLTTransactions"),
		beego.NSRouter("/plttransactionsofuser/", &controllers.PLTContractController{}, "post:PLTTransactionsOfUser"),
		beego.NSRouter("/nftholdersofuser/", &controllers.NFTContractController{}, "post:NFTHoldersOfUser"),
		beego.NSRouter("/allnftsholdersofuser/", &controllers.NFTContractController{}, "post:AllNFTsHoldersOfUser"),
		beego.NSRouter("/nfttokeninfo/", &controllers.NFTContractController{}, "post:NFTTokenInfo"),
		beego.NSRouter("/nfttokentransactions/", &controllers.NFTContractController{}, "post:NFTTokenTransactions"),
		beego.NSRouter("/nfts/", &controllers.NFTContractController{}, "post:NFTs"),
		beego.NSRouter("/nft/", &controllers.NFTContractController{}, "post:NFT"),
		beego.NSRouter("/nftholders/", &controllers.NFTContractController{}, "post:NFTHolders"),
		beego.NSRouter("/nftusers/", &controllers.NFTContractController{}, "post:NFTUsers"),
		beego.NSRouter("/nfttransactions/", &controllers.NFTContractController{}, "post:NFTTransactions"),
		beego.NSRouter("/nfttransactionsofuser/", &controllers.NFTContractController{}, "post:NFTTransactionsOfUser"),
		beego.NSRouter("/allnftstransactionsofuser/", &controllers.NFTContractController{}, "post:AllNFTsTransactionsOfUser"),
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
