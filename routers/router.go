package routers

import (
	"github.com/ryonzhang/uvatel_carrier/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.BalanceController{},"get:Get")
	beego.Router("/balance/deduct", &controllers.BalanceController{},"post:DeductBalance")
	beego.Router("/balance/add", &controllers.BalanceController{}, "post:AddBalance")
	beego.Router("/balance/set", &controllers.BalanceController{}, "post:SetBalance")
	beego.Router("/balance/:msisdn", &controllers.BalanceController{}, "get:GetBalance")
	beego.Router("/account/", &controllers.BalanceController{}, "post:InsertAccount")

	beego.Router("/user/", &controllers.UserController{}, "post:InsertUser")
	beego.Router("/user/:msisdn", &controllers.UserController{}, "get:GetUser")

	beego.Router("/package/", &controllers.PackageController{}, "post:InsertPackage")
	beego.Router("/package/:name", &controllers.PackageController{}, "get:GetPackage")
	beego.Router("/packages/:msisdn", &controllers.PackageController{}, "get:GetPackages")
	beego.Router("/user/:msisdn/package/:name", &controllers.PackageController{}, "post:AddPackageToUser")

	ns :=beego.NewNamespace("/v1/",
		beego.NSInclude(
			&controllers.BalanceController{},
		),
		beego.NSInclude(
			&controllers.PackageController{},
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
