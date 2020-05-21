// 文章模块
package admin

import (
	"Go-Beego-Admin-Blog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ArticleController struct {
	BaseController
}

func(a *ArticleController) List() {

	o := orm.NewOrm()
	art := []models.Article{}
	_, _ = o.QueryTable(new(models.Article)).All(&art)

	a.Data["artList"] = art
	a.TplName = "admin/article/list.html"
}

func (a *ArticleController) Add() {

	if a.IsPost() {
		artTitle := a.GetString("arttitle", "")
		artContent := a.GetString("content", "")
		tagId, _ := a.GetUint16("tagstatus", 0)

		if artTitle == "" || artContent == "" || tagId == 0 {
			a.ResponseJson(1001, "", "文章标题或者文章内容不能为空")
		}

		// 插入操作
		o := orm.NewOrm()
		art := models.Article{}
		art.Tagid = tagId
		art.Title = artTitle
		art.Content = artContent
		if _, err := o.Insert(&art); err != nil {
			beego.Info(err, 2121)
			a.ResponseJson(1002, "", "添加失败")
		}
		a.ResponseJson(0, "", "添加成功")
	}

	var tag []*models.Tag
	o := orm.NewOrm()
	o.QueryTable(new(models.Tag)).Filter("tagstatus", 0).All(&tag, "id", "tagname")

	a.Data["tag"] = tag
	a.TplName = "admin/article/add.html"
}

func (a *ArticleController) Edit() {

	var artId uint
	a.Ctx.Input.Bind(&artId, "id")

	// 去数据库获取相应的标签信息
	o := orm.NewOrm()
	art := models.Article{}
	// tag.Id = tagId
	art.Id = artId
	err := o.Read(&art)
	if err == orm.ErrNoRows {
		a.ResponseJson(2004, "", "编辑失败，没有此文章")
	}

	var tag []*models.Tag
	o.QueryTable(new(models.Tag)).Filter("tagstatus", 0).All(&tag, "id", "tagname")

	a.Data["tag"] = tag
	a.Data["art"] = art
	beego.Info(art)
	a.TplName = "admin/article/edit.html"
}