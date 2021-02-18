package jwtmiddleware

import  (
	"github.com/gin-gonic/gin"

	"api/service/jwtmanager"
)

func ApiGuard() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.AbortWithStatusJSON(200, gin.H{
				"code": -1,
				"msg" : "没有获取到token",
			})
		}

		claims, err := jwtmanager.ParseJwtTOken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(200, gin.H{
				"code": -1,
				"msg" : "token验证失败",
			})
		}

		ctx.Set("username", claims["username"])

		//ctx.Next()
	}
}