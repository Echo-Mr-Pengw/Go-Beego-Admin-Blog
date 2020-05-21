package admin

import (
	"Go-Beego-Admin-Blog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type TagController struct {
	BaseController
}

// 标签列表
func (t *TagController) List() {

	tag := []models.Tag{}
	o := orm.NewOrm()
	_, _ = o.QueryTable("tbl_tag").All(&tag)

	t.Data["tagList"] = tag
	t.TplName = "admin/tag/list.html"
}

// 添加标签
func (t *TagController) Add() {

	if t.IsPost() {

		tagName := t.GetString("tagname", "")
		tagStatus, _ := t.GetUint8("tagstatus", 0)

		if tagName == "" {
			t.ResponseJson(2001, "", "标签名不能为空")
		}

		// 插入操作
		o := orm.NewOrm()
		tag := models.Tag{}
		tag.Tagname = tagName
		tag.Tagstatus = tagStatus
		if _, err := o.Insert(&tag); err != nil {
			t.ResponseJson(2002, "", "添加失败")
		}
		t.ResponseJson(0, "", "添加成功")
	}

	t.TplName = "admin/tag/add.html"
}

// 编辑标签
func (t *TagController) Edit() {

	if t.IsPost() {

		id, _ := t.GetUint64("id", 0)
		tagName := t.GetString("tagname", "")
		tagStatus, _ := t.GetUint8("tagstatus", 0)

		if id == 0 {
			t.ResponseJson(2006, "", "参数错误")
		}

		if tagName == "" {
			t.ResponseJson(2005, "", "标签名不能为空")
		}

		o := orm.NewOrm()
		tag := models.Tag{Id: uint(id)}
		if o.Read(&tag) != nil {
			t.ResponseJson(1007, "", "编辑失败")
		}

		tag.Tagname = tagName
		tag.Tagstatus = tagStatus
		if _, err := o.Update(&tag, "Tagname", "Tagstatus"); err != nil {
			t.ResponseJson(1007, "", "编辑失败")
		}
		t.ResponseJson(0, "", "编辑成功")
	}

	var tagId uint
	t.Ctx.Input.Bind(&tagId, "id")

	// 去数据库获取相应的标签信息
	o := orm.NewOrm()
	tag := models.Tag{}
	// tag.Id = tagId
	tag.Id = tagId
	err := o.Read(&tag)
	if err == orm.ErrNoRows {
		t.ResponseJson(2004, "", "编辑失败，没有此标签")
	}

	t.Data["id"]        = tag.Id
	t.Data["tagName"]   = tag.Tagname
	t.Data["tagStatus"] = tag.Tagstatus

	beego.Info(t.Data)

	t.TplName = "admin/tag/edit.html"
}