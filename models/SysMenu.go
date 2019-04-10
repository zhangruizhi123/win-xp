package models

import (
	"time"
)

type SysMenu struct {
	Id       int       `json:"id" 		xorm:"id pk"`           //主键
	ParentId int       `json:"parentId" 		xorm:"parent_id"` //父id
	Name     string    `json:"name" 		xorm:"name"`          //名字
	Url      string    `json:"url" 		xorm:"url"`            //连接
	Icon     string    `json:"icon" 		xorm:"icon"`          //图标
	Remarks  string    `json:"remarks" 		xorm:"remarks"`    //备注
	CUser    int       `json:"cUser" 		xorm:"c_user"`       //创建用户
	CTime    time.Time `json:"cTime" 		xorm:"c_time"`       //创建时间
	Od       int       `json:"od" 		xorm:"od"`              //排序
}

type SysMenuTree struct {
	Menu    SysMenu   //菜单
	SubMenu []SysMenu //子菜单
}
