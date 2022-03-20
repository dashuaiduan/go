package __jwt

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type Userinfo struct {
	User
	Id       uint
}
type User struct {
	UserName string `form:"username" binding:"min=4,max=20,alpha,required"`
	NickName string `form:"nickname" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Age      int    `form:"age" binding:"required,gte=12"`
	Phone    string `form:"phone" binding:"required"`
}

// 服务器关机重启之后 所有已颁发token同样有效
//自定义一个字符串  秘钥
var jwtkey = []byte("duandashuai")

type Claims struct {
	UserInfo           Userinfo //  自定义业务数据存储结构  务必用大写啊  小写导不出 会导致解析tokenk 解析不到数据.深坑
	jwt.StandardClaims                    // 加密配置
}

//颁发token
func GetToken(userInfo Userinfo) (string, error) {
	expireTime := time.Now().Add(7 * 24 * time.Hour) // 一周后过期，过期功能jwt已实现
	//expireTime := time.Now().Add( time.Minute) // 一分钟后过期
	claims := &Claims{
		UserInfo: userInfo, // 业务数据
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
		return "", err
	}
	return tokenString, err
}

//解析token
func CheckToken(tokenString string) (interface{}, bool) {
	//tokenString := ctx.GetHeader("Authorization")

	if tokenString == "" {
		return nil, false
	}

	token, claims, err := ParseToken(tokenString)
	if err != nil || !token.Valid {
		return nil, false
	}
	return claims.UserInfo, true
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	return token, Claims, err
}
