package staff

import (
	"jindouyunERP/app/model/staffs"
	"net/http"

	"github.com/gogf/gf/net/ghttp"
)

//查询员工信息
func SelectStaff(r *ghttp.Request) {
	staffs, err := staffs.SelectStaff()
	if err != nil {
		r.Response.WriteJsonExit(http.StatusNotFound)
	}
	r.Response.WriteJson(staffs)
}

//添加或修改员工信息
func UpdateStaff(r *ghttp.Request) {
	ID := r.GetQueryString("id")
	_, err := staffs.FindOne(ID)
	if err != nil {
		r.Response.WriteJsonExit(http.StatusNotFound)
	}

	Name := r.GetQueryString("name")
	Age := r.GetQueryString("age")
	Address := r.GetQueryString("address")

	err2 := staffs.UpdateStaff(ID, Name, Age, Address)
	if err2 != nil {
		r.Response.WriteJsonExit(http.StatusInternalServerError)
	}

	r.Response.WriteJsonExit(http.StatusCreated)
}

//删除员工信息
func DeleteStaff(r *ghttp.Request) {
	ID := r.GetQueryString("id")
	_, err := staffs.FindOne(ID)
	if err != nil {
		r.Response.WriteJsonExit(http.StatusNotFound)
	}

	err2 := staffs.DeleteStaff(ID)
	if err2 != nil {
		r.Response.WriteJsonExit(http.StatusInternalServerError)
	}

	r.Response.WriteJsonExit(http.StatusOK)
}
