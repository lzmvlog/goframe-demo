package hello

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// Hello is a demonstration route handler for output "Hello World!".
func Hello(r *ghttp.Request) {
	db := g.DB()
	list, err := db.Query("select * from student")
	r.Response.Writeln("Hello World!", list, err)
}

func Test(r *ghttp.Request) {
	db := g.DB()
	list, _ := db.GetAll("select * from student")
	r.Response.WriteJson(list)
}

func Select(r *ghttp.Request) {
	id := r.Get("id")
	db := g.DB()
	list, _ := db.GetOne("select * from student where id=?", id)
	r.Response.WriteJson(list)
}

type RegisterRes struct {
	Code  int         `json:"code"`
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

func Save(r *ghttp.Request) {
	// form-data
	//id := r.GetForm("id")
	// url
	//id := r.GetQuery("id")

	// 获取任意提交类型的值
	//id := r.GetRequest("id")
	//name := r.GetRequest("name")
	//id := r.GetBody()
	//id := r.GetStruct(r.Parse(&stud))

	//var stud *model.Student
	//if err := r.Parse(&stud); err != nil {
	//	r.Response.WriteJsonExit(RegisterRes{
	//		Code:  1,
	//		Error: err.Error(),
	//	})
	//}
	//r.Response.WriteJsonExit(RegisterRes{
	//	Data: stud,
	//})
	//r.Response.WriteJson(r.Parse(&stud))

	//stu := r.GetBodyString()
	//var stu *student.Entity
	//name := r.GetRequest("name").(string)
	//id := r.GetRequest("id")
	////fmt.Println(name.(string))
	//db := g.DB()
	//db.Insert("student", student.Entity{
	//	Id:      id.(string),
	//	Name:    name,
	//})
	//if err := r.Parse(&stu); err != nil {
	//	// Validation error.
	//	if v, ok := err.(*gvalid.Error); ok {
	//		r.Response.WriteJsonExit(RegisterRes{
	//			Code:  1,
	//			Error: v.FirstString(),
	//		})
	//	}
	//	// Other error.
	//	r.Response.WriteJsonExit(RegisterRes{
	//		Code:  1,
	//		Error: err.Error(),
	//	})
	//}
	//// ...
	//r.Response.WriteJsonExit(RegisterRes{
	//	Data: stu,
	//})

	//var data *student.Entity
	//// 这里没有使用Parse而是仅用GetStruct获取对象，
	//// 数据校验交给后续的service层统一处理。
	//if err := r.GetStruct(&data); err != nil {
	//	response.JsonExit(r, 1, err.Error())
	//}
}

func TestTpl(r *ghttp.Request) {
	////tplContent := `id:{{.id}}, name:{{.name}}`
	r.Response.WriteTpl("index.html", g.Map{
		"id":   123,
		"name": "john",
	})
	//r.Response.WriteTplContent(tplContent, g.Map{
	//	"id"   : 123,
	//	"name" : "john",
	//})
}
