package router

import (
	"github.com/gogf/gf/container/gtype"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"goframe/app/api/hello"
)

func Total(r *ghttp.Request) {
	r.Response.Write("包方法注册")
}

type Controller struct {
	total *gtype.Int
}

func (c *Controller) Total(r *ghttp.Request) {
	r.Response.Write("对象方法注册:", c.total.Add(1))
}

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/", hello.Hello)
		group.ALL("/post", hello.Test)
		group.ALL("/select", hello.Select)
		//group.ALL("/save", hello.Save)
	})

	s.BindHandler("/hello", func(r *ghttp.Request) {
		r.Response.Write("哈喽世界！")
	})

	group := s.Group("/")
	group.POST("/save", hello.Save)

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

}