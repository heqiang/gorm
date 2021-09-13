package main

import (
	"gorm/gorm_conn"
	"gorm/gorm_delete"
	"gorm/model"
)

func main() {

	db, err := gorm_conn.Conn("root", "142212", model.Product1{}, model.Product2{})
	if err != nil {
		panic(err)
	}
	// 创建
	//gorm_create.SingleGet(db)
	// 批量创建
	//gorm_create.BatchCreate(db)

	//查询
	//gorm_query.SingleQuery(db)
	//gorm_query.QueryByPrimaryId(db)
	//gorm_query.QueryByStringCondition(db)
	//gorm_query.QueryByStructAndMap(db)
	//gorm_query.QueryBySpecialFiled(db)
	//gorm_query.QueryByOrder(db)
	//gorm_query.QueryLimitAndOffset(db)

	//更新
	//gorm_update.Update(db)
	//gorm_update.UpdateRow(db)
	//gorm_update.UpdateRows(db)

	//删除
	//gorm_delete.Delete(db)
	gorm_delete.DeleteBulk(db)
}
