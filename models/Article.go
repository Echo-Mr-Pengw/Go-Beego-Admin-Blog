// 文章列表
package models

import "time"

type Article struct {
	Id uint `json:"id" orm:"pk, description(主键)"`                                                 // 自增主键
	Tagid uint16 `json:"tagid" orm:"description(标签id)"`                                           // 标签id
	Title string `json:"title" orm:"description(文章标题)"`                                         // 文章标题
	Content string `json:"content" orm:" type(text); description(文章内容)"`                        // 文章内容
	Artstatus uint8 `json:"artstatus" orm:"size(1); default(0); description(状态：0正常1失效)"`     // 文章状态 0正常1失效
	Addtime time.Time `json:"addtime" orm:"auto_now_add; type(datetime); description(添加时间)"`    // 添加时间
}