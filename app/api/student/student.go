package student

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"goframe/app/model/student"
)

// 获取数据库连接
var db = g.DB().Table("student")

// 新增
func Save(r *ghttp.Request) {
	db.Data(student.Entity{
		Id:   r.GetRequestString("id"),
		Name: r.GetRequestString("name"),
		Age:  r.GetRequestInt("age"),
	}).Insert()
	r.Response.WriteJsonExit("新增成功")
}

// 根据id查询
func Select(r *ghttp.Request) {
	stu := (*student.Entity)(nil)
	db.Where("id", r.Get("id")).Struct(&stu)
	r.Response.WriteJsonExit(stu)
}

// 更新
func Update(r *ghttp.Request) {
	db.Data(student.Entity{
		Id:   r.GetRequestString("id"),
		Name: r.GetRequestString("name"),
		Age:  r.GetRequestInt("age"),
	}).Where("id", r.Get("id")).Update()
	r.Response.WriteJsonExit("更新成功")
}

// 更新
func Delete(r *ghttp.Request) {
	db.Where("id", r.Get("id")).Delete()
}
