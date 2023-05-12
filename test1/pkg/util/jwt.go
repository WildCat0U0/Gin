//package util
//
//import (
//	"awesomeProject/test1/pkg/setting"
//	"github.com/dgrijalva/jwt-go"
//	"time"
//)
//
//var jwtSecret = []byte(setting.JwtSecret) // 这里的JwtSecret是在conf\app.ini中定义的
//
//type Claims struct {
//	Username string `json:"username"`
//	Password string `json:"password"`
//	jwt.StandardClaims
//}
//
//// GenerateToken 生成token
//func GenerateToken(username, password string) (string, error) {
//	nowTime := time.Now()
//	expireTime := nowTime.Add(3 * time.Hour)
//	claims := Claims{
//		Username: username,
//		Password: password,
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: expireTime.Unix(),
//			Issuer:    "awesomeProject",
//		},
//	}
//	tokenCLaims := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
//	token, err := tokenCLaims.SignedString(jwtSecret)
//
//	return token, err
//}
//
//// ParseToken 解析token
//func ParseToken(token string) (*Claims, error) {
//	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
//		return jwtSecret, nil
//	})
//
//	if tokenClaims != nil {
//		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
//			return claims, nil
//		}
//	}
//
//	return nil, err
//}

package util

import (
	"awesomeProject/test1/pkg/setting"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
