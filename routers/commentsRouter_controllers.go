package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/ryonzhang/uvatel_carrier/controllers:BalanceController"] = append(beego.GlobalControllerRouter["github.com/ryonzhang/uvatel_carrier/controllers:BalanceController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ryonzhang/uvatel_carrier/controllers:BalanceController"] = append(beego.GlobalControllerRouter["github.com/ryonzhang/uvatel_carrier/controllers:BalanceController"],
        beego.ControllerComments{
            Method: "InsertAccount",
            Router: `/account/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ryonzhang/uvatel_carrier/controllers:BalanceController"] = append(beego.GlobalControllerRouter["github.com/ryonzhang/uvatel_carrier/controllers:BalanceController"],
        beego.ControllerComments{
            Method: "GetBalance",
            Router: `/balance/:msisdn`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ryonzhang/uvatel_carrier/controllers:BalanceController"] = append(beego.GlobalControllerRouter["github.com/ryonzhang/uvatel_carrier/controllers:BalanceController"],
        beego.ControllerComments{
            Method: "AddBalance",
            Router: `/balance/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ryonzhang/uvatel_carrier/controllers:BalanceController"] = append(beego.GlobalControllerRouter["github.com/ryonzhang/uvatel_carrier/controllers:BalanceController"],
        beego.ControllerComments{
            Method: "DeductBalance",
            Router: `/balance/deduct`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ryonzhang/uvatel_carrier/controllers:BalanceController"] = append(beego.GlobalControllerRouter["github.com/ryonzhang/uvatel_carrier/controllers:BalanceController"],
        beego.ControllerComments{
            Method: "SetBalance",
            Router: `/balance/set`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ryonzhang/uvatel_carrier/controllers:PackageController"] = append(beego.GlobalControllerRouter["github.com/ryonzhang/uvatel_carrier/controllers:PackageController"],
        beego.ControllerComments{
            Method: "InsertPackage",
            Router: `/package/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ryonzhang/uvatel_carrier/controllers:PackageController"] = append(beego.GlobalControllerRouter["github.com/ryonzhang/uvatel_carrier/controllers:PackageController"],
        beego.ControllerComments{
            Method: "GetPackage",
            Router: `/package/:name`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ryonzhang/uvatel_carrier/controllers:PackageController"] = append(beego.GlobalControllerRouter["github.com/ryonzhang/uvatel_carrier/controllers:PackageController"],
        beego.ControllerComments{
            Method: "AddPackageToUser",
            Router: `/user/:msisdn/package/:name`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ryonzhang/uvatel_carrier/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ryonzhang/uvatel_carrier/controllers:UserController"],
        beego.ControllerComments{
            Method: "InsertUser",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ryonzhang/uvatel_carrier/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ryonzhang/uvatel_carrier/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUser",
            Router: `/:msisdn/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
