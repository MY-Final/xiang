package middleware

import (
	"strings"

	"xiang/backend/internal/pkg/apperror"
	jwtutil "xiang/backend/internal/pkg/jwt"
	"xiang/backend/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

const (
	CtxAdminIDKey   = "admin_id"
	CtxAdminNameKey = "admin_username"
)

func JWT(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			response.Fail(c, 401, apperror.CodeUnauthorized, "未授权")
			c.Abort()
			return
		}

		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			response.Fail(c, 401, apperror.CodeUnauthorized, "未授权")
			c.Abort()
			return
		}

		claims, err := jwtutil.Parse(secret, parts[1])
		if err != nil {
			response.Fail(c, 401, apperror.CodeUnauthorized, "token 无效")
			c.Abort()
			return
		}

		c.Set(CtxAdminIDKey, claims.AdminID)
		c.Set(CtxAdminNameKey, claims.Username)
		c.Next()
	}
}
