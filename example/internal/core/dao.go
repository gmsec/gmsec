package core

import (
	"example/internal/config"

	"github.com/xxjwxc/public/errors"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/mysqldb"
)

// Dao ...
var Dao DaoCore

// DaoCore core dao
type DaoCore struct {
	dbr *mysqldb.MySqlDB // 读库
	dbw *mysqldb.MySqlDB // 写库
}

func init() {
	Dao.InitDao()
}

// GetDBr 获取读数据库
func (dao *DaoCore) GetDBr() *mysqldb.MySqlDB {
	if dao.dbr == nil {
		mylog.Error(errors.New("dbr is nil. "))
	}
	return dao.dbr
}

// GetDBw 获取写数据库
func (dao *DaoCore) GetDBw() *mysqldb.MySqlDB {
	if dao.dbw == nil {
		mylog.Error(errors.New("dbw is nil. "))
	}
	return dao.dbw
}

// InitDao 初始化dao
func (dao *DaoCore) InitDao() {
	// 原实现有 runtime.SetFinalizer(dao, dao.Destroy):
	// dao 是包级全局变量,永远不会变成不可达对象,这个 finalizer 一次都不会执行,
	// 反而让人误以为连接会被自动释放。进程退出前需要释放请显式调 Destroy。
	dao.dbr = mysqldb.OnInitDBOrm(config.GetMysqlConStr(), 5, 20, true)
	dao.dbw = mysqldb.OnInitDBOrm(config.GetMysqlConStr(), 5, 20, true)
}

// Destroy 释放
func (dao *DaoCore) Destroy() {
	if dao.dbr != nil {
		dao.dbr.OnDestoryDB()
	}

	if dao.dbw != nil {
		dao.dbw.OnDestoryDB()
	}
}
