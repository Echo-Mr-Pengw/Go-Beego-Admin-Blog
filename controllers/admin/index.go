// 后台登录界面
package admin

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"Go-Beego-Admin-Blog/models"
	"strconv"
)

type IndexController struct {
	// 继承基类
	// BaseController
	beego.Controller
}

// 进去后台登录页面
func (c *IndexController) Index() {
	//c.Ctx.WriteString("111")
	c.TplName = "admin/login.html"
}

// 登录
func (c *IndexController) Login() {

	// 获取用户名和密码
	userName := c.GetString("username", "")
	passWord := c.GetString("password", "")

	// 用户名或者密码为空
	if userName == "" || passWord == "" {
		r := Resp{1001, "", "用户名或者密码不能为空"}
		c.Data["json"] = r
		c.ServeJSON()
		return
	}

	o := orm.NewOrm()
	user := models.User{}
	user.Username = userName
	err := o.Read(&user, "username")
	//用户不存在
	if err != nil {
		r := Resp{1002, "", "用户名不存在，请联系管理员"}
		c.Data["json"] = r
		c.ServeJSON()
		return
	}

	// 用户存在，但密码错误
	if user.Password != passWord {
		r := Resp{1003, "", "密码错误"}
		c.Data["json"] = r
		c.ServeJSON()
		return
	}

	// 设置cookie和session
	id := strconv.Itoa(user.Id)
	signSlice := md5.Sum([]byte(id + user.Username + user.Password + Key))
	signStr := fmt.Sprintf("%x", signSlice)

	// 设置cookie
	c.Ctx.SetCookie("uid", id, 86400)
	c.Ctx.SetCookie("token", signStr, 86400)
	// 设置session
	c.SetSession(id, signStr)

	r := Resp{0, "", "登录成功"}
	c.Data["json"] = r
	c.ServeJSON()
}