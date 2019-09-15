package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/ryonzhang/uvatel_carrier/models"
	"github.com/ryonzhang/uvatel_carrier/services"
	"strconv"
)

type UserController struct {
	beego.Controller
}
// @Title Insert user into db
// @Summary Add user into db
// @Description Add user
// @Param   msisdn body    int  true    "ID of user"
// @Param   active body    boolean  true    "active"
// @Param   country_code body    string  true    "active"
// @Param   imsi body    string  true    "active"
// @Param   expired_at body    datetime  true    "active"
// @Success 200 {object} ZDT.ZDTMisc.CmsResponse
// @Failure 400 Bad request
// @Failure 404 Not found
// @router / [post]
func (c *UserController) InsertUser() {
	var request models.User
	data := c.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &request)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	user,err := services.InsertUser(request)
	if err==nil{
		c.Data["json"] = &models.UserResponse{Data:user,Code:200,Error:models.Error{Message:""}}
	}else{
		c.Data["json"] = &models.UserResponse{Data:models.User{},Code:500,Error:models.Error{Message:err.Error()}}
	}
	c.ServeJSON()
}
// @Title Get user from msisdn
// @Summary Get user from msisdn
// @Description Get user from msisdn
// @Param   msisdn path    int  true    "ID of user"
// @Success 200 {object} ZDT.ZDTMisc.CmsResponse
// @Failure 400 Bad request
// @Failure 404 Not found
// @router /:msisdn/ [get]
func (c *UserController) GetUser() {
	var request models.User
	request.Msisdn,_ = strconv.Atoi(c.Ctx.Input.Param(":msisdn"))
	user,err := services.GetUser(request)
	if err==nil{
		c.Data["json"] = &models.UserResponse{Data:user,Code:200,Error:models.Error{Message:""}}
	}else{
		c.Data["json"] = &models.UserResponse{Data:models.User{},Code:500,Error:models.Error{Message:err.Error()}}
	}
	c.ServeJSON()
}
