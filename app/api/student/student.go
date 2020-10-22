package student

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"goframe/app/api/response"
	"goframe/app/model/student"
)

// 获取数据库连接
var db = g.DB().Table("student")

// 新增
func Save(r *ghttp.Request) {
	var stu *student.Entity
	if err := r.Parse(&stu); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	db.Data(stu).Insert()
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
	var stu *student.Entity
	if err := r.Parse(&stu); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	db.Data(stu).Where("id", stu.Id).Update()
	r.Response.WriteJsonExit("更新成功")
}

// 删除
func Delete(r *ghttp.Request) {
	db.Where("id", r.Get("id")).Delete()
}
