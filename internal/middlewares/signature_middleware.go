package middlewares

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	myconfig "github.com/devopsfaith/krakend-ce/v2/internal/config"
)

// InitSignatureMiddleware InitSignatureMiddleware
func InitSignatureMiddleware() *SignatureMiddleware {
	return &SignatureMiddleware{}
}

// SignatureMiddleware SignatureMiddleware
type SignatureMiddleware struct{}

// Apply Apply
func (sm *SignatureMiddleware) Apply(c *gin.Context) {
	log.Println("this is a signature valid middleware")

	fmt.Println(myconfig.Case.Signature.Salts)

	typeContent := c.Request.Header.Get("Content-Type")
	if "" == typeContent {
		typeContent = "application/json"
	}

	if c.Request.Body == nil {
		AbortWithError(c, http.StatusNotAcceptable, GetLangContent("", "", "HTTP请求Body错误"))
		return
	}

	var bytesBody []byte
	bytesBody, _ = ioutil.ReadAll(c.Request.Body)

	props, err := GenPropsByBody(typeContent, bytesBody)
	if err != nil {
		AbortWithError(c, http.StatusNotAcceptable, GetLangContent("", "", "解析HTTP参数错误"))
		return
	}

	appID := c.Request.Header.Get("X-APPID")
	saltSign := "f90f4ec04b10"

	if "" != appID {
		mapSalt := myconfig.Case.Signature.Salts[appID]
		if mapSalt != nil {
			saltSign = mapSalt.(string)
		}
	}

	// 处理各个appid对应的salt
	log.Println("salt:", saltSign[:4])

	_sign := WechatSign(props, saltSign, "signature")
	log.Println("====== api signed : ", _sign, props["signature"])
	if props["signature"] != _sign {
		AbortWithError(c, http.StatusNotAcceptable, GetLangContent("", "", "验签错误"))
		return
	}

	// 把刚刚读出来的再写进去
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bytesBody))

	c.Next()
}

// AbortWithError AbortWithError
func AbortWithError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"status": code,
		"msg":    message,
		"data":   "",
	})
	c.Abort()
}

// GetLangContent GetLangContent
func GetLangContent(code, lang, _default string) string {
	if len(code) == 0 {
		return _default
	}

	if len(lang) == 0 {
		lang = "cn"
	}

	return _default
}
