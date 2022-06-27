package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

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

	//db, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/week2?charset=utf8mb4&parseTime=True"))
	//db.Exec(producttableSQL)

	if err != nil {
		panic("failed to connect database")
	}

	var user User
	db.First(&user, "username = ?", "lisi") // 查找 code 字段值为 D42 的记录
	fmt.Printf("查询fist记录:%x", user)
	db.Create(&User{Username: "D42", Password: "123", Phone: "15267169537", IsDeleted: false})
	db.First(&user, "username = ?", "D42") // 查找 code 字段值为 D42 的记录
	fmt.Printf("查询fist记录:%x", user)

	// Delete - 删除 product
	//db.Delete(&user, 1)
}
