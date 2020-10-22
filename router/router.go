package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"goframe/app/api/hello"
	"goframe/app/api/student"
)

//func Total(r *ghttp.Request) {
//	r.Response.Write("包方法注册")
//}
//
//type Controller struct {
//	total *gtype.Int
//}
//
//func (c *Controller) Total(r *ghttp.Request) {
//	r.Response.Write("对象方法注册:", c.total.Add(1))
//}

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.PUT("/save", student.Save)
		group.GET("/select", student.Select)
		group.POST("/update", student.Update)
		group.DELETE("/delete", student.Delete)
		group.ALL("/template", hello.TestTpl)
	})

	//s.BindHandler("/template", func(r *ghttp.Request) {
	//	tplContent := `id:{{.id}}, name:{{.name}}`
	//	r.Response.WriteTplContent(tplContent, g.Map{
	//		"id"   : 123,
	//		"name" : "john",
	//	})
	//	//r.Response.WriteTpl("index.tpl", g.Map{
	//	//	"id":   123,
	//	//	"name": "john",
	//	//})
	//})
	//s.BindHandler("/template", func(r *ghttp.Request) {
	//	r.Response.WriteTpl("index.tpl", g.Map{
	//		"id":   123,
	//		"name": "john",
	//	})
	//})

	//s.BindHandler("/template", func(r *ghttp.Request){
	//	tplContent := `id:{{.id}}, name:{{.name}}`
	//	r.Response.WriteTplContent(tplContent, g.Map{
	//		"id"   : 123,
	//		"name" : "john",
	//	})
	//})
	//group := s.Group("/")
	//group.PUT("/save", student.Save)
	//group.GET("/select", student.Select)
	//group.POST("/update", student.Update)
	//group.DELETE("/delete", student.Delete)
}

//s.Group("/", func(group *ghttp.RouterGroup) {
//
//})

//group.GET("/get",)

//s.BindHandler("/hello", func(r *ghttp.Request) {
//	r.Response.Write("哈喽世界！")
//})
//
//group := s.Group("/")
//group.POST("/save", hello.Save)

//s.BindHandler("/total1", hello.Test)
//
//c := &Controller{
//	total: gtype.NewInt(),
//}
//s.BindHandler("/total", c.Total)
//
//group := s.Group("/api")
//group.ALL("/all", func(r *ghttp.Request) {
//	r.Response.Write("all")
//})
//
//group.GET("/get", func(r *ghttp.Request) {
//	r.Response.Write("get")
//})
//
//group.POST("/post", func(r *ghttp.Request) {
//	r.Response.Write("post")
//})

//}
