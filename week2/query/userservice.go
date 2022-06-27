package query

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type User struct {
	ID         uint `gorm:"primarykey"`
	Username   string
	Password   string
	Phone      string
	CreateTime time.Time
	UpdateTime time.Time
	IsDeleted  bool `gorm:"index"`
}

var MyDB, err = gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/week2?charset=utf8mb4&parseTime=True"), &gorm.Config{
	NamingStrategy: schema.NamingStrategy{
		SingularTable: true, // 使用单数表名
	},
})

func QuryByid(uid int64) User {
	var user User
	MyDB.First(&user, uid) // 查找 code 字段值为 D42 的记录
	return user
}

func QuryAllUser() []User {
	var users []User
	result := MyDB.Find(&users)
	fmt.Println(result)
	return users
}

func AddUser(user User) {
	create := MyDB.Create(&user)
	fmt.Println(create)
}

func Del(uid int64) {
	var user User
	MyDB.Delete(&user, uid)
}

/*# 根据 id 查询用户
/user/${id} get
# 返回用户列表
/user/list  get
# 新增用户
# 用户名只能由字母、数字和下划线组成，长度小于等于16，手机号符和通用手机号规则
/user/add post
# 删除用户
/user/${id}  delete*/

//func main() {
//	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/week2?charset=utf8mb4&parseTime=True"), &gorm.Config{
//		NamingStrategy: schema.NamingStrategy{
//			SingularTable: true, // 使用单数表名
//		},
//	})
//
//	if err != nil {
//		panic("failed to connect database")
//	}
//
//	user := User{Username: "Jinzhu", Password: "123", Phone: "152666666555", CreateTime: time.Now(), UpdateTime: time.Now()}
//
//	result := db.Create(&user) // 通过数据的指针来创建
//	fmt.Println(result)
//
//}

func main() {

	//user := QuryByid(1)
	//fmt.Println(user)
	//QuryAllUser()
	usernew := User{Username: "lisi", Password: "123", Phone: "152666666535", CreateTime: time.Now(), UpdateTime: time.Now()}
	AddUser(usernew)
}
