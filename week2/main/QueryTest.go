package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

/*# 根据 id 查询用户
/user/${id} get
# 返回用户列表
/user/list  get
# 新增用户
# 用户名只能由字母、数字和下划线组成，长度小于等于16，手机号符和通用手机号规则
/user/add post
# 删除用户
/user/${id}  delete*/

func main() {
	type User struct {
		ID         uint `gorm:"primarykey"`
		Username   string
		Password   string
		Phone      string
		CreateTime time.Time
		UpdateTime time.Time
		IsDeleted  bool `gorm:"index"`
	}

	type DB struct {
		db *gorm.DB
	}

	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/week2?charset=utf8mb4&parseTime=True"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})

	if err != nil {
		panic("failed to connect database")
	}

	user := User{Username: "Jinzhu", Password: "123", Phone: "152666666555", CreateTime: time.Now(), UpdateTime: time.Now()}

	result := db.Create(&user) // 通过数据的指针来创建
	fmt.Println(result)

}
