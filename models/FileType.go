package models

import (
	"time"
)

type FileType struct {
	Id    int64     `json:"id" 	xorm:"id pk autoincr"`
	Name  string    `json:"name" 		xorm:"name"`
	Icon  string    `json:"icon" 		xorm:"icon"`
	CTime time.Time `json:"cTime" 		xorm:"c_time"`
}
