package routers

import (
	"Go-Beego-Admin-Blog/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {

	// 后台路由
	// 后台登录页面
    beego.Router("/admin", &admin.IndexController{}, "*:Index")
    // 后台登录请求接口
    beego.Router("/admin/login", &admin.UserController{}, "post:Login")
    // 后台登录成功后的主页面
    beego.Router("/admin/index", &admin.UserController{}, "*:Index")
    // 后台用户注销接口
    beego.Router("/admin/logout", &admin.UserController{}, "get:Logout")

    // 后台用户管理模块
    beego.Router("/admin/user/ls" , &admin.UserController{}, "get:List")
    beego.Router("/admin/user/add", &admin.UserController{}, "*:Add")
    beego.Router("/admin/user/edit", &admin.UserController{}, "get,post:Edit")
    beego.Router("/admin/user/del", &admin.UserController{}, "post:Del")
}
