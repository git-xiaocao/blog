package model

import "time"

type BaseModel struct {
	Id       int64      `json:"id" gorm:"primaryKey"`
	CreateAt *time.Time `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt *time.Time `json:"update_at" gorm:"autoCreateTime"`
}
