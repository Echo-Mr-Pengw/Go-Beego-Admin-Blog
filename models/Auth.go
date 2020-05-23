// 权限模块
package models

import "time"

// 权限列表
type AuthList struct {
	Id uint `json:"id" orm:"description(主键),pk"`                                                   // 自增主键(权限id)
	Authname string `json:"authname" orm:"description(权限列表名)"`                                  // 权限列表名
	Ctlaction string `json:"ctlaction" orm:"description(控制器方法)"`                                // 对应请求方法
	Addtime time.Time `json:"addtime, "orm:"auto_now_add; type(datetime); description(添加时间)"`    // 添加时间
}

// 权限分组
type AuthGroup struct {
	Id uint `json:"id" orm:"description(主键),pk"`                                                 // 主键
	Authid string `json:"authid" orm:"description(权限id)"`                                        // 权限id
	Groupname string `json:"groupname" orm:"description(分组名)"`                                  // 分组名
	Addtime time.Time `json:"addtime" orm:"auto_now_add; type(datetime); description(添加时间)"`  // 添加时间
}
