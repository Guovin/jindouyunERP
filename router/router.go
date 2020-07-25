package router

import (
	"jindouyunERP/app/service/middleware"
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
	UserGroup := s.Group("/api")
	//UserGroup.POST("/login", Order, user.Login, func(g *ghttp.RouterGroup) {
	//	g.Middleware(MiddlewareCORS)
	//})
	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareCORS)
		group.POST("/login", user.Login)
	})
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareCORS, middleware.Auth)
		group.GET("/logout", user.LogOut)
		group.GET("/user", user.GetUser)
	})
	//UserGroup.GET("/logout", user.LogOut)

	UserGroup.PATCH("/password", func(r *ghttp.Request) {
		r.Response.Write("Password")
	})
	UserGroup.GET("/users", func(r *ghttp.Request) {
		r.Response.Write("listusers")
	})
	UserGroup.GET("/user", func(r *ghttp.Request) {
		r.Response.Write("getuser")
	})

	UserGroup.GET("/commodities", func(r *ghttp.Request) {
		r.Response.Write("commodities")
	})

	UserGroup.POST("/order/custormer", func(r *ghttp.Request) {
		r.Response.Write("order/custormer")
	})
	UserGroup.GET("/orders/custormer", func(r *ghttp.Request) {
		r.Response.Write("order/custormer")
	})
	UserGroup.DELETE("/order/custormer/:id", func(r *ghttp.Request) {
		r.Response.Write("order/custormer")
	})
	UserGroup.PATCH("/order/custormer/:id/confirm", func(r *ghttp.Request) {
		r.Response.Write("order/custormer")
	})

	//需要admin限权的API

}
