// 文章标签
package models

import "time"

type Tag struct {
	Id uint `json:"id" orm:"description(主键); pk"`                                              // 自增主键
	Tagname string `json:"tagname" orm:"description(标签名)"`                                    // 标签名
	Tagstatus uint8 `json:"tagstatus" orm:"size(1); default(0); description(状态：0正常1失效)"`   // 标签状态 0正常1失效
	Addtime time.Time `json:"addtime" orm:"auto_now_add; type(datetime); description(添加时间)"`  // 添加时间
}