package services

import (
	"github.com/astaxie/beego/orm"
	"github.com/ryonzhang/uvatel_carrier/models"
)

func InsertUser(request models.User) (user models.User,err error) {
	o := orm.NewOrm()
	o.Using("default")
	_,err = o.Insert(&request)
	user = request
	return
}

func GetUser(request models.User) (user models.User,err error){
	o := orm.NewOrm()
	o.Using("default")
	user = models.User{Msisdn: request.Msisdn,Active:true}
	err = o.Read(&user, "msisdn","active")
	return
}

