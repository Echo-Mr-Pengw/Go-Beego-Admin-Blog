// 后台用户
package models

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id int `json:"id", orm:"description(主键),pk"`
	Username string `json:"username, "orm:"size(64); description(用户名)"`
	Password string `json:"password, "orm:"type(char); size(32); description(密码)"`
	Status   uint8   `json:"status,   "orm:"size(1); description(员工状态0:正常 1:停用)"`
	Addtime time.Time `json:"addtime, "orm:"auto_now_add; type(datetime); description(添加时间)"`
}