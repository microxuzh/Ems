package controllers

import (
	"Ems/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Register() {

	o := orm.NewOrm()
	o.Using("default")
	user := models.User{Name: "Vivian"}
	err := o.Read(&user, "Name")
	if err == orm.ErrNoRows {
		id, err := o.Insert(&user)
		if err != nil {
			fmt.Printf("ID: %d, ERR: %v\n", id, err)
		}
	}

	c.TplName = "index.tpl"
}
