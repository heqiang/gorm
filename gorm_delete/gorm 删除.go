package gorm_delete

import (
	"gorm.io/gorm"
	"gorm/model"
)

// 单条删除 需要指定主键 否则则会触发批量删除
// gorm 中的删除都是逻辑删除
func Delete(db *gorm.DB) {
	var p1 model.Product1
	db.Delete(&p1, 22)
}

//批量删除
func DeleteBulk(db *gorm.DB) {
	var p1 model.Product1
	db.Where("name =?", "hq").Delete(&p1)
}
