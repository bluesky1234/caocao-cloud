package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	type User2 struct {
		ID   uint `gorm:"primarykey"`
		Name string
		Age  uint
	}

	//db, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/week2?charset=utf8mb4&parseTime=True"))

	if err != nil {
		panic("failed to connect database")
	}

	user2 := User2{Name: "Jinzhu", Age: 18}

	result := db.Create(&user2) // 通过数据的指针来创建
	fmt.Println(result)
}
