package datasource

import (
	"QianfengCmsProject/model"
	_ "github.com/go-sql-driver/mysql" //不能忘记导入
	"github.com/go-xorm/xorm"
)

/**
 * 实例化数据库引擎方法：mysql的数据引擎
 */
func NewMysqlEngine() *xorm.Engine {
	//数据库引擎
	engine, err := xorm.NewEngine("mysql", "root:xxl123@/testcms?charset=utf8")
	err = engine.Sync2(new(model.Admin), new(model.AdminPermission), new(model.City), new(model.Permission))
	if err != nil {
		panic(err.Error())
	}
	// 设置显示sql语句
	engine.ShowSQL(true)
	engine.SetMaxOpenConns(10)
	return engine
}
