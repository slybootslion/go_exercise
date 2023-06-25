package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/slybootslion/miniblog-t/internal/pkg/known"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.Request.Header.Get(known.XRequestIdKey)
		if requestId == "" {
			requestId = uuid.New().String()
		}
		// 将RequestId保存在gin.Context中，方便后边程序使用
		c.Set(known.XRequestIdKey, requestId)
		// 将RequestId保存在Http返回头重，header的键为"X-Request-ID"
		c.Writer.Header().Set(known.XRequestIdKey, requestId)

		c.Next()
	}
}
