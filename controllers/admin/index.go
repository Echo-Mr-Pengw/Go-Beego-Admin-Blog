// 后台登录界面
package admin

type IndexController struct {
	// 继承基类
	BaseController
}

// 进去后台登录页面
func (c *IndexController) Index() {
	//c.Ctx.WriteString("111")
	c.TplName = "admin/login.html"
}