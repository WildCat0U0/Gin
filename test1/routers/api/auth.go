// package api
//
// import (
//
//	"awesomeProject/test1/models"
//	"awesomeProject/test1/pkg/e"
//	"awesomeProject/test1/pkg/util"
//	"github.com/astaxie/beego/validation"
//	"github.com/gin-gonic/gin"
//	"log"
//	"net/http"
//
// )
//
//	type auth struct {
//		Username string `valid:"Required; MaxSize(50)`
//		Password string `valid:"Required; MaxSize(50)`
//	}
//
// //中文
//
// // GetAuth 获取token
//
//	func GetAuth(c *gin.Context) {
//		username := c.Query("username")
//		password := c.Query("password")
//
//		valid := validation.Validation{}
//		a := auth{Username: username, Password: password}
//		ok, _ := valid.Valid(&a)
//
//		data := make(map[string]interface{}) // data is a map of string and interface{}
//		code := e.INVALID_PARAMS             // 400
//
//		if ok {
//			isExist := models.CheckAuth(username, password) // CheckAuth is in models/user.go
//			if isExist {
//				token, err := util.GenerateToken(username, password)
//				if err != nil {
//					code = e.ERROR_AUTH_TOKEN
//				} else {
//					data["token"] = token
//					code = e.SUCCESS
//				}
//			} else {
//				code = e.ERROR_AUTH // 401
//			}
//		} else { // if valid.Valid(&a) is false
//			for _, err := range valid.Errors { // valid.Errors is a slice of Error
//				log.Println(err.Key, err.Message) // err.Key is "Username" or "Password"
//			}
//		}
//
//		c.JSON(http.StatusOK, gin.H{ // gin.H is a shortcut for map[string]interface{}
//			"code": code,           // e.SUCCESS is 200
//			"msg":  e.GetMsg(code), // e.GetMsg(e.SUCCESS) is "OK"
//			"data": data,
//		})
//
// }
package api

import (
	"awesomeProject/test1/models"
	"awesomeProject/test1/pkg/e"
	"awesomeProject/test1/pkg/logging"
	"awesomeProject/test1/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)
	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message) // err.Key is "Username" or "Password"
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
