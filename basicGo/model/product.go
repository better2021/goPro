package model

import (
	"time"
)

type Classification struct {
	Id   int    `gorm:"primary_key" json:"id"`
	Name string `gorm:"not null" json:"name"`
	TimeModel
}

type Product struct {
	Id          int     `gorm:"primary_key" json:"id"`
	ProductName string  `gorm:"not null" json:"productName"`
	Price       float64 `json:"price"`
	Desc        string  `json:"desc"`
	MenuId      int     `gorm:"not null" json:"menuId"`
	TimeModel
}

type TimeModel struct {
	CreatedAt time.Time `json:"createAt" example:"创建时间"`
	UpdatedAt time.Time `json:"updateAt" example:"更新时间"`
}
