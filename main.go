package main

import (
	"net/http"
	"strings"
	_ "win-xp/routers"

	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
)

func TransparentStatic(ctx *context.Context) {
	//请求controller中的数据
	//fmt.Println("#####url:", ctx.Request.URL.Path)
	if strings.Index(ctx.Request.URL.Path, "/admin/") == 0 || ctx.Request.URL.Path == "/admin" {
		//fmt.Println("hello word")
		return
	}
	//请求controller中的数据
	if strings.LastIndex(ctx.Request.URL.Path, ".do") >= 0 {
		//fmt.Println("hello word")
		return
	}
	//正常的数据

	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/"+ctx.Request.URL.Path)
}

func main() {
	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
	//刚加入的数据
	//设置session
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
