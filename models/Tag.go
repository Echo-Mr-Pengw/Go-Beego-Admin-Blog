// 文章标签
package models

type Tag struct {
	Id uint `json:"id", orm:"description(主键),pk"`
}