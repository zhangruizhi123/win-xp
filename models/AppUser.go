package models

type AppUser struct {
	NickName string `json:"nickName" 	xorm:"nick_name pk"` //用户名
	Age      int    `json:"age" 		xorm:"age"`              //年龄
	Sex      bool   `json:"sex" 		xorm:"sex"`              //性别
	Id       int    `json:"id" 			xorm:"id"`               //主键
	Password string `json:"password"	xorm:"password"`      //密码
}
