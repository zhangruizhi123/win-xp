package controllers

/*
*
*图标管理,需要将一些图标添加到对应的icon文件夹中
*
 */
import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"win-xp/service"

	"github.com/astaxie/beego"
)

type IconController struct {
	beego.Controller
}

/**
*查询目录下的图标
**/
func (this *IconController) ListIcon() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	realPath := dir + "/static/icon"
	filelist, err := ioutil.ReadDir(realPath)
	if err == nil {
		list := make([]string, 0)
		for _, f := range filelist {
			name := f.Name()
			list = append(list, "/icon/"+name)
		}
		this.Data["json"] = map[string]interface{}{"success": 0, "message": "获取图标成功", "data": list}
	} else {
		this.Data["json"] = map[string]interface{}{"success": 100, "message": "获取图标失败"}
	}

	this.ServeJSON()
}

/**
*查询数据库图标
**/
func (this *IconController) SelectIcon() {
	fileTypeService := service.FileTypeService{}
	list, err := fileTypeService.SelectList()
	if err == nil {
		this.Data["json"] = map[string]interface{}{"success": 0, "message": "查询成功", "data": list}
	} else {
		this.Data["json"] = map[string]interface{}{"success": 100, "message": "查询失败"}
	}
	this.ServeJSON()
}

/**
*更新或者插入
**/
func (this *IconController) IUItem() {
	id, err1 := this.GetInt("id")
	name := this.GetString("name")
	icon := this.GetString("icon")
	fileTypeService := service.FileTypeService{}
	err2 := fileTypeService.IUItem(id, name, icon)
	if err2 == nil {
		if id > 0 {
			this.Data["json"] = map[string]interface{}{"success": 0, "message": "更新成功"}
		} else {
			this.Data["json"] = map[string]interface{}{"success": 0, "message": "插入成功"}
		}
	} else {
		if id > 0 {
			this.Data["json"] = map[string]interface{}{"success": 100, "message": "更新失败"}
		} else {
			this.Data["json"] = map[string]interface{}{"success": 100, "message": "插入失败"}
		}

	}
	fmt.Println(id, err1)
	this.ServeJSON()
}
