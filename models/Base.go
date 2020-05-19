// 数据库初始操作
package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 数据库的初始化
func init() {

	driverName := beego.AppConfig.String("drivername")
	mysqlHost  := beego.AppConfig.String("mysqlhost")
	mysqlPort  := beego.AppConfig.String("mysqlport")
	mysqlUser  := beego.AppConfig.String("mysqluser")
	mysqlPass  := beego.AppConfig.String("mysqlpasswd")
	mysqlDb    := beego.AppConfig.String("mysqldb")

	dataSource := mysqlUser + ":" + mysqlPass + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + mysqlDb + "?charset=utf8"
	// beego.Info(dataSource)
	orm.RegisterModelWithPrefix("tbl_", new(User), new(Tag), new(Article), new(AuthList), new(AuthGroup))
	// orm.RegisterModel(new(User), new(Tag), new(Article))
	orm.RegisterDriver(driverName, orm.DRMySQL)
	orm.RegisterDataBase("default", driverName, dataSource)
	orm.RunSyncdb("default", false, true)
}