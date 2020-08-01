package staff

import (
	"encoding/json"
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

//员工信息结构
var Data []struct {
	ID      string `json:"id" v:"required"`
	Name    string `json:"name" v:"required"`
	Age     string `json:"age" v:"required"`
	Address string `json:"address" v:"required"`
}

//保存全部员工信息
func SaveStaff(r *ghttp.Request) {
	source := r.GetQueryString("source")       //获取到的是string: “[{"xxx":"xxx"},{"xxx":"xxx"}]”
	s := json.Unmarshal([]byte(source), &Data) //将string的json数组转换为Data结构对象
	if s != nil {
		r.Response.WriteJsonExit(http.StatusInternalServerError)
	}
	staff := make([]staffs.Staff, len(Data)) //创建Staff结构对象数组，并设置容量
	//遍历导入
	for i, v := range Data {
		staff[i].ID = v.ID
		staff[i].Name = v.Name
		staff[i].Age = v.Age
		staff[i].Address = v.Address
	}
	err := staffs.SaveStaff(staff) //保存所有成员
	if err != nil {
		r.Response.WriteJsonExit(http.StatusInternalServerError)
	}
	r.Response.WriteJsonExit(http.StatusCreated)
}
