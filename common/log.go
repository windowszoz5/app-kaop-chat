package common

import (
	"drone/compose"
	"drone/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

type KibnanaLog struct {
	Url           string `json:"url"`  //服务名称
	Body          string `json:"body"` //服务名称
	Ip            string `json:"ip"`
	RequestHeader string `json:"request_header"`
	Method        string `json:"method"`
	CtxId         string `json:"ctx_id"`
	Message       string `json:"message"`
}

func KibLog(ctx *gin.Context, write string) {
	buf, _ := ctx.GetRawData()

	//必须先解析Form
	ctx.Request.ParseForm()
	dataMap := make(map[string]string)
	//说明:须post方法,加: 'Content-Type': 'application/x-www-form-urlencoded'
	for key, _ := range ctx.Request.PostForm {
		dataMap[key] = ctx.PostForm(key)
	}
	//
	for key, _ := range ctx.Request.URL.Query() {
		dataMap[key] = ctx.Query(key)
	}
	req := ctx.Request
	var header string
	for i, v := range req.Header {
		header += fmt.Sprintf("%v:%v", i, v)
	}

	data := KibnanaLog{
		Url:           req.Host + req.URL.String(),
		Body:          (string)(buf),
		CtxId:         ctx.Request.Context().Value(X_REQ_UUID).(string),
		RequestHeader: header,
		Ip:            ctx.ClientIP(),
		Method:        req.Method,
		Message:       write,
	}
	_, err := compose.EsClient.Index().
		Index(config.RunConf.Branch).
		Type(config.RunConf.Name).
		BodyJson(data).
		Do()
	if err != nil {
		// Handle error
		panic(err)
		return
	}
}
