package jwtmiddleware

import  (
	"github.com/gin-gonic/gin"
	"api/service/jwtmanager"
	"api/model"
)

var fullPaths map[string]interface{}

func init() {
	fullPaths = map[string]interface{}{}
	fullPaths["/api/admin/login"] = "api/admin/login"
	fullPaths["/api/admin/menu/check"] = "api/admin/menu/check"
}

func JwtGuard() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if _, ok := fullPaths[ctx.FullPath()]; ok {
			ctx.Next()
			return
		}
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.AbortWithStatusJSON(200, gin.H{
				"code": 40001,
				"msg" : "没有获取到token",
			})
			return
		}
		
		claims, err := jwtmanager.ParseJwtToken(token)
		
		if err != nil {
			ctx.AbortWithStatusJSON(200, gin.H{
				"code": 40001,
				"msg" : "token验证失败, " + err.Error(),
			})
			return
		}

		ctx.Set("username", claims["username"])

		ctx.Next()
	}
}

func ApiGuard() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if _, ok := fullPaths[ctx.FullPath()]; ok {
			ctx.Next()
			return
		}
		username := ctx.MustGet("username")
		var admin model.Admin
		model.ORM.Where("username = ?", username).Find(&admin)
		var api model.AdminActionApi
		model.ORM.Where("method = ? and path = ?", ctx.Request.Method, ctx.FullPath()).Find(&api)
		var roleApi model.AdminRoleActionApi
		model.ORM.Where("role_id = ? and api_id = ?", admin.RoleId, api.ID).First(&roleApi)
		if roleApi.ID == 0 { //or errors.Is(, ErrResultNotFound)
			ctx.AbortWithStatusJSON(200, gin.H{
				"code": 1,
				"msg" : "没有api访问权限",
			})
			return
		}
		ctx.Next()
	}
}