package main

import (
	_ "Go-Beego-Admin-Blog/models"
	_ "Go-Beego-Admin-Blog/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

