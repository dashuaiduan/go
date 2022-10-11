package main

// https://www.cnblogs.com/sparkdev/p/10833186.html
import (
	"fmt"
	flag "github.com/spf13/pflag"
	"strings"
)

//格式化name  将变量的des-detail  转换成des.detail  命令行写.即可
func wordSepNormalizeFunc(f *flag.FlagSet, name string) flag.NormalizedName {
	from := []string{"-", "_"}
	to := "."
	for _, sep := range from {
		name = strings.Replace(name, sep, to, -1)
	}
	return flag.NormalizedName(name)
}
func main() {

	// 定义命令行参数对应的变量  value是默认值 就是没有传递的默认值
	var cliName = flag.StringP("name", "n", "nick", "Input Your Name")
	var cliAge = flag.IntP("age", "a", 0, "Input Your Age")
	var cliGender = flag.StringP("gender", "g", "male", "Input Your Gender")
	var cliOK = flag.BoolP("ok", "o", true, "Input Are You OK")
	var cliDes = flag.StringP("des-detail", "d", "", "Input Description")
	var cliOldFlag = flag.StringP("badflag", "b", "just for test", "Input badflag")
	// 设置标准化参数名称的函数
	flag.CommandLine.SetNormalizeFunc(wordSepNormalizeFunc)

	// 为 age 参数设置 NoOptDefVal  命令行传了 --age  但是没有指定值得时候  默认值就是这个，并不是没有传递的默认值
	flag.Lookup("age").NoOptDefVal = "25"

	// 把 badflag 参数标记为即将废弃的，请用户使用 des-detail 参数
	flag.CommandLine.MarkDeprecated("badflag", "please use --des-detail instead")
	// 把 badflag 参数的 shorthand 标记为即将废弃的，请用户使用 des-detail 的 shorthand 参数
	flag.CommandLine.MarkShorthandDeprecated("badflag", "please use -d instead")

	// 在帮助文档中隐藏参数 badflag   --h 说明中一次当前参数说明
	flag.CommandLine.MarkHidden("badflag")

	// 把用户传递的命令行参数解析为对应变量的值
	flag.Parse()

	fmt.Println("name=", *cliName)
	fmt.Println("age=", *cliAge)
	fmt.Println("gender=", *cliGender)
	fmt.Println("ok=", *cliOK)
	fmt.Println("des=", *cliDes)
	fmt.Println("badflag=", *cliOldFlag)
}
