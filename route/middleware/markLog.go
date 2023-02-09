package middleware

import (
	"bytes"
	"context"
	"drone/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
)

func MarkLog(ctx *gin.Context) {
	reqUUID := ctx.GetHeader(common.X_REQ_UUID)
	if reqUUID == "" {
		reqUUID = fmt.Sprintf("%v", rand.Int())
	}

	//初始响应
	writer := responseWriter{
		ctx.Writer,
		bytes.NewBuffer([]byte{}),
	}
	ctx.Writer = writer
	ctx.Next()

	//otime := time.Now()
	ctxc := context.WithValue(ctx.Request.Context(), common.X_REQ_UUID, reqUUID)
	ctx.Request = ctx.Request.WithContext(ctxc)

	// 响应后执行
	//latency := time.Since(otime)
	common.KibLog(ctx, writer.b.String())
}

type responseWriter struct {
	gin.ResponseWriter
	b *bytes.Buffer
}

// 重写 Write([]byte) (int, error) 方法
func (w responseWriter) Write(b []byte) (int, error) {
	//向一个bytes.buffer中写一份数据来为获取body使用
	w.b.Write(b)
	//完成gin.Context.Writer.Write()原有功能
	return w.ResponseWriter.Write(b)
}

func PrintResponse(c *gin.Context) {

}
