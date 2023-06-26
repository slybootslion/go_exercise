package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"miniblog/internal/pkg/known"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查请求头中是否有 `X-Request-ID`，如果有则复用，没有则新建
		requestID := c.Request.Header.Get(known.XRequestIDKey)
		if requestID == "" {
			requestID = uuid.New().String()
		}
		c.Set(known.XRequestIDKey, requestID)
		c.Writer.Header().Set(known.XRequestIDKey, requestID)
		c.Next()
	}
}
