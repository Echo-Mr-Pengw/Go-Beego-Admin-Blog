package admin

import (
	"Go-Beego-Admin-Blog/models"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"strings"
)

type AuthController struct {
	BaseController
}

func (a *AuthController) List() {
	o := orm.NewOrm()
	auth := []models.AuthList{}
	_, _ = o.QueryTable(new(models.AuthList)).All(&auth)

	a.Data["authList"] = auth
	a.TplName = "admin/auth/list.html"
}

func (a * AuthController) Add() {

	if a.IsPost() {
		authName := a.GetString("authname", "")
		ctlAction := a.GetString("url", "")

		if authName == "" || ctlAction == "" {
			a.ResponseJson(1001, "", "权限名或者URL不能为空")
		}

		// 插入操作
		o := orm.NewOrm()
		auth := models.AuthList{}
		auth.Authname = authName
		auth.Ctlaction = ctlAction
		if _, err := o.Insert(&auth); err != nil {
			a.ResponseJson(1002, "", "添加失败")
		}
		a.ResponseJson(0, "", "添加成功")
	}

	a.TplName = "admin/auth/add.html"
}

func (a * AuthController) Edit() {

	var authId uint
	a.Ctx.Input.Bind(&authId, "id")

	// 去数据库获取相应的标签信息
	o := orm.NewOrm()
	auth := models.AuthList{}
	// tag.Id = tagId
	auth.Id = authId
	err := o.Read(&auth)
	if err == orm.ErrNoRows {
		a.ResponseJson(2004, "", "编辑失败，没有此权限列表")
	}

	a.Data["auth"] = auth
	a.TplName = "admin/auth/add.html"
}

func (a *AuthController) GroupList() {

	o := orm.NewOrm()
	group := []models.AuthGroup{}
	_, _ = o.QueryTable(new(models.AuthGroup)).All(&group)

	a.Data["groupList"] = group
	a.TplName = "admin/auth/groupls.html"
}

func (a *AuthController) GroupAdd() {

	if a.IsPost() {

		var groupName string
		authId := make([]string, 0)

		for i, v := range a.Input() {
			if strings.HasPrefix(i, "authid") {
				authId = append(authId, v[0])
			}else{
				groupName = v[0]
			}
		}

		json, _ := json.Marshal(authId)

		o := orm.NewOrm()
		g := models.AuthGroup{}
		g.Groupname = groupName
		g.Authid = string(json)
		if _, err := o.Insert(&g); err != nil {
			a.ResponseJson(2002, "", "添加失败")
		}
		a.ResponseJson(0, "", "添加成功")
	}

	o := orm.NewOrm()
	auth := []models.AuthList{}
	_, _ = o.QueryTable(new(models.AuthList)).All(&auth)

	a.Data["authList"] = auth
	a.TplName = "admin/auth/groupadd.html"
}

func (a *AuthController) GroupEdit() {

	var groupId uint
	a.Ctx.Input.Bind(&groupId, "id")

	// 去数据库获取相应的标签信息
	o := orm.NewOrm()
	g := models.AuthGroup{}
	g.Id = groupId
	err := o.Read(&g)
	if err == orm.ErrNoRows {
		a.ResponseJson(2004, "", "编辑失败，没有此权限列表")
	}

	authIds := g.Authid
	sliceAuthIds := make([]string, 10)
	json.Unmarshal([]byte(authIds), &sliceAuthIds)

	auth := []models.AuthList{}
	_, _ = o.QueryTable(new(models.AuthList)).All(&auth)

	a.Data["sliceAuthIds"] = sliceAuthIds
	a.Data["authList"] = auth
	a.TplName = "admin/auth/groupadd.html"
}