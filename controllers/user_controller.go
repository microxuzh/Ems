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
	user := models.Users{Name: "Tom"}
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)
	c.TplName = "index.tpl"
}
