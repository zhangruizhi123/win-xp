package service

import (
	"time"
	"win-xp/common"
	"win-xp/models"
)

type FileTypeService struct {
}

/**
*查询列表
**/
func (this *FileTypeService) SelectList() ([]models.FileType, error) {
	fileList := make([]models.FileType, 0)
	err := common.Engine.Find(&fileList)
	return fileList, err
}

func (this *FileTypeService) IUItem(id int, name string, icon string) error {
	types := new(models.FileType)
	types.Icon = icon
	types.Name = name
	//row int64
	//err error
	if id > 0 {
		_, err := common.Engine.Where("id=?", id).Update(types)
		return err
	} else {
		types.CTime = time.Now()
		_, err := common.Engine.InsertOne(types)
		return err
	}
}
