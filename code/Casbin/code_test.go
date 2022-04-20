package casbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
)

func TestFile(t *testing.T) {
	e, err := casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		log.Fatal(err.Error())
	}

	sub := "alice" // 想要访问资源的用户。
	obj := "data1" // 将被访问的资源。
	act := "read"  // 用户对资源执行的操作。
	// 权限校验
	ok, err := e.Enforce(sub, obj, act)

	if err != nil {
		// 处理err
		log.Fatal(err.Error())
	}

	if ok == true {
		// 允许alice读取data1
		fmt.Println("true")
	} else {
		// 拒绝请求，抛出异常
		fmt.Println("false")
	}

	// 您可以使用BatchEnforce()来批量执行一些请求
	// 这个方法返回布尔切片，此切片的索引对应于二维数组的行索引。
	// 例如 results[0] 是 {"alice", "data1", "read"} 的结果
	results, err := e.BatchEnforce([][]interface{}{{"alice", "data1", "read"}, {"bob", "data2", "write"}, {"jack", "data3", "read"}})
	fmt.Println(results)
	//常用函数 使用 e.
	fmt.Println(e.GetModel())

}

//官网的 xorm 适配器 不好使  报错 深坑  使用gorm
func TestMysql(t *testing.T) {
	a, _ := gormadapter.NewAdapter("mysql", "root:djf136169@tcp(120.77.95.121:3306)/test", "test", "casbin_rule", true)
	e, _ := casbin.NewEnforcer("model.conf", a)
	e.LoadPolicy() // 加载所有策略
	// 禁用AutoSave机制
	e.EnableAutoSave(false)
	fmt.Println(e)
	//其他策略管理api 查看官网  https://casbin.org/docs/zh-CN/api-overview
}
