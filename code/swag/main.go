package main

//使用函数头注释的方式 自动生成api文档,直接http服务请求api文档内容
//官网地址： https://github.com/swaggo/swag
// 地鼠文档 ： https://www.topgoer.cn/docs/swaggo/swaggo-1disep3530rmo
//支持的web框架
//gin
//echo
//buffalo
//net/http
import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "test/code/swag/docs"
)

// gin的helloWorld

// @BasePath /api/v1

// PingExample godoc
// @Summary ping 示例
// @Schemes
// @Description 执行 ping
// @Tags 示例
// @Accept json
// @Produce json
// @Success 200 {string } main
// @Router /example/main [get]
func main() {
	//docs.SwaggerInfo.BasePath = "/v1"
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// 路由组1 ，处理GET请求
	v1 := r.Group("/v1")
	// {} 是书写规范
	{
		v1.GET("/login", login)
		v1.GET("submit", submit)
	}
	v2 := r.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8800")
}

// @BasePath /api/v1

// PingExample godoc
// @Summary 登录接口
// @Schemes 计划
// @Description 用户登录接口 用于登录
// @Tags example
// @Accept mpfd
// @Produce json
// @Success 200 {string} Helloworld
// @dfgdg 5646546d
// @Router /v1/login [get]
func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  Context
// @Failure      400  {object}  Context
// @Failure      404  {object}  Context
// @Failure      500  {object}  Context
// @Router       /accounts/{id} [get]
func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}
