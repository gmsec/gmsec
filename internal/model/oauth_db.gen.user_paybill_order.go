package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type _UserPaybillOrderMgr struct {
	*_BaseMgr
}

// UserPaybillOrderMgr open func
func UserPaybillOrderMgr(db *gorm.DB) *_UserPaybillOrderMgr {
	if db == nil {
		panic(fmt.Errorf("UserPaybillOrderMgr need init by db"))
	}
	return &_UserPaybillOrderMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: gloabIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserPaybillOrderMgr) GetTableName() string {
	return "user_paybill_order"
}

// Get 获取
func (obj *_UserPaybillOrderMgr) Get() (result UserPaybillOrder, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserPaybillOrderMgr) Gets() (results []*UserPaybillOrder, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UserPaybillOrderMgr) WithID(ID int) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithPaybillID paybill_id获取 二次账单id
func (obj *_UserPaybillOrderMgr) WithPaybillID(PaybillID int64) Option {
	return optionFunc(func(o *options) { o.query["paybill_id"] = PaybillID })
}

// WithOrderIDMysql order_id_mysql获取 MySql中的订单Id
func (obj *_UserPaybillOrderMgr) WithOrderIDMysql(OrderIDMysql int64) Option {
	return optionFunc(func(o *options) { o.query["order_id_mysql"] = OrderIDMysql })
}

// WithOrderItemIDMysql order_item_id_mysql获取 MySql中的订单ItemId
func (obj *_UserPaybillOrderMgr) WithOrderItemIDMysql(OrderItemIDMysql int64) Option {
	return optionFunc(func(o *options) { o.query["order_item_id_mysql"] = OrderItemIDMysql })
}

// WithOrderIDMssql order_id_mssql获取 MsSql中的订单Id
func (obj *_UserPaybillOrderMgr) WithOrderIDMssql(OrderIDMssql int64) Option {
	return optionFunc(func(o *options) { o.query["order_id_mssql"] = OrderIDMssql })
}

// GetByOption 功能选项模式获取
func (obj *_UserPaybillOrderMgr) GetByOption(opts ...Option) (result UserPaybillOrder, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UserPaybillOrderMgr) GetByOptions(opts ...Option) (results []*UserPaybillOrder, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_UserPaybillOrderMgr) GetFromID(ID int) (result UserPaybillOrder, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_UserPaybillOrderMgr) GetBatchFromID(IDs []int) (results []*UserPaybillOrder, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

// GetFromPaybillID 通过paybill_id获取内容 二次账单id
func (obj *_UserPaybillOrderMgr) GetFromPaybillID(PaybillID int64) (results []*UserPaybillOrder, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("paybill_id = ?", PaybillID).Find(&results).Error

	return
}

// GetBatchFromPaybillID 批量唯一主键查找 二次账单id
func (obj *_UserPaybillOrderMgr) GetBatchFromPaybillID(PaybillIDs []int64) (results []*UserPaybillOrder, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("paybill_id IN (?)", PaybillIDs).Find(&results).Error

	return
}

// GetFromOrderIDMysql 通过order_id_mysql获取内容 MySql中的订单Id
func (obj *_UserPaybillOrderMgr) GetFromOrderIDMysql(OrderIDMysql int64) (results []*UserPaybillOrder, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("order_id_mysql = ?", OrderIDMysql).Find(&results).Error

	return
}

// GetBatchFromOrderIDMysql 批量唯一主键查找 MySql中的订单Id
func (obj *_UserPaybillOrderMgr) GetBatchFromOrderIDMysql(OrderIDMysqls []int64) (results []*UserPaybillOrder, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("order_id_mysql IN (?)", OrderIDMysqls).Find(&results).Error

	return
}

// GetFromOrderItemIDMysql 通过order_item_id_mysql获取内容 MySql中的订单ItemId
func (obj *_UserPaybillOrderMgr) GetFromOrderItemIDMysql(OrderItemIDMysql int64) (results []*UserPaybillOrder, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("order_item_id_mysql = ?", OrderItemIDMysql).Find(&results).Error

	return
}

// GetBatchFromOrderItemIDMysql 批量唯一主键查找 MySql中的订单ItemId
func (obj *_UserPaybillOrderMgr) GetBatchFromOrderItemIDMysql(OrderItemIDMysqls []int64) (results []*UserPaybillOrder, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("order_item_id_mysql IN (?)", OrderItemIDMysqls).Find(&results).Error

	return
}

// GetFromOrderIDMssql 通过order_id_mssql获取内容 MsSql中的订单Id
func (obj *_UserPaybillOrderMgr) GetFromOrderIDMssql(OrderIDMssql int64) (results []*UserPaybillOrder, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("order_id_mssql = ?", OrderIDMssql).Find(&results).Error

	return
}

// GetBatchFromOrderIDMssql 批量唯一主键查找 MsSql中的订单Id
func (obj *_UserPaybillOrderMgr) GetBatchFromOrderIDMssql(OrderIDMssqls []int64) (results []*UserPaybillOrder, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("order_id_mssql IN (?)", OrderIDMssqls).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_UserPaybillOrderMgr) FetchByPrimaryKey(ID int) (result UserPaybillOrder, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// FetchByOrderIDIndex  获取多个内容
func (obj *_UserPaybillOrderMgr) FetchByOrderIDIndex(PaybillID int64, OrderIDMysql int64, OrderItemIDMysql int64) (results []*UserPaybillOrder, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("paybill_id = ? AND order_id_mysql = ? AND order_item_id_mysql = ?", PaybillID, OrderIDMysql, OrderItemIDMysql).Find(&results).Error

	return
}
