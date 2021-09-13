package gorm_update

import (
	"gorm.io/gorm"
	"gorm/model"
)

// Update
func Update(db *gorm.DB) {
	var p1 model.Product1
	db.First(&p1)
	p1.Price = 123455
	db.Save(&p1)
}

//UpdateRow 更新单列 就是更新某个字段的值
func UpdateRow(db *gorm.DB) {
	var p1 model.Product1
	db.Model(&p1).Where("id=?", 20).Update("price", 5666)
}

//UpdateRows 更新多列
func UpdateRows(db *gorm.DB) {
	var p1 model.Product1
	db.Model(&p1).Where("name like ?", "%测试%").Updates(model.Product1{Name: "hq", Price: 123, Code: "10000"})
}
