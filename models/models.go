package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)
// Account
type Account struct{
	Id int `json:"id"`
	Msisdn	int `json:"msisdn"`
	Active	bool `json:"active"`
	CoreBalance	float32	`json:"core_balance"`
	Currency string `json:"currency"`
}
func (a *Account) TableIndex() [][]string {
	return [][]string{
		[]string{"Msisdn", "Active"},
	}
}
func (a *Account) TableUnique() [][]string {
	return [][]string{
		[]string{"Msisdn", "Active"},
	}
}
// User
type User struct {
	Id int `json:"id"`
	Msisdn	int `json:"msisdn"`
	Active	bool `json:"active"`
	CountryCode	string	`json:"country_code"`
	CreatedAt time.Time `orm:"auto_now_add" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now" json:"updated_at"`
	Imsi    int `json:"imsi"`
	ExpiredAt time.Time `orm:"auto_now" json:"expired_at"`
}
func (u *User) TableIndex() [][]string {
	return [][]string{
		[]string{"Msisdn", "Active"},
	}
}
func (u *User) TableUnique() [][]string {
	return [][]string{
		[]string{"Msisdn", "Active"},
	}
}
type Package struct{
	Id int `json:"id"`
	Name string `orm:"name"`
	Total float32 `json:"total"`
	Unit	string `json:"unit"`
	Type	int `json:"type"`
	Amount	float32 `json:"amount"`
	Expiration int `json:"expiration"`
}

type UserPackage struct{
	Id int `json:"id"`
	Total	float32 `json:"total"`
	Used	bool `json:"used"`
	Remaining 	float32 `json:"remaining"`
	Unit 	string `json:"unit"`
	ExpiredAt	time.Time `orm:"auto_now" json:"expired_at"`
	User	*User `orm:"rel(fk);null;on_delete(set_null)" json:"user"`
	Package	*Package	`orm:"rel(fk)" json:"package"`
}

func init() {
	orm.RegisterModel(new(User),new(Package),new(UserPackage),new(Account))
}



