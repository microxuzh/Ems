package controllers

import (
	"Ems/models"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Register() {

	// o := orm.NewOrm()
	// o.Using("default")
	// user := models.User{UserName: "MIKE", Level: models.Admin}
	// err := o.Read(&user, "Name")
	// if err == orm.ErrNoRows {
	// 	id, err := o.Insert(&user)
	// 	if err != nil {
	// 		fmt.Printf("ID: %d, ERR: %v\n", id, err)
	// 	}
	// }
	models.PostUser()
	c.TplName = "index.tpl"
}
