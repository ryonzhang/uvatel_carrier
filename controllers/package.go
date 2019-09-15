package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/ryonzhang/uvatel_carrier/models"
	"github.com/ryonzhang/uvatel_carrier/services"
	"strconv"
)

type PackageController struct {
	beego.Controller
}
// @Title Insert Package
// @Summary Insert package to database
// @Description Insert package
// @Param   name body    string  true    "name of the package"
// @Param   total body    int  true    "total amount"
// @Param   unit body    string  true    "unit"
// @Param   type body    int  true    "type"
// @Param   amount body    int  true    "amount"
// @Param   expiration body    int  true    "expiration"
// @Success 200 {object} ZDT.ZDTMisc.CmsResponse
// @Failure 400 Bad request
// @Failure 404 Not found
// @router /package/ [post]
func (c *PackageController) InsertPackage() {
	var request models.Package
	data := c.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &request)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	pack,err := services.InsertPackage(request)
	if err==nil{
		c.Data["json"] = &models.PackageResponse{Data:pack,Code:200,Error:models.Error{Message:""}}
	}else{
		c.Data["json"] = &models.PackageResponse{Data:models.Package{},Code:500,Error:models.Error{Message:err.Error()}}
	}
	c.ServeJSON()
}
// @Title Get Package
// @Summary Insert package to database
// @Description Get package from name
// @Param   name path    string  true    "name of the package"
// @Success 200 {object} ZDT.ZDTMisc.CmsResponse
// @Failure 400 Bad request
// @Failure 404 Not found
// @router /package/:name [get]
func (c *PackageController) GetPackage() {
	var request models.Package
	request.Name= c.Ctx.Input.Param(":name")
	pack,err := services.GetPackage(request)
	if err==nil{
		c.Data["json"] = &models.PackageResponse{Data:pack,Code:200,Error:models.Error{Message:""}}
	}else{
		c.Data["json"] = &models.PackageResponse{Data:models.Package{},Code:500,Error:models.Error{Message:err.Error()}}
	}
	c.ServeJSON()
}
// @Title Add package to certain user
// @Summary Add package to user
// @Description Add package to user
// @Param   name path    string  true    "name of the package"
// @Param   msisdn path    int  true    "ID of user"
// @Success 200 {object} ZDT.ZDTMisc.CmsResponse
// @Failure 400 Bad request
// @Failure 404 Not found
// @router /user/:msisdn/package/:name [post]
func (c *PackageController) AddPackageToUser() {
	var pack models.Package
	var user models.User

	pack.Name= c.Ctx.Input.Param(":name")
	user.Msisdn,_ = strconv.Atoi(c.Ctx.Input.Param(":msisdn"))
	userpack,err := services.AddPackageToUser(pack,user)
	if err==nil{
		c.Data["json"] = &models.UserPackageResponse{Data:userpack,Code:200,Error:models.Error{Message:""}}
	}else{
		c.Data["json"] = &models.UserPackageResponse{Data:models.UserPackage{},Code:500,Error:models.Error{Message:err.Error()}}
	}
	c.ServeJSON()
}
