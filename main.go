package main

import (
	"GinProject/06_go_gorm/gorm_conn"
	"GinProject/06_go_gorm/gorm_query"
	"GinProject/06_go_gorm/model"
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
	gorm_query.QueryLimitAndOffset(db)
}
