package main

import (
	_ "github.com/neoandroid/blackbird/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	_ "github.com/neoandroid/blackbird/tasks"
)

func init() {
	mysqlurl := beego.AppConfig.String("mysqluser") + ":" +  beego.AppConfig.String("mysqlpass") + "@tcp(" + beego.AppConfig.String("mysqlhost") + ")/" + beego.AppConfig.String("mysqldb")
	orm.RegisterDataBase("default", "mysql", mysqlurl)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

