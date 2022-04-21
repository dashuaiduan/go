package viper

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"testing"
)

// 官网文档  ，或者地鼠文档 https://www.topgoer.cn/docs/goday/goday-1crg2dneqeek8
func TestName(t *testing.T) {
	//以yaml文件为实例  其他json ini 等文件同理
	// 读取配置文件   一般 一个v对应一个配置文件
	v := viper.New()
	v.SetConfigName("conf")   //找寻文件的名字
	v.SetConfigType("yaml")   // 找寻文件的类型
	v.AddConfigPath("config") //从当前目录下的哪个文件夹找寻，
	err := v.ReadInConfig()   //读取配置文件
	if err != nil {
		if v, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 文件没有找到错误
			fmt.Println(v)
		} else {
			// 读取时候发生的其他错误
			panic(fmt.Errorf("read config err=%s", err))
		}
	}
	cas := v.GetString("cas")         //Get cas Config
	host := v.GetString("mysql.host") //Get mysql.host Config  支持嵌套读取 用. 分割
	a := v.Get("cas")
	fmt.Println("Get Config")
	fmt.Println(cas, host)
	fmt.Println(a)

	//三、写入配置文件
	//从配置文件中读取配置文件是有用的，但是有时你想要存储在运行时所做的所有修改。为此，可以使用下面一组命令，每个命令都有自己的用途:
	//1、WriteConfig - 将当前的viper配置写入 预定义 的路径并覆盖（如果存在的话）。如果没有预定义的路径，则报错。
	//2、SafeWriteConfig - 将当前的viper配置写入 预定义 的路径。如果没有预定义的路径，则报错。如果存在，将不会覆盖当前的配置文件。
	//3、WriteConfigAs - 将当前的viper配置写入 给定的文件路径。将覆盖给定的文件(如果它存在的话)。
	//4、SafeWriteConfigAs - 将当前的viper配置写入 给定的文件路径。不会覆盖给定的文件(如果它存在的话)。
	v.Set("cas", "cas21")
	err = v.WriteConfigAs("config/tmp.yaml") /// 写入到指定文件
	if err != nil {
		panic(err.Error())
	}

	//四 监控并重新读取配置文件
	v.OnConfigChange(func(in fsnotify.Event) {
		//Can Add Some Logics
		fmt.Println("Config Was Changed", in.Name)
		fmt.Println("------", in.Op)
	})
	v.WatchConfig()

	v.WriteConfig() // 写入到预定义文件 set的配置路径

	//	 支持提取子树
	fmt.Println("____________支持提取子树_______________")
	c := v.Sub("b.c")
	d := v.Sub("b.c")
	fmt.Println(c.Get("c1"), d.Get("d2"))

	//	七、反序列化与序列化
	fmt.Println("____________反序列化与序列化_______________")
	//	tmp := v.AllSettings()
	var tmp Config
	err = v.Unmarshal(&tmp) // 将配置解析到结构体
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(tmp)

	m := v.AllSettings() // 将所有配置 转换成map数据
	fmt.Println(m)

	//	往后其他读取 redis etcd 远程配置 请参照官网...

}

type Config struct {
	Port        int    `mapstructure:"port"`
	Cas         string `mapstructure:"cas"`
	MySqlConfig `mapstructure:"mysql"`
}

type MySqlConfig struct {
	Host   string `mapstructure:"host"`
	Port   int    `mapstructure:"port"`
	Dbname string `mapstructure:"dbname"`
}
