package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("todoList")

type Claims struct {
	Id uint `json:"id"`
	jwt.StandardClaims
}

// 签发token
func ReleaseToken(id uint) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(8 * time.Hour) // 8小时token过期
	claims := Claims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 过期时间
			IssuedAt:  time.Now().Unix(), // 发布时间
			Issuer:    "todoList",        // 发布者
			Subject:   "todo_List",       // 主题
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 加密
	token, err := tokenClaims.SignedString(jwtKey)                   // 签发token
	return token, err
}

// 校验token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
