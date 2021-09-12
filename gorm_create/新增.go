package gorm_create

import (
	"gorm/model"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

// SingleGet 新增单条数据
func SingleGet(db *gorm.DB) {
	product := model.Product1{Code: "1024", Price: 128, Name: "测试"}
	res := db.Create(&product)
	fmt.Println(res.RowsAffected)
	fmt.Println(product.ID)
	db.Select("Code", "Name").Create(&product)
	db.Omit("Code", "Name").Create(&product)
}

// BatchCreate 批量插入
func BatchCreate(db *gorm.DB) {
	var products []model.Product2
	for x := 0; x < 10; x++ {
		pro := model.Product2{
			Name:  "测试" + strconv.Itoa(x),
			Code:  strconv.Itoa(x),
			Price: 500 * (x + 1),
		}
		products = append(products, pro)
	}
	db.Create(&products)
}
