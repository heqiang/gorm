package model

import "gorm.io/gorm"

type Product1 struct {
	gorm.Model
	Code  string
	Price int
	Name  string
}

type Product2 struct {
	gorm.Model
	Code  string
	Price int
	Name  string
}

type Product3 struct {
	gorm.Model
	Code  string
	Price int
	Name  string
}
