package admin

import "github.com/astaxie/beego"

// 参与签名的秘钥
const Key = "eea9d13bc394514eb6838b1b43e93e5b"

// 定义后台用户的状态
var AdminUserStatus = [2]string{0 : "正常", 1 : "失效"}

// 定义JSON返回的消息体
type Resp struct {
	Errno int         `json:"errno"`      // 编号
	Data interface{}  `json:"data"`       // 返回的数据
	Errmsg string     `json:"errmsg"`     // 消息
}

type BaseController struct {
	beego.Controller
}

// 权限认准
func (b *BaseController) Prepare() {
	// 获取客户端传的cookie
	clicentUid   := b.Ctx.GetCookie("uid")
	clientToken := b.Ctx.GetCookie("token")
	beego.Info(clicentUid, clientToken, 1111)
	if clicentUid == "" || clientToken == "" {
		b.Redirect("/admin", 302)
	}

	// 通过uid获取session
	sess := b.GetSession(clicentUid)
	beego.Info(clicentUid, clientToken, sess)
	if sess == nil {
		b.Redirect("/admin", 302)
	}

	// 客户端传的token与服务端不匹配
	if clientToken != sess {
		b.Redirect("/admin", 302)
	}
}

// POST请求判断
func (b *BaseController) IsPost() bool {
	return b.Ctx.Request.Method == "POST"
}

// 返回JSON
func (b * BaseController) ResponseJson(errno int, data interface{}, errmsg string) {
	retJson := &Resp{
		Errno: errno,
		Data: data,
		Errmsg: errmsg,
	}
	b.Data["json"] = retJson
	b.ServeJSON()
}

