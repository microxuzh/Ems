package routers

import (
	"Ems/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/Register", &controllers.UserController{}, "Post:Register")
}
