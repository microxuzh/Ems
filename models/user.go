package models

import "github.com/astaxie/beego/orm"

import "github.com/astaxie/beego"

type Users struct {
	Id   int
	Name string
}

func init() {
	orm.RegisterModel(new(Users))
	orm.RunSyncdb("default", false, true)
	beego.Info("Register Users Model")
}
