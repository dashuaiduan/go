package __jwt

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	// 获取 token
	user := Userinfo{User{"zszszszs","张萨姆","7554317612@qq.com",45,"1771252465465"},55}
	tokenStr,err := GetToken(user)
	fmt.Println(tokenStr,err)
//	解析token
	tmp,ok := CheckToken(tokenStr)
	fmt.Println(tmp,ok)


}
