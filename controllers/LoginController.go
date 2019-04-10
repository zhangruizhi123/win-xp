package controllers

import (
	"MyWeb2/service"
	"fmt"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

/**
**用户登录接口
**/
func (this *LoginController) Login() {
	userName := this.GetString("userName")
	if userName == "" {
		this.SetSession("msg", "请输入用户名")
		this.Redirect("/admin/index.html", 302)
		return
	}
	password := this.GetString("password")
	if password == "" {
		this.SetSession("msg", "请输入密码")
		this.Redirect("/admin/index.html", 302)
		return
	}
	appUserService := &service.AppUserService{}
	user := appUserService.UserLogin(userName, password)
	if user != nil {
		this.Data["user"] = user
		this.SetSession("Login", user)
		this.Redirect("/admin/admin.html", 302)
		//this.TplName = "admin/index.tpl"
	} else {
		fmt.Println("密码错误")
		this.SetSession("msg", "用户名或者密码错误")
		this.Redirect("/admin/index.html", 302)
	}
}
