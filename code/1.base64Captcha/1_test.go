// example of HTTP server that uses the captcha package.
package ddd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"net/http"
	"testing"
)

// base64Captcha create http handler
func generateCaptchaHandler(c *gin.Context) {
	//github.com/mojocn/base64Captcha
	//字符,公式,验证码配置
	var configC = base64Captcha.ConfigCharacter{
		Height: 60,
		Width:  240,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeNumber,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         4,
	}
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	idKeyC, capC := base64Captcha.GenerateCaptcha("", configC)
	//以base64编码
	base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(capC)

	//fmt.Println(idKeyC)
	c.JSON(200, gin.H{"code": http.StatusOK, "id": idKeyC, "data": base64stringC})
	//c.JSON(http.StatusOK,base64stringC)
}
// base64Captcha verify http handler
func captchaVerifyHandle(c *gin.Context) {
	IdKey := c.Query("IdKey")
	ValidateCode := c.Query("ValidateCode")
	//比对验证码 通过给定的id密钥验证验证码并 删除存储中的验证码值，返回布尔值
	//return base64Captcha.VerifyCaptcha(IdKey ,ValidateCode)
	// 不删除缓存 验证码 可以多次验证
	fmt.Println( base64Captcha.VerifyCaptchaAndIsClear(IdKey, ValidateCode, false))
}
func TestName(t *testing.T) {

	//start a net/http server
	//启动golang net/http 服务器

	r := gin.Default()
	r.GET("/get",generateCaptchaHandler)
	r.GET("/ver",captchaVerifyHandle)
	r.Run(":33333")


}
