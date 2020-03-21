package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	beego.Info("Initialize...")
	driverName := beego.AppConfig.String("dbDriver")
	//注册数据库驱动
	orm.RegisterDriver(driverName, orm.DRMySQL)

	//数据库连接参数
	dbUser := beego.AppConfig.String("dbUser")
	dbPassword := beego.AppConfig.String("dbPassword")
	dbHost := beego.AppConfig.String("dbHost")
	dbPort := beego.AppConfig.String("dbPort")
	dbName := beego.AppConfig.String("dbName")

	//数据库连接串
	dbConn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"

	//注册数据库
	err := orm.RegisterDataBase("default", driverName, dbConn)
	if err != nil {
		beego.Info("连接数据库出错！")
		return
	}
	beego.Info("连接数据库成功!")
	//注册Model
	// orm.RegisterModel(new(Users))
	//开启orm调试模式；开发过程中打开，release时需要关闭
	orm.Debug = true
	//自动建表
	// orm.RunSyncdb("default", true, true)

}
