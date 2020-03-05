package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type _SignClientTblMgr struct {
	*_BaseMgr
}

// SignClientTblMgr open func
func SignClientTblMgr(db *gorm.DB) *_SignClientTblMgr {
	if db == nil {
		panic(fmt.Errorf("SignClientTblMgr need init by db"))
	}
	return &_SignClientTblMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: gloabIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SignClientTblMgr) GetTableName() string {
	return "sign_client_tbl"
}

// Get 获取
func (obj *_SignClientTblMgr) Get() (result SignClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SignClientTblMgr) Gets() (results []*SignClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_SignClientTblMgr) WithID(ID int) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithAppKey app_key获取
func (obj *_SignClientTblMgr) WithAppKey(AppKey string) Option {
	return optionFunc(func(o *options) { o.query["app_key"] = AppKey })
}

// WithAppSecret app_secret获取
func (obj *_SignClientTblMgr) WithAppSecret(AppSecret string) Option {
	return optionFunc(func(o *options) { o.query["app_secret"] = AppSecret })
}

// WithExpireTime expire_time获取 超时时间
func (obj *_SignClientTblMgr) WithExpireTime(ExpireTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["expire_time"] = ExpireTime })
}

// WithStrictSign strict_sign获取 是否强制验签:0：用户自定义，1：强制
func (obj *_SignClientTblMgr) WithStrictSign(StrictSign int) Option {
	return optionFunc(func(o *options) { o.query["strict_sign"] = StrictSign })
}

// WithStrictVerify strict_verify获取 是否强制验证码:0：用户自定义，1：强制
func (obj *_SignClientTblMgr) WithStrictVerify(StrictVerify int) Option {
	return optionFunc(func(o *options) { o.query["strict_verify"] = StrictVerify })
}

// WithTokenExpireTime token_expire_time获取 token过期时间
func (obj *_SignClientTblMgr) WithTokenExpireTime(TokenExpireTime int) Option {
	return optionFunc(func(o *options) { o.query["token_expire_time"] = TokenExpireTime })
}

// GetByOption 功能选项模式获取
func (obj *_SignClientTblMgr) GetByOption(opts ...Option) (result SignClientTbl, err error) {
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
func (obj *_SignClientTblMgr) GetByOptions(opts ...Option) (results []*SignClientTbl, err error) {
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
func (obj *_SignClientTblMgr) GetFromID(ID int) (result SignClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_SignClientTblMgr) GetBatchFromID(IDs []int) (results []*SignClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

// GetFromAppKey 通过app_key获取内容
func (obj *_SignClientTblMgr) GetFromAppKey(AppKey string) (results []*SignClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_key = ?", AppKey).Find(&results).Error

	return
}

// GetBatchFromAppKey 批量唯一主键查找
func (obj *_SignClientTblMgr) GetBatchFromAppKey(AppKeys []string) (results []*SignClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_key IN (?)", AppKeys).Find(&results).Error

	return
}

// GetFromAppSecret 通过app_secret获取内容
func (obj *_SignClientTblMgr) GetFromAppSecret(AppSecret string) (results []*SignClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_secret = ?", AppSecret).Find(&results).Error

	return
}

// GetBatchFromAppSecret 批量唯一主键查找
func (obj *_SignClientTblMgr) GetBatchFromAppSecret(AppSecrets []string) (results []*SignClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_secret IN (?)", AppSecrets).Find(&results).Error

	return
}

// GetFromExpireTime 通过expire_time获取内容 超时时间
func (obj *_SignClientTblMgr) GetFromExpireTime(ExpireTime time.Time) (results []*SignClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("expire_time = ?", ExpireTime).Find(&results).Error

	return
}

// GetBatchFromExpireTime 批量唯一主键查找 超时时间
func (obj *_SignClientTblMgr) GetBatchFromExpireTime(ExpireTimes []time.Time) (results []*SignClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("expire_time IN (?)", ExpireTimes).Find(&results).Error

	return
}

// GetFromStrictSign 通过strict_sign获取内容 是否强制验签:0：用户自定义，1：强制
func (obj *_SignClientTblMgr) GetFromStrictSign(StrictSign int) (results []*SignClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("strict_sign = ?", StrictSign).Find(&results).Error

	return
}

// GetBatchFromStrictSign 批量唯一主键查找 是否强制验签:0：用户自定义，1：强制
func (obj *_SignClientTblMgr) GetBatchFromStrictSign(StrictSigns []int) (results []*SignClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("strict_sign IN (?)", StrictSigns).Find(&results).Error

	return
}

// GetFromStrictVerify 通过strict_verify获取内容 是否强制验证码:0：用户自定义，1：强制
func (obj *_SignClientTblMgr) GetFromStrictVerify(StrictVerify int) (results []*SignClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("strict_verify = ?", StrictVerify).Find(&results).Error

	return
}

// GetBatchFromStrictVerify 批量唯一主键查找 是否强制验证码:0：用户自定义，1：强制
func (obj *_SignClientTblMgr) GetBatchFromStrictVerify(StrictVerifys []int) (results []*SignClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("strict_verify IN (?)", StrictVerifys).Find(&results).Error

	return
}

// GetFromTokenExpireTime 通过token_expire_time获取内容 token过期时间
func (obj *_SignClientTblMgr) GetFromTokenExpireTime(TokenExpireTime int) (results []*SignClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("token_expire_time = ?", TokenExpireTime).Find(&results).Error

	return
}

// GetBatchFromTokenExpireTime 批量唯一主键查找 token过期时间
func (obj *_SignClientTblMgr) GetBatchFromTokenExpireTime(TokenExpireTimes []int) (results []*SignClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("token_expire_time IN (?)", TokenExpireTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SignClientTblMgr) FetchByPrimaryKey(ID int) (result SignClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}
