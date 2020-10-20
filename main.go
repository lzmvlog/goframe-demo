package main

import (
	_ "goframe/boot"
	_ "goframe/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
