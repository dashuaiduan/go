package aa

import (
	"bytes"
	"fmt"
	"html/template"
	"regexp"
	"testing"
)

type AA int

const (
	AA1 = 2
	AA2 = 3
)

type aa struct {
	a AA
}

func reg(s string) string {
	b := []byte(s)
	fmt.Println(string(b[2 : len(b)-2]))
	return "AA"
}
func Test111(t *testing.T) {
	data := "dfgd{{lk}}jn{{gdng}}a{{a}}"
	re3, _ := regexp.Compile("{{.*?}}")
	rep := re3.ReplaceAllStringFunc(data, reg)
	fmt.Println(rep)
}

type Person struct {
	Name string
	Age  int
}
type Data struct {
	User Person
}

//替换模板变量
func Test11121(t *testing.T) {
	var tpl bytes.Buffer
	var data Data
	m := make(map[string]string)
	m["user_Name"] = "5456465"
	m["user_Age"] = "33"
	p := Person{"longshuai11", 2113}
	//p2 := Person{"longshuai1122222",222222}
	data.User = p
	tmpl, _ := template.New("test").Parse("Name: {{.user_Name}}, Age: {{.user_Age}}")
	_ = tmpl.Execute(&tpl, m)
	fmt.Println(tpl.String())
}

func Test65664(t *testing.T) {

}
