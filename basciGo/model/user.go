package model

import "time"

type BasicModel struct {
	ID        int       `json:"id,omitempty"`
	CreatedAt time.Time `json:"createAt" example:"创建时间"`
	UpdatedAt time.Time `json:"updateAt" example:"更新时间"`
}

type User struct {
	BasicModel
	UserName  string `json:"userName" example:"用户名" binding:"required"`
	Birthday  string `json:"birthday" example:"出生年月"`
	Sex string `json:"sex" example:"性别"`
	Desc string `json:"desc" example:" 介绍"`
}