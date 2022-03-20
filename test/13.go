package main

import (
	"fmt"
)

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

type user struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
	Sex  string `json:"Sex"`
}

func (user) MarshalJSON() ([]byte, error)       { return nil, nil }
func (*user) UnmarshalJSON([]byte) error        { return nil }
func (user) MarshalEasyJSON(w *jwriter.Writer)  {}
func (*user) UnmarshalEasyJSON(l *jlexer.Lexer) {}

type EasyJSON_exporter_user *user

func main() {
	user := new(user)
	//user.Age=11
	//user.Name="dfdf"
	//user.Sex = "男"
	str := `{"Name":"zs","Age":15,"Sex":"男"}`
	//json.Unmarshal([]byte(str),user)
	user.UnmarshalJSON([]byte(str))
	fmt.Println(user)
	//ss ,err := json.Marshal(user)
	//fmt.Println(string(ss),err)
}
