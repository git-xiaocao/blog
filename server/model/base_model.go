package model

import "time"

type BaseModel struct {
	Id       int64      `json:"id" gorm:"primaryKey"`
	CreateAt *time.Time `json:"createAt" gorm:"autoCreateTime"`
	UpdateAt *time.Time `json:"updateAt" gorm:"autoCreateTime"`
}
