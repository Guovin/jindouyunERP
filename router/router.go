package router

import (
	"jindouyunERP/app/service/middleware"
	"jindouyunERP/app/service/staff"
	"jindouyunERP/app/service/user"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

//授权跨域请求
func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	corsOptions := r.Response.DefaultCORSOptions()
	corsOptions.ExposeHeaders = "Authorization"
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}

func init() {
	s := g.Server()

	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareCORS)
		group.POST("/login", user.Login)
	})
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareCORS, middleware.Auth)
		group.GET("/logout", user.LogOut)
		group.GET("/user", user.GetUser)
		group.POST("/updatePassword", user.UpdatePassword)
		group.POST("/updateStaff", staff.UpdateStaff)
		group.POST("/deleteStaff", staff.DeleteStaff)
		group.GET("/selectStaff", staff.SelectStaff)
		group.POST("/saveStaff", staff.SaveStaff)
	})
}
