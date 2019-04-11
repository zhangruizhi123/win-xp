package controllers

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/astaxie/beego"
)

type FileController struct {
	beego.Controller
}

func (this *FileController) List() {
	path := this.GetString("name")
	fmt.Println("hello word", path)
	//dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	dir := beego.AppConfig.String("path")
	realDir := dir + path
	filelist, err := ioutil.ReadDir(realDir)
	list := make([]map[string]interface{}, 0)
	if err == nil {
		for _, f := range filelist {
			item := make(map[string]interface{})
			name := f.Name()
			item["name"] = name
			item["isDir"] = f.IsDir()
			item["size"] = f.Size()
			item["time"] = f.ModTime().Format("2006-01-02 15:04:05")
			list = append(list, item)
		}
		this.Data["json"] = map[string]interface{}{"success": 0, "message": "查询成功", "data": list}
	} else {
		this.Data["json"] = map[string]interface{}{"success": 100, "message": "文件或者目录不存在"}
	}

	this.ServeJSON()
}

func (this *FileController) Read() {

}

func (this *FileController) Delete() {
	//获取要删除的文件
	path := this.GetString("name")
	//dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	dir := beego.AppConfig.String("path")
	realDir := dir + path
	fmt.Println("删除文件:", realDir)
	err := os.Remove(realDir)
	if err == nil {
		this.Data["json"] = map[string]interface{}{"success": 0, "message": "删除成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"success": 100, "message": err.Error()}
	}
	this.ServeJSON()
}
