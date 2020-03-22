package main

import (
	_ "Ems/models"
	_ "Ems/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
