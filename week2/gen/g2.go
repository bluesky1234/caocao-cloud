package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

const producttableSQL = "CREATE TABLE `products` (" +
	"  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT," +
	"  `Code` varchar(255) DEFAULT NULL," +
	"  `Price` varchar(255) DEFAULT NULL," +
	"  PRIMARY KEY (`id`)" +
	") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4"

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

	// 迁移 schema
	db.AutoMigrate(&User{})

	// Create
	db.Create(&User{Username: "D42", Password: "123", Phone: "15267169537", IsDeleted: false})

	// Read
	var user User
	db.First(&user, 1)                     // 根据整形主键查找
	db.First(&user, "username = ?", "D42") // 查找 code 字段值为 D42 的记录
	fmt.Printf("查询fist记录:%x", user)
	// Update - 将 product 的 price 更新为 200
	db.Model(&user).Update("Password", "222")
	// Update - 更新多个字段
	db.Model(&user).Updates(User{Username: "zhansan", Password: "zhansan"}) // 仅更新非零值字段
	db.Model(&user).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	fmt.Printf("更新记录:%x", user)
	db.Model(&user).Commit()
	// Delete - 删除 product
	//db.Delete(&user, 1)
}
