package main

//
//import (
//	"gorm.io/gen/examples/dal"
//	"time"
//	"week2/model"
//	"week2/query"
//)
//
//func main() {
//	// u refer to query.user
//	user := model.User{Username: "Modi", Password: "123", IsDeleted: 0, CreateTime: time.Now(), UpdateTime: time.Now()}
//
//	u := query.Use(dal.DB).User
//	err := u.WithContext(ctx).Create(&user) // pass pointer of data to Create
//
//	err // returns error
//}
