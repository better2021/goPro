package model

import "github.com/jinzhu/gorm"

type Classification struct {
	gorm.Model
	Name   string
	MenuId int
}

type Product struct {
	gorm.Model
	ProductName   string
	Price   int
	Desc   string
	Classifications []Classification `gorm:"ForeignKey:MenuId;AssociationForeignKey:ProductName"`
}
