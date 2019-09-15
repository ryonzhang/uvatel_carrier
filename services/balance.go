package services

import (
	"github.com/astaxie/beego/orm"
	"github.com/ryonzhang/uvatel_carrier/models"
)

func  DeductBalance(request models.BalanceRequest) (account models.Account,err error){
	o := orm.NewOrm()
	o.Using("default")
	account = models.Account{Msisdn: request.Msisdn,Active:true}
	err = o.Read(&account, "msisdn","active")
	if err == nil {
		account.CoreBalance = account.CoreBalance-float32(request.Amount)
		_, err = o.Update(&account)
	}
	return
}
func AddBalance(request models.BalanceRequest)(account models.Account,err error) {
	o := orm.NewOrm()
	o.Using("default")
	account = models.Account{Msisdn: request.Msisdn,Active:true}
	err = o.Read(&account, "msisdn","active")
	if err == nil {
		account.CoreBalance = account.CoreBalance+float32(request.Amount)
		_, err = o.Update(&account)
	}
	return
}
func SetBalance(request models.BalanceRequest) (account models.Account,err error){
	o := orm.NewOrm()
	o.Using("default")
	account = models.Account{Msisdn: request.Msisdn,Active:true}
	err = o.Read(&account, "msisdn","active")
	if err == nil {
		account.CoreBalance = float32(request.Amount)
		_, err = o.Update(&account)
	}
	return
}
func GetBalance(request models.BalanceRequest) (account models.Account,err error){
	o := orm.NewOrm()
	o.Using("default")
	account = models.Account{Msisdn: request.Msisdn,Active:true}
	err = o.Read(&account, "msisdn","active")
	return
}
func InsertAccount(request models.BalanceRequest) (account models.Account,err error) {
	o := orm.NewOrm()
	o.Using("default")
	account = models.Account{Msisdn: request.Msisdn,Active:true,CoreBalance:float32(request.Amount)}
	_,err = o.Insert(&account)
	return
}
