package common

import (
	"os"

	"github.com/astaxie/beego"

	_ "mysql"

	"github.com/go-xorm/xorm"
)

var configFile string = "conf/app.conf"
var Engine *xorm.Engine

func init() {
	var err error
	Engine, err = xorm.NewEngine("mysql", getMysqlUrl())
	if showsql() == "true" {
		Engine.ShowSQL(true)
	}
	if err != nil {
		os.Exit(100)
	}
}

/**
**通过配置文件获取数据库连接
**/
func getMysqlUrl() string {

	url := beego.AppConfig.String("url")
	return url
}

/**
**通过配置文件获取数据库连接
**/
func showsql() string {
	showsql := beego.AppConfig.String("showsql")
	return showsql
}
