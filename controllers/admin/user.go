// 后台用户登录、注销
package admin

import (
	"Go-Beego-Admin-Blog/models"
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type UserController struct {
	// 继承基类
	BaseController
}

// 登录
func (c *UserController) Login() {

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

// 登录成功后的主界面
func (c * UserController) Index() {
	//c.Ctx.WriteString("1111")
	c.TplName = "admin/index.html"
}

// 用户列表
func (c *UserController) List() {

	user := []models.User{}
	o := orm.NewOrm()
	_, _ = o.QueryTable("user").All(&user)

	// c.Data["statusDescrip"]   = AdminUserStatus
	c.Data["userList"] = user
	c.TplName = "admin/user/list.html"
}

// 用户添加
func (c *UserController) Add() {

	// 如果是POST，表示添加
	if c.IsPost() {

		userName := c.GetString("username", "")
		passWd   := c.GetString("pass", "")
		rePassWd := c.GetString("repass", "")

		// 用户名不能为空
		if userName == "" {
			c.ResponseJson(1001, "", "用户名不能为空")
		}

		// 密码不能为空
		if passWd == "" || rePassWd == "" {
			c.ResponseJson(1002, "", "密码不能为空")
		}

		// 密码不一致
		if passWd != rePassWd {
			c.ResponseJson(1003, "", "两次密码不一致")
		}

		// 判断用户是否已经存在
		o := orm.NewOrm()
		user := models.User{}
		user.Username = userName
		err := o.Read(&user, "username")
		if err ==  nil {
			c.ResponseJson(1004, "", "用户已经存在")
		}
		// 不存在则添加
		user.Password = passWd
		_, err = o.Insert(&user)
		if err != nil {
			c.ResponseJson(1005, "", "用户添加失败")
		}
		c.ResponseJson(0, "", "用户添加成功")
	}
	c.TplName = "admin/user/add.html"
}

// 用户编辑
func (c *UserController) Edit() {

	// 编辑POST提交
	if c.IsPost() {
		id, _ := c.GetInt("id", 0)
		userName := c.GetString("username", "")
		passWord := c.GetString("pass", "")

		if id == 0 || userName == "" || passWord == "" {
			c.ResponseJson(1001, "", "用户名或者密码不能为空")
		}

		o := orm.NewOrm()
		u := models.User{Id: id}
		if o.Read(&u) != nil {
			c.ResponseJson(1007, "", "编辑失败")
		}

		u.Username = userName
		u.Password = passWord
		if _, err := o.Update(&u, "UserName", "Password"); err != nil {
			c.ResponseJson(1007, "", "编辑失败")
		}
		c.ResponseJson(0, "", "编辑成功")
	}

	// 获取用户对应的id
	var userId int
	c.Ctx.Input.Bind(&userId, "id")

	// 去数据库获取相应的用户信息
	o := orm.NewOrm()
	u := models.User{}
	u.Id = userId
	err := o.Read(&u)
	if err == orm.ErrNoRows {
		c.ResponseJson(1006, "", "编辑失败，没有此用户")
	}

	c.Data["id"]       = u.Id
	c.Data["userName"] = u.Username
	// c.Data["passWord"] = u.Password
	// c.Data["status"]   = u.Status
	beego.Info(c.Data)
	c.TplName = "admin/user/edit.html"
}

// POST删除用户
func (c *UserController) Del() {

	userId, _ := c.GetInt("id")
	if userId == 0 {
		c.ResponseJson(1008, "", "删除失败")
	}

	o := orm.NewOrm()
	u := models.User{Id:userId}
	u.Status = 1
	if _, err := o.Update(&u, "Status"); err != nil {
		c.ResponseJson(1008, "", "删除失败")
	}
	c.ResponseJson(0, "", "删除成功")
}

// 注销
func (c *UserController) Logout() {
	// 删除cookie
	c.Ctx.SetCookie("uid", "")
	c.Ctx.SetCookie("token", "")
	// 删除session
	c.DestroySession()

	// 跳转到后台登录
	c.Redirect("/admin", 302)
}