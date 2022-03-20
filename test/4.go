package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

//自定义一个字符串
var jwtkey = []byte("www.fafaseo.com")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func main() {
	r := gin.Default()
	r.GET("/set", setting)
	r.GET("/get", getting)
	//监听端口默认为8080
	r.Run(":1123")
}

//颁发token
func setting(ctx *gin.Context) {
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: 2,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",  // 签名颁发者
			Subject:   "user token", //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Println(token)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		fmt.Println(err)
	}

	ctx.JSON(200, gin.H{"token": tokenString})
}

//解析token
func getting(ctx *gin.Context) {

	//fmt.Println(ctx.Request.Header)

	tokenString := ctx.GetHeader("Authorization")
	//fmt.Println(tokenString)

	//vcalidate token formate
	if tokenString == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		ctx.Abort()
		return
	}

	token, claims, err := ParseToken(tokenString)
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		ctx.Abort()
		return
	}
	fmt.Println(111)
	fmt.Println(claims.UserId)
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "验证通过"})
	ctx.Abort()
	return
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	return token, Claims, err
}
