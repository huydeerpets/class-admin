package main

import (
	_ "class-admin/routers"
	"github.com/astaxie/beego"
	"class-admin/lib"
)

func main() {
	lib.InitOssBucket()
	beego.Run()
}

