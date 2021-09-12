package gorm_query

import (
	"gorm/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

// SingleQuery 查询单个对象
func SingleQuery(db *gorm.DB) {

	// First 和 Last 会根据主键进行排序,take 不排序
	// 使用前两个方法的时候 会看指定的model是否含有主键 没有主键就按照给定的struct的第一个字段进行排序
	var pro model.Product1
	//res:= db.First(&pro)
	res := db.Take(&pro)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		fmt.Println("没有数据")
	}
	db.Take(&pro)
	db.Last(&pro)
	fmt.Println(pro.ID)
	var pros []model.Product3
	db.Limit(2).Find(&pros)
	for _, value := range pros {
		fmt.Println(value)
	}

	var pp []model.Product1
	db.Limit(10).Find(&pp)
	fmt.Println(pp)
}

// QueryByPrimaryId 根据主键查询
func QueryByPrimaryId(db *gorm.DB) {
	var p1 model.Product1
	db.First(&p1, 10)
	fmt.Println(p1)
	var p2 []model.Product1
	db.Find(&p2, []int{1, 2, 3})
	fmt.Println(p2)
}

// QueryByStringCondition  条件查询
func QueryByStringCondition(db *gorm.DB) {
	// 获取第一条匹配的记录
	var p1 model.Product1
	// select * from product1 where name = "测试" order by id Limit 1;
	db.Where("name =?", "测试0").First(&p1)
	// 获取所有的匹配记录
	var p2 []model.Product1
	db.Where("name =?", "测试").Find(&p2)
	// IN
	db.Where("name in", []string{"ceshi", "admin"}).Find(&p2)
	//Like
	db.Where("name Like ?", "%测%").Find(&p2)
	// And
	db.Where("name =? And code>=?", "测试", "1020")
	//Time 获取当前时间之前的记录
	db.Where("updated_at<?", time.Now()).Find(&p2)
	//Between
	db.Where("code between ? and ?", 2, 6).Find(&p2)
	for _, value := range p2 {
		fmt.Println(value)
	}

}

// QueryByStructAndMap struct 和map 条件查询
// 区别:struct中查询字段为0,"",false或其他零值的会被过滤不会被用于构建查询条件
// map会将所有查询的字段当做条件进行查询
func QueryByStructAndMap(db *gorm.DB) {
	var p1 model.Product1
	//SELECT * FROM users WHERE name = "测试1"  ;
	db.Where(&model.Product1{Name: "测试1", Code: "0"}).Find(&p1)
	//SELECT * FROM users WHERE name = "测试1 and code =0"
	db.Where(map[string]interface{}{"name": "测试1", "code": "0"})
	// 主键切片条件
	//SELECT * FROM product1 WHERE id IN (1,3,5,6)
	db.Where([]int{1, 3, 5, 6}).Find(&p1)
	var p2 []model.Product1
	db.Where(&model.Product1{Name: "测试1"}, "name").Find(&p2)
	for _, value := range p2 {
		fmt.Println(value)
	}
}

// QueryBySpecialFiled 选择特定字段
func QueryBySpecialFiled(db *gorm.DB) {
	var p1 []model.Product1
	db.Select("name", "price").Find(&p1)
	for _, value := range p1 {
		//fmt.Println(value.ID)
		fmt.Println(value.Name)
	}
	//效果同上
	db.Select([]string{"name", "price"}).Find(&p1)
}

// 指定数据库检索结果的排序方式
func QueryByOrder(db *gorm.DB) {
	var p1 []model.Product1
	//SELECT * FROM product1 ORDER BY code desc, name;
	db.Order("code desc,name").Find(&p1)
	for _, value := range p1 {
		//fmt.Println(value.ID)
		fmt.Println(value)
	}

}

// limit && offset
//Limit 指定获取记录的最大数量 Offset 指定在开始返回记录之前要跳过的记录数量
func QueryLimitAndOffset(db *gorm.DB) {
	var p1 []model.Product1
	//select * from product1 limit 3;
	db.Limit(3).Find(&p1)
	var p2 []model.Product2
	// SELECT * FROM product2 LIMIT 6; (product2)
	// SELECT * FROM product1; (product1)
	db.Limit(6).Find(&p2).Limit(-1).Find(&p1)
	for _, value := range p1 {
		//fmt.Println(value.ID)
		fmt.Println("表1：", value)
	}
	for _, value := range p2 {
		//fmt.Println(value.ID)
		fmt.Println("表2：", value)
	}
	// select * from product1 offset 3 limit 6;
	// 同 limit语法一样 offset(-1)也可以消除offset条件
	db.Offset(3).Limit(6).Find(&p1)
	for _, value := range p1 {
		//fmt.Println(value.ID)
		fmt.Println(value)
	}

}
