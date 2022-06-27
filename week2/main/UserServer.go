package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"week2/query"
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
	router := gin.Default()

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		parseInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.String(http.StatusOK, "Hello %s", "parse id error!")
		} else {
			user := query.QuryByid(parseInt)
			json, _ := json.Marshal(user)
			c.String(http.StatusOK, "Hello %s", json)
		}
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/list", func(c *gin.Context) {
		allUser := query.QuryAllUser()
		allstudentjson, _ := json.Marshal(allUser)
		c.String(http.StatusOK, string(allstudentjson))
	})

	// This handler will match /user/john but will not match /user/ or /user
	router.POST("/add", func(c *gin.Context) {
		var userIn query.User
		if err := c.ShouldBindJSON(&userIn); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		query.AddUser(userIn)
		c.JSON(http.StatusOK, gin.H{"status": "add success"})
	})

	router.DELETE("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		parseInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.String(http.StatusOK, "del %s", "parse id error!")
		} else {
			query.Del(parseInt)
			c.String(http.StatusOK, "del %s", parseInt)
		}
	})

	router.Run(":8080")
}
