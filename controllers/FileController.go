package controllers

import (
	"fmt"
	"io/ioutil"
	"os"

	"path"
	"win-xp/service"

	"github.com/astaxie/beego"
)

type FileController struct {
	beego.Controller
}

func (this *FileController) List() {
	path2 := this.GetString("name")
	fmt.Println("hello word", path2)
	//dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	dir := beego.AppConfig.String("path")
	realDir := dir + path2
	filelist, err := ioutil.ReadDir(realDir)
	//查出所有的文件图标

	fileTypeService := service.FileTypeService{}
	iconList, err := fileTypeService.SelectList()

	iconMap := make(map[string]string)
	if err == nil {
		for _, iconF := range iconList {
			iconMap["."+iconF.Name] = iconF.Icon
		}
	}

	list := make([]map[string]interface{}, 0)
	if err == nil {
		for _, f := range filelist {
			item := make(map[string]interface{})
			name := f.Name()
			item["name"] = name
			item["isDir"] = f.IsDir()
			if f.IsDir() {
				item["img"] = iconMap[".fold"]
			} else {
				end := path.Ext(name)
				if end == "" {
					item["img"] = iconMap[".default"]
				} else {
					item["img"] = iconMap[end]
				}
			}
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
	path := this.GetString("name")
	fmt.Println("pp:", path)
	dir := beego.AppConfig.String("path")
	realDir := dir + path
	byt, err := ioutil.ReadFile(realDir)
	fmt.Println("path:", realDir)
	if err == nil {
		this.Data["json"] = map[string]interface{}{"success": 0, "message": "查询成功", "data": string(byt)}
	} else {
		this.Data["json"] = map[string]interface{}{"success": 100, "message": err.Error()}
	}
	this.ServeJSON()
}

/****
*
*删除文件
*
*
**/
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

/**
*重命名文件
*
 */
func (this *FileController) Rename() {
	//文件现在的名字
	path := this.GetString("name")
	//文件原来的名字
	old := this.GetString("old")
	//文件的根目录
	dir := beego.AppConfig.String("path")

	err := os.Rename(dir+old, dir+path)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"success": 0, "message": "修改成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"success": 100, "message": err.Error()}
	}
	this.ServeJSON()
}

func (this *FileController) Create() {
	path := this.GetString("name")
	type2 := this.GetString("type")
	dir := beego.AppConfig.String("path")
	name := dir + path
	var err error
	var f *os.File
	if type2 == "dir" {
		err = os.Mkdir(name, os.ModePerm)
	} else {
		f, err = os.OpenFile(name, os.O_CREATE, 0766)
		f.Close()
	}
	if err == nil {
		this.Data["json"] = map[string]interface{}{"success": 0, "message": "创建成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"success": 100, "message": err.Error()}
	}
	this.ServeJSON()
}

func (this *FileController) Upload() {
	_, h, e := this.GetFile("file")
	if e != nil {
		this.Data["json"] = map[string]interface{}{"success": 101, "message": "获取文件失败"}
	} else {
		path := this.GetString("path")
		dir := beego.AppConfig.String("path")
		err := this.SaveToFile("file", dir+path+h.Filename)
		if err == nil {
			this.Data["json"] = map[string]interface{}{"success": 0, "message": "上传成功"}
		} else {
			this.Data["json"] = map[string]interface{}{"success": 102, "message": "上传失败"}
		}

	}

	this.ServeJSON()
}
func (this *FileController) Download() {
	path := this.GetString("path")
	name := this.GetString("name")
	dir := beego.AppConfig.String("path")
	this.Ctx.Output.Download(dir+path, name)
}

func (this *FileController) ShowImage() {
	path := this.GetString("path")
	//name := this.GetString("name")
	dir := beego.AppConfig.String("path")
	byt, err := ioutil.ReadFile(dir + path)
	if err == nil {
		this.Ctx.ResponseWriter.Write(byt)
		this.Ctx.Output.ContentType("image/png")
	} else {
		this.Data["json"] = map[string]interface{}{"success": 102, "message": "获取文件失败"}
		this.ServeJSON()
	}

}
