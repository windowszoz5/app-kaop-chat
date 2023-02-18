package middleware

import (
	"bytes"
	"drone/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// MarkLog 标记请求
func MarkLog(ctx *gin.Context) {
	reqUUID := ctx.GetHeader(common.X_TRACK_ID)
	if reqUUID == "" {
		reqUUID = uuid.New().String()
	}

	//全局便利
	ctx.Set(common.X_TRACK_ID, reqUUID)

	//初始响应
	writer := responseWriter{
		ctx.Writer,
		bytes.NewBuffer([]byte{}),
	}
	ctx.Writer = writer

	ctx.Next()

	//响应后执行
	common.KibLog(ctx, writer.b.String())
}

type responseWriter struct {
	gin.ResponseWriter
	b *bytes.Buffer
}

// 重写 Write([]byte) (int, base) 方法
func (w responseWriter) Write(b []byte) (int, error) {
	//向一个bytes.buffer中写一份数据来为获取body使用
	w.b.Write(b)
	//完成gin.Context.Writer.Write()原有功能
	return w.ResponseWriter.Write(b)
}
