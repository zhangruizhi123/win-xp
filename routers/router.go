package routers

import (
	"win-xp/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/v1/test.do", &controllers.MainController{})
	beego.Router("/v1/json.do", &controllers.MainController{}, "*:TestJson")
	beego.Router("/appUser/findOne.do", &controllers.AppUserController{}, "*:SelectById")
	beego.Router("/admin/login.do", &controllers.LoginController{}, "*:Login")
	beego.Router("/admin/*.html", &controllers.PageController{}, "*:ToPage")

	//文件操作
	beego.Router("/file/list.do", &controllers.FileController{}, "*:List")
	beego.Router("/file/del.do", &controllers.FileController{}, "*:Delete")
	beego.Router("/file/rename.do", &controllers.FileController{}, "*:Rename")
	beego.Router("/file/create.do", &controllers.FileController{}, "*:Create")
	beego.Router("/file/upload.do", &controllers.FileController{}, "*:Upload")
	beego.Router("/file/read.do", &controllers.FileController{}, "*:Read")
	beego.Router("/file/download.do", &controllers.FileController{}, "*:Download")
	beego.Router("/file/showImg.do", &controllers.FileController{}, "*:ShowImage")
	//图标操作
	beego.Router("/icon/iconlist.do", &controllers.IconController{}, "*:ListIcon")
	beego.Router("/icon/selectlist.do", &controllers.IconController{}, "*:SelectIcon")
	beego.Router("/icon/iuItem.do", &controllers.IconController{}, "*:IUItem")
}
