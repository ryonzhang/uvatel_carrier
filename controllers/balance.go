package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/ryonzhang/uvatel_carrier/models"
	"github.com/ryonzhang/uvatel_carrier/services"
	"strconv"
)

type BalanceController struct {
	beego.Controller
}
// @Title Get
// @Summary URL for playing around
// @Description URL for trial and testing new features
// @Success 200 {object} ZDT.ZDTMisc.CmsResponse
// @Failure 400 Bad request
// @Failure 404 Not found
// @router / [get]
func (c *BalanceController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["json"] = "astaxie@gmail.com"
	c.ServeJSON()
}
// @Title Deduct balance
// @Summary Deduct the balance from account
// @Description
// @Param   msisdn body    int  true    "The ID of the caller"
// @Param   amount body    decimal  true    "The amount to deduct"
// @Success 200 {} ZDT.ZDTMisc.CmsResponse
// @Failure 400 Bad request
// @Failure 404 Not found
// @router /balance/deduct [post]
func (c *BalanceController) DeductBalance() {
	var request models.BalanceRequest
	data := c.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &request)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	account,err := services.DeductBalance(request)
	if err==nil{
		c.Data["json"] = &models.BalanceResponse{Data:account,Code:200,Error:models.Error{Message:""}}
	}else{
		c.Data["json"] = &models.BalanceResponse{Data:models.Account{},Code:500,Error:models.Error{Message:err.Error()}}
	}
	c.ServeJSON()
}
// @Title Add balance
// @Summary Add the balance to account
// @Description
// @Param   msisdn body    int  true    "The ID of the caller"
// @Param   amount body    decimal  true    "The amount to add"
// @Success 200 {} ZDT.ZDTMisc.CmsResponse
// @Failure 400 Bad request
// @Failure 404 Not found
// @router /balance/add [post]
func (c *BalanceController) AddBalance() {
	var request models.BalanceRequest
	data := c.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &request)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	account,err := services.AddBalance(request)
	if err==nil{
		c.Data["json"] = &models.BalanceResponse{Data:account,Code:200,Error:models.Error{Message:""}}
	}else{
		c.Data["json"] = &models.BalanceResponse{Data:models.Account{},Code:500,Error:models.Error{Message:err.Error()}}
	}
	c.ServeJSON()
}
// @Title Set balance
// @Summary Set the balance to account
// @Description
// @Param   msisdn body    int  true    "The ID of the caller"
// @Param   amount body    decimal  true    "The amount to set"
// @Success 200 {} ZDT.ZDTMisc.CmsResponse
// @Failure 400 Bad request
// @Failure 404 Not found
// @router /balance/set [post]
func (c *BalanceController) SetBalance() {
	var request models.BalanceRequest
	data := c.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &request)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	account,err := services.SetBalance(request)
	if err==nil{
		c.Data["json"] = &models.BalanceResponse{Data:account,Code:200,Error:models.Error{Message:""}}
	}else{
		c.Data["json"] = &models.BalanceResponse{Data:models.Account{},Code:500,Error:models.Error{Message:err.Error()}}
	}
	c.ServeJSON()
}
// @Title Get balance
// @Summary Get the balance of account
// @Description
// @Param   msisdn path    string  true    "The ID of the caller"
// @Success 200 {} ZDT.ZDTMisc.CmsResponse
// @Failure 400 Bad request
// @Failure 404 Not found
// @router /balance/:msisdn [get]
func (c *BalanceController) GetBalance() {
	var request models.BalanceRequest
	request.Msisdn,_ = strconv.Atoi(c.Ctx.Input.Param(":msisdn"))
	account,err := services.GetBalance(request)
	if err==nil{
		c.Data["json"] = &models.BalanceResponse{Data:account,Code:200,Error:models.Error{Message:""}}
	}else{
		c.Data["json"] = &models.BalanceResponse{Data:models.Account{},Code:500,Error:models.Error{Message:err.Error()}}
	}
	c.ServeJSON()
}
// @Title Insert account
// @Summary Insert new row of account
// @Description
// @Param   msisdn body    int  true    "The ID of the caller"
// @Param   amount body    decimal  true    "The amount to set"
// @Success 200 {} ZDT.ZDTMisc.CmsResponse
// @Failure 400 Bad request
// @Failure 404 Not found
// @router /account/ [post]
func (c *BalanceController) InsertAccount() {
	var request models.BalanceRequest
	data := c.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &request)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	account,err := services.InsertAccount(request)
	if err==nil{
		c.Data["json"] = &models.BalanceResponse{Data:account,Code:200,Error:models.Error{Message:""}}
	}else{
		c.Data["json"] = &models.BalanceResponse{Data:models.Account{},Code:500,Error:models.Error{Message:err.Error()}}
	}
	c.ServeJSON()
}


