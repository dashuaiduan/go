package email

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	err := Send("755431761@qq.com","测试一下","测试主体")
	fmt.Println(err)
}