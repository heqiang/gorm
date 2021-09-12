package gorm_conn

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Conn(root, password string, tableStruct ...interface{}) (db *gorm.DB, err error) {

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/spider?charset=utf8mb4&parseTime=True&loc=Local", root, password)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 禁用外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Println("mysql  conn err:", err)
	}
	for _, table := range tableStruct {
		// AutoMigrate 会自动生产外键约束 在配置中已禁用
		err := db.AutoMigrate(table)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
