package grpc1

import (
	"a/code/grpc/service"
	"fmt"
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestName(t *testing.T) {
	user := &service.User{
		Username: "mszlu",
		Age:      20,
	}
	//转换为protobuf
	marshal, err := proto.Marshal(user)
	if err != nil {
		panic(err)
	}
	newUser := &service.User{}
	err = proto.Unmarshal(marshal, newUser)
	if err != nil {
		panic(err)
	}
	fmt.Println(newUser.String())
}
