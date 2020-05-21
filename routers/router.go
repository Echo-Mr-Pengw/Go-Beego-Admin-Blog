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
    beego.Router("/admin/login", &admin.IndexController{}, "post:Login")
    // 后台登录成功后的主页面
    beego.Router("/admin/index", &admin.UserController{}, "*:Index")
    // 后台用户注销接口
    beego.Router("/admin/logout", &admin.UserController{}, "get:Logout")

    // 后台用户管理模块
    beego.Router("/admin/user/ls" , &admin.UserController{}, "get:List")
    beego.Router("/admin/user/add", &admin.UserController{}, "*:Add")
    beego.Router("/admin/user/edit", &admin.UserController{}, "get,post:Edit")
    beego.Router("/admin/user/del", &admin.UserController{}, "post:Del")

    // 标签模块
    beego.Router("/admin/tag/ls", &admin.TagController{}, "get:List")
    beego.Router("/admin/tag/add", &admin.TagController{}, "get,post:Add")
    beego.Router("/admin/tag/edit", &admin.TagController{}, "get,post:Edit")

    // 文章模块
    beego.Router("/admin/article/ls", &admin.ArticleController{}, "get:List")
    beego.Router("/admin/article/add", &admin.ArticleController{}, "get,post:Add")
    beego.Router("/admin/article/edit", &admin.ArticleController{}, "get,post:Edit")

    // 权限模块
    beego.Router("/admin/auth/ls", &admin.AuthController{}, "get:List")
    beego.Router("/admin/auth/add", &admin.AuthController{}, "get,post:Add")
    beego.Router("/admin/auth/edit", &admin.AuthController{}, "get,post:Edit")
    beego.Router("/admin/auth/groupls", &admin.AuthController{}, "get:GroupList")
    beego.Router("/admin/auth/groupadd", &admin.AuthController{}, "get,post:GroupAdd")
    beego.Router("/admin/auth/groupedit", &admin.AuthController{}, "get,post:GroupEdit")
}
