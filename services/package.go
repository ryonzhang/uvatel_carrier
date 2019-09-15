package services

import (
	"github.com/astaxie/beego/orm"
	"github.com/ryonzhang/uvatel_carrier/models"
)

func InsertPackage(request models.Package) (pack models.Package,err error) {
	o := orm.NewOrm()
	o.Using("default")
	_,err = o.Insert(&request)
	pack = request
	return
}

func GetPackage(request models.Package) (pack models.Package,err error){
	o := orm.NewOrm()
	o.Using("default")
	err = o.Read(&request, "name")
	pack=request
	return
}
func AddPackageToUser(pack models.Package,user models.User) (userpack models.UserPackage,err error){
	pack,err=GetPackage(pack)
	user,err=GetUser(user)
	if err!=nil{
		return
	}
	userpack = models.UserPackage{User:&user,Package:&pack,Total:pack.Total}
	o := orm.NewOrm()
	o.Using("default")
	_,err = o.Insert(&userpack)
	return
}