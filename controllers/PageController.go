package controllers

/**
**后台管理员页面的跳转
**
**/
import (
	"MyWeb2/service"
	"strings"

	"github.com/astaxie/beego"
)

type PageController struct {
	beego.Controller
}

func (this *PageController) ToPage() {

	path := this.Ctx.Request.URL.Path

	if path == "/admin/" || path == "/admin/index.html" {
		msg := this.GetSession("msg")
		this.Data["msg"] = msg
		this.TplName = "login.tpl"
		return
	}
	user := this.GetSession("Login")
	if user == nil {
		this.TplName = "login.tpl"
		this.Redirect("/admin/index.html", 302)
	} else {
		index := strings.LastIndex(path, ".html")
		if index > 0 {
			pt := subString(path, 1, index-1)
			sysMenuService := &service.SysMenuService{}
			menu := sysMenuService.ListMenu()
			this.Data["adminUser"] = user
			this.Data["menu"] = menu
			this.TplName = pt + ".tpl"
		} else {
			this.TplName = "login.tpl"
		}
	}

}

func (this *PageController) Page() {
	msg := this.GetSession("msg")
	this.Data["msg"] = msg
	this.TplName = "login.tpl"
}

/**
**字符串截取
**/
func subString(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	// 返回子串
	return string(rs[begin:end])
}
