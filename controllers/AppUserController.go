package controllers

import (
	"MyWeb2/service"
	"fmt"

	"github.com/astaxie/beego"
)

type AppUserController struct {
	beego.Controller
}

func (this *AppUserController) SelectById() {
	userIds := this.GetSession("userId")
	id, err := this.GetInt("id")
	if err != nil {
		this.Data["json"] = map[string]interface{}{"success": 100, "message": "请输入用户id"}
	} else {
		appUserService := service.AppUserService{}
		appUser := appUserService.SelectById(id)
		this.Data["json"] = appUser
		if userIds == nil {
			this.SetSession("userId", appUser)
			fmt.Println("没有session")
		} else {
			fmt.Println("session:", userIds)
		}

	}

	this.ServeJSON()
}
