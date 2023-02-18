package common

import (
	"drone/compose"
	"drone/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

type KibnanaLog struct {
	Url       string `json:"url"`       //请求地址
	Body      string `json:"body"`      //post请求体
	Ip        string `json:"ip"`        //请求IP
	ReqHeader string `json:"reqHeader"` //请求头
	TractId   string `json:"tractId"`   //请求UID
	Method    string `json:"method"`    //请求方法
	Message   string `json:"message"`   //请求相应
}

func KibLog(ctx *gin.Context, write string) {
	//post请求头数据
	buf, _ := ctx.GetRawData()

	//获取参数
	ctx.Request.ParseForm()
	var data string
	for key, _ := range ctx.Request.PostForm {
		data += fmt.Sprintf("%v:%v ", key, ctx.PostForm(key))
	}
	for key, _ := range ctx.Request.URL.Query() {
		data += fmt.Sprintf("%v:%v ", key, ctx.PostForm(key))
	}

	//请求头
	req := ctx.Request
	var header string
	for i, v := range req.Header {
		header += fmt.Sprintf("%v:%v", i, v)
	}

	//写入es
	esData := KibnanaLog{
		Url:       req.Host + req.URL.String(),
		Body:      (string)(buf),
		TractId:   ctx.GetString(X_TRACK_ID),
		ReqHeader: header,
		Ip:        ctx.ClientIP(),
		Method:    req.Method,
		Message:   write,
	}
	_, err := compose.EsClient.Index().
		Index(config.RunConf.Branch).
		Type("server-product").
		BodyJson(esData).
		Do()
	if err != nil {
		// Handle base
		panic(err)
		return
	}
}
