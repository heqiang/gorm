### Gorm新手入门教程

> 官网：[gorm中文官网](https://gorm.io/zh_CN/docs/index.html )


#### <font style="color:blue">一、 环境配置</font>

安装

```go
go get -u gorm.io/gorm
# mysql驱动 本文以mysql驱动为例
go get -u gorm.io/driver/mysql
# sqlite 驱动
go get -u gorm.io/driver/sqlite
```

#### <font style="color:blue">二 、mysql数据库连接</font>

```go
dsn := "用户名:密码@tcp(ip:port)/dbName?charset=utf8mb4&parseTime=True&loc=Local"
// 这个db就是我们想要的*gorm.db实例了,后续就可以在这个基础上进行增删改查
db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
if err!=nil{
    fmt.Println("mysql  conn err:",err)
}
```

#### <font style="color:blue">三、MySql 模型迁移</font> 

```GO
type Model struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt DeletedAt `gorm:"index"`
}

type  Product1 struct {
	gorm.Model // 模型嵌套 自动生成以上字段
	Code  string
	Price uint
	Name  string
}

// AutoMigrate 用于自动迁移您的 schema，保持您的 schema 是最新的
err := db.AutoMigrate(Product{})
if err != nil {
    return nil, err
}

```

> AutoMigrate 会自动创建数据库外键约束，您可以在初始化时禁用此功能

 ```go
  gorm.Open(mysql.Open(dsn), &gorm.Config{
       DisableForeignKeyConstraintWhenMigrating: true,
  })
 ```

GORM 提供了 Migrator 接口，该接口为每个数据库提供了统一的 API 接口

```go
type Migrator interface {
  // 自动迁移
  AutoMigrate(dst ...interface{}) error

  // 数据库
  CurrentDatabase() string
  FullDataTypeOf(*schema.Field) clause.Expr

  // 数据库表的操作
  CreateTable(dst ...interface{}) error
  DropTable(dst ...interface{}) error
  HasTable(dst interface{}) bool
  RenameTable(oldName, newName interface{}) error

  // 表的字段相关操作
  AddColumn(dst interface{}, field string) error
  DropColumn(dst interface{}, field string) error
  AlterColumn(dst interface{}, field string) error
  HasColumn(dst interface{}, field string) bool
  RenameColumn(dst interface{}, oldName, field string) error
  MigrateColumn(dst interface{}, field *schema.Field, columnType *sql.ColumnType) error
  ColumnTypes(dst interface{}) ([]*sql.ColumnType, error)

  // 约束相关操作
  CreateConstraint(dst interface{}, name string) error
  DropConstraint(dst interface{}, name string) error
  HasConstraint(dst interface{}, name string) bool

  //索引相关操作
  CreateIndex(dst interface{}, name string) error
  DropIndex(dst interface{}, name string) error
  HasIndex(dst interface{}, name string) bool
  RenameIndex(dst interface{}, oldName, newName string) error
}
```

#### <font style="color:blue">四、gorm增删改查</font>

##### 1 增的相关操作

```go
// gorm.model中的字段会在创建记录的时候自动插入对应数据，我们只需关心我们定义的字段数据即可
product := Product1{Code: "1024",Price: 128,Name:"测试"}
// 新增一条
// 通过数据的指针来创建，
//对应sql语句 "insert into product values('xx','xx')"
result:=db.Create(&product)
//out
product.ID             // 返回插入数据的主键
result.Error           // 返回 error
result.RowsAffected    // 返回插入记录的条数
```

**用指定的字段插入数据**

```go
对应的sql语句 "insert into (Code,Name)values('xx','xx')"
db.Select("Code","Name").Create(&product)
对应的sql语句 "insert into (price)values('xx','xx')"
db.Omit("Code","Name").Create(&product)
```

**批量插入**

```go
var products  = []Product1{{Code: "1024",Price: 128,Name:"测试"}, {Code: "1023",Price: 128,Name:"测试"},{Code: "1023",Price: 128,Name:"测试"}}
//1 通过构建一个数据切片 将 这个切片直接扔给Create方法
db.Create(&products)
// 2 使用CreateInBatches 分批创建,指定每一批创建的数量
db.CreateInBatches(products, 100)
```

##### 2 查的相关操作 

 **单条查询**

```
var pro Product1
// 查询第一个
// SELECT * FROM users ORDER BY id LIMIT 1;
db.First(&pro)
// 查询最后一条
// SELECT * FROM users ORDER BY id DESC LIMIT 1;
db.Last(&pro)
// SELECT * FROM users LIMIT 1; 
db.Take(&pro)

result := db.First(&user)
result.RowsAffected // returns count of records found
result.Error        // returns error or nil

// check error ErrRecordNotFound
errors.Is(result.Error, gorm.ErrRecordNotFound)
// 以上单挑查询当表中没有数据的时候会报一个gorm.ErrRecordNotFound的错误
```

> 如何避免`ErrRecordNotFound`错误？
>
> 你可以使用`Find`，比如`db.Limit(1).Find(&user)`，`Find`方法可以接受struct和slice的数据,
>
> `Find`方法主要用于检索所有的数据

**通过主键进行查询**

```go
var p1 model.Product1
db.First(&p1,10)
fmt.Println(p1)
var p2 []model.Product1
db.Find(&p2,[]int{1,2,3})
fmt.Println(p2)
```

**String条件查询**

```go
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
    db.Where("updated_at<?",time.Now()).Find(&p2)
    //Between
    db.Where("code between ? and ?", 2, 6).Find(&p2)
    for _, value := range p2 {
        fmt.Println(value)
    }
}
```

**struct 和map 条件查询 **

> 这种查询可以理解位String查询的升级版
>
> QueryByStructAndMap struct 和map 条件查询
> 区别:struct中查询字段为0,"",false或其他零值的会被过滤不会被用于构建查询条件
> map会将所有查询的字段当做条件进行查询

```go
func QueryByStructAndMap(db *gorm.DB) {
	var p1 model.Product1
	//SELECT * FROM users WHERE name = "测试1"  ;
	db.Where(&model.Product1{Name: "测试1", Code: "0",}).Find(&p1)
	//SELECT * FROM users WHERE name = "测试1 and code =0"
	db.Where(map[string]interface{}{"name":"测试1","code":"0"})
	// 主键切片条件
	//SELECT * FROM product1 WHERE id IN (1,3,5,6)
	db.Where([]int{1, 3, 5, 6}).Find(&p1)
	var p2 []model.Product1
	db.Where(&model.Product1{Name: "测试1"}, "name").Find(&p2)
	for _, value := range p2 {
		fmt.Println(value)
	}
}
```

**选择特定字段进行查询**

```go
func  QueryBySpecialFiled(db * gorm.DB)  {
	var  p1 []model.Product1
	db.Select("name","price").Find(&p1)
	for _,value:=range p1{
		//fmt.Println(value.ID)
		fmt.Println(value.Name)
	}
	//效果同上
	db.Select([]string{"name","price"}).Find(&p1)
}
```

**排序方式的查询**

```go
func  QueryByOrder(db *gorm.DB)  {
	var  p1 []model.Product1
	//SELECT * FROM product1 ORDER BY code desc, name;
	db.Order("code desc,name").Find(&p1)
	for _,value:=range p1{
		//fmt.Println(value.ID)
		fmt.Println(value)
	}

}
```

**limit && offset**

>  Limit 指定获取记录的最大数量 Offset 指定在开始返回记录之前要跳过的记录数量

```go
func  QueryLimitAndOffset(db *gorm.DB)  {
	var  p1 []model.Product1
	//select * from product1 limit 3;
	db.Limit(3).Find(&p1)
	var  p2 []model.Product2
	// SELECT * FROM product2 LIMIT 6; (product2)
	// SELECT * FROM product1; (product1)
	db.Limit(6).Find(&p2).Limit(-1).Find(&p1)
	for _,value:=range p1{
		//fmt.Println(value.ID)
		fmt.Println("表1：",value)
	}
	for _,value:=range p2{
		//fmt.Println(value.ID)
		fmt.Println("表2：",value)
	}
	// select * from product1 offset 3 limit 6;
	// 同 limit语法一样 offset(-1)也可以消除offset条件
	db.Offset(3).Limit(6).Find(&p1)
	for _,value:=range p1{
		//fmt.Println(value.ID)
		fmt.Println(value)
	}
}
```

##### **3 更新的相关操作**

```go
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
```

##### **4 删除的相关操作**

```go
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

```

