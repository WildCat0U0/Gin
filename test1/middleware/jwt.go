//package middleware
//
//import (
//	"awesomeProject/test1/pkg/e"
//	"awesomeProject/test1/pkg/util"
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"time"
//)
//
//func JWT() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		var code int              // 这里的code是为了保持和其他接口一致
//		var data interface{}      // 这里的data是为了保持和其他接口一致
//		code = e.SUCCESS          // 这里的SUCCESS是200
//		token := c.Query("token") // 这里的token是在前端传过来的
//		if token == "" {          // 这里的token是在前端传过来的
//			code = e.INVALID_PARAMS // 这里的INVALID_PARAMS是400
//		} else {
//			claims, err := util.ParseToken(token) // 这里的ParseToken是在pkg\util\jwt.go中定义的
//			if err != nil {
//				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL // 这里的ERROR_AUTH_CHECK_TOKEN_FAIL是20001
//			} else if time.Now().Unix() > claims.ExpiresAt { // 这里的time.Now().Unix()是为了获取当前时间戳
//				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT // 这里的ERROR_AUTH_CHECK_TOKEN_TIMEOUT是20002
//			}
//		}
//		if code != e.SUCCESS {
//			c.JSON(http.StatusUnauthorized, gin.H{ // 这里的StatusUnauthorized是401
//				"code": code,           // 这里的code是为了保持和其他接口一致
//				"msg":  e.GetMsg(code), // 这里的GetMsg是为了获取对应的错误信息
//				"data": data,           // 这里的data是为了保持和其他接口一致
//			}) // 这里的StatusUnauthorized是401
//			c.Abort() // 这里的Abort是为了防止后续的中间件还会执行
//			return    // 这里的return是为了防止后续的中间件还会执行
//		}
//	}
//}

package jwt

import (
	"awesomeProject/test1/pkg/e"
	"awesomeProject/test1/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = e.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
