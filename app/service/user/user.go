package user

import (
	"jindouyunERP/app/service/middleware"
	"jindouyunERP/app/service/token"
	"log"
	"net/http"
	"time"

	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/net/ghttp"

	"jindouyunERP/app/model/users"
)

var User struct {
	Name     string `string:"tel" v:"required"`
	PassWord string `string:"password v:"required"`
}

// 用户登录
func Login(r *ghttp.Request) {
	User.Name = r.GetQueryString("tel")
	User.PassWord = r.GetQueryString("password")
	//one, err := users.FindOne("tel=? and password=?", User.Name, User.PassWord)
	one, err := users.VerifyUser(User.Name, User.PassWord)
	if err != nil {
		r.Response.WriteJsonExit(http.StatusUnauthorized)
	}

	role, err := users.GetUserRole(one.UserID)
	if err != nil {
		r.Response.WriteJsonExit(http.StatusInternalServerError)
	}

	clientIP := r.GetClientIp()
	//time.Now().Add(time.Hour * 24 * 3).Unix() 三天后过期
	tokenStr, err := token.Encode(one.Name, one.UserID, clientIP, role, "jindouyun", time.Now().Add(time.Hour*24*3).Unix())
	if err != nil {
		log.Fatal(err)
		r.Response.WriteJsonExit(http.StatusInternalServerError)
	}
	r.Response.Header().Set("Authorization", tokenStr)
	r.Response.WriteJsonExit(http.StatusOK)
}

//用户注销
func LogOut(r *ghttp.Request) {
	r.Response.Header().Set("Authorization", "")
	r.Response.WriteJsonExit(http.StatusOK)
}

//获取用户信息
func GetUser(r *ghttp.Request) {
	//从auth中获取鉴权成功后的userID
	userID := r.GetString(middleware.RequestUserIDKey)
	_, err := users.FindOne("user_id = ?", userID)
	if err != nil {
		r.Response.WriteJsonExit(http.StatusBadRequest)
	}
	user := users.GetUser(userID)
	roles, err2 := users.GetUserRole(userID)
	if err2 != nil {
		r.Response.WriteJsonExit(http.StatusInternalServerError)
	}
	r.Response.WriteJson(g.Map{
		"name":     user.Name,
		"tel":      user.Tel,
		"position": user.Position,
		"roles":    roles,
	})
}
