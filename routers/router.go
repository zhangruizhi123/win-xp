package routers

import (
	"MyWeb2/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/v1/test.do", &controllers.MainController{})
	beego.Router("/v1/json.do", &controllers.MainController{}, "*:TestJson")
	beego.Router("/appUser/findOne.do", &controllers.AppUserController{}, "*:SelectById")
	beego.Router("/admin/login.do", &controllers.LoginController{}, "*:Login")
	beego.Router("/admin/*.html", &controllers.PageController{}, "*:ToPage")

	//
	beego.Router("/file/list.do", &controllers.FileController{}, "*:List")
	beego.Router("/file/del.do", &controllers.FileController{}, "*:Delete")
}
