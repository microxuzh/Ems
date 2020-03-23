package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type AccessLevel int32

const (
	guest AccessLevel = iota
	operator
	manager
	admin
	superAdmin
)

type User struct {
	Id       int
	UserName string
	Name     string
	Password string
	Email    string
	Active   bool
	Level    AccessLevel
}

func init() {
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
}

func PostUser() {
	o := orm.NewOrm()
	user := User{UserName: "Tom", Name: "李三", Level: superAdmin}
	err := o.Read(&user, "UserName")
	if err == orm.ErrNoRows {
		id, err := o.Insert(&user)
		if err != nil {
			fmt.Printf("ID: %d, ERR: %v\n", id, err)
		}
	}
}
