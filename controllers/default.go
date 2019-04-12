package controllers

import (
	"win-xp/common"
	"win-xp/models"
	"fmt"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["test"] = "hello word"
	c.TplName = "index.tpl"
}

func (this *MainController) TestJson() {
	fmt.Println("json::::")
	user := &models.AppUser{}
	common.Engine.Get(user)
	fmt.Println(user)
	this.Data["json"] = map[string]interface{}{"success": 0, "message": "111"}
	this.ServeJSON(true)
	return
}
