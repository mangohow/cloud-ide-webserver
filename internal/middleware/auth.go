package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mangohow/cloud-ide-webserver/pkg/logger"
	"github.com/mangohow/cloud-ide-webserver/pkg/utils/encrypt"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			logger.Logger().Warningf("未获得授权, ip:%s", ctx.Request.RemoteAddr)
			ctx.JSON(http.StatusUnauthorized, nil)
			ctx.Abort()
			return
		}

		username, uid, id, err := encrypt.VerifyToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, nil)
			ctx.Abort()
			return
		}
		ctx.Set("id", id)
		ctx.Set("username", username)
		ctx.Set("uid", uid)

		ctx.Next()
	}
}
