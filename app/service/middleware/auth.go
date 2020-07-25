package middleware

import (
	"jindouyunERP/app/service/token"
	"net/http"

	"github.com/gogf/gf/net/ghttp"
)

//RequestUserNameKey context key值
const RequestUserNameKey string = "jindouyun-UserName"

//RequestUserIDKey context key值
const RequestUserIDKey string = "jindouyun-UserID"

//RequestRoleKey context key值
const RequestRoleKey string = "jindouyun-Role"

//用户信息jwt鉴权
func Auth(r *ghttp.Request) {
	jwtToken := r.Header.Get("Authorization")
	info, err := token.Decode(jwtToken)
	if err != nil {
		r.Response.WriteJsonExit(http.StatusUnauthorized)
	}
	r.SetParam(RequestUserNameKey, info.UserName)
	r.SetParam(RequestUserIDKey, info.UserID)
	r.SetParam(RequestRoleKey, info.Roles)
	r.Middleware.Next()
}
