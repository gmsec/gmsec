package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type _Oauth2ClientTblMgr struct {
	*_BaseMgr
}

// Oauth2ClientTblMgr open func
func Oauth2ClientTblMgr(db *gorm.DB) *_Oauth2ClientTblMgr {
	if db == nil {
		panic(fmt.Errorf("Oauth2ClientTblMgr need init by db"))
	}
	return &_Oauth2ClientTblMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: gloabIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Oauth2ClientTblMgr) GetTableName() string {
	return "oauth2_client_tbl"
}

// Get 获取
func (obj *_Oauth2ClientTblMgr) Get() (result Oauth2ClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_Oauth2ClientTblMgr) Gets() (results []*Oauth2ClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_Oauth2ClientTblMgr) WithID(ID int) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithAppKey app_key获取
func (obj *_Oauth2ClientTblMgr) WithAppKey(AppKey string) Option {
	return optionFunc(func(o *options) { o.query["app_key"] = AppKey })
}

// WithAppSecret app_secret获取
func (obj *_Oauth2ClientTblMgr) WithAppSecret(AppSecret string) Option {
	return optionFunc(func(o *options) { o.query["app_secret"] = AppSecret })
}

// WithExpireTime expire_time获取 超时时间
func (obj *_Oauth2ClientTblMgr) WithExpireTime(ExpireTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["expire_time"] = ExpireTime })
}

// WithStrictSign strict_sign获取 是否强制验签:0：用户自定义，1：强制
func (obj *_Oauth2ClientTblMgr) WithStrictSign(StrictSign int) Option {
	return optionFunc(func(o *options) { o.query["strict_sign"] = StrictSign })
}

// WithStrictVerify strict_verify获取 是否强制验证码:0：用户自定义，1：强制
func (obj *_Oauth2ClientTblMgr) WithStrictVerify(StrictVerify int) Option {
	return optionFunc(func(o *options) { o.query["strict_verify"] = StrictVerify })
}

// WithTokenExpireTime token_expire_time获取 token过期时间
func (obj *_Oauth2ClientTblMgr) WithTokenExpireTime(TokenExpireTime int) Option {
	return optionFunc(func(o *options) { o.query["token_expire_time"] = TokenExpireTime })
}

// WithAaa aaa获取
func (obj *_Oauth2ClientTblMgr) WithAaa(Aaa string) Option {
	return optionFunc(func(o *options) { o.query["aaa"] = Aaa })
}

// GetByOption 功能选项模式获取
func (obj *_Oauth2ClientTblMgr) GetByOption(opts ...Option) (result Oauth2ClientTbl, err error) {
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
func (obj *_Oauth2ClientTblMgr) GetByOptions(opts ...Option) (results []*Oauth2ClientTbl, err error) {
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
func (obj *_Oauth2ClientTblMgr) GetFromID(ID int) (result Oauth2ClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_Oauth2ClientTblMgr) GetBatchFromID(IDs []int) (results []*Oauth2ClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

// GetFromAppKey 通过app_key获取内容
func (obj *_Oauth2ClientTblMgr) GetFromAppKey(AppKey string) (results []*Oauth2ClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_key = ?", AppKey).Find(&results).Error

	return
}

// GetBatchFromAppKey 批量唯一主键查找
func (obj *_Oauth2ClientTblMgr) GetBatchFromAppKey(AppKeys []string) (results []*Oauth2ClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_key IN (?)", AppKeys).Find(&results).Error

	return
}

// GetFromAppSecret 通过app_secret获取内容
func (obj *_Oauth2ClientTblMgr) GetFromAppSecret(AppSecret string) (results []*Oauth2ClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_secret = ?", AppSecret).Find(&results).Error

	return
}

// GetBatchFromAppSecret 批量唯一主键查找
func (obj *_Oauth2ClientTblMgr) GetBatchFromAppSecret(AppSecrets []string) (results []*Oauth2ClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_secret IN (?)", AppSecrets).Find(&results).Error

	return
}

// GetFromExpireTime 通过expire_time获取内容 超时时间
func (obj *_Oauth2ClientTblMgr) GetFromExpireTime(ExpireTime time.Time) (results []*Oauth2ClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("expire_time = ?", ExpireTime).Find(&results).Error

	return
}

// GetBatchFromExpireTime 批量唯一主键查找 超时时间
func (obj *_Oauth2ClientTblMgr) GetBatchFromExpireTime(ExpireTimes []time.Time) (results []*Oauth2ClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("expire_time IN (?)", ExpireTimes).Find(&results).Error

	return
}

// GetFromStrictSign 通过strict_sign获取内容 是否强制验签:0：用户自定义，1：强制
func (obj *_Oauth2ClientTblMgr) GetFromStrictSign(StrictSign int) (results []*Oauth2ClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("strict_sign = ?", StrictSign).Find(&results).Error

	return
}

// GetBatchFromStrictSign 批量唯一主键查找 是否强制验签:0：用户自定义，1：强制
func (obj *_Oauth2ClientTblMgr) GetBatchFromStrictSign(StrictSigns []int) (results []*Oauth2ClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("strict_sign IN (?)", StrictSigns).Find(&results).Error

	return
}

// GetFromStrictVerify 通过strict_verify获取内容 是否强制验证码:0：用户自定义，1：强制
func (obj *_Oauth2ClientTblMgr) GetFromStrictVerify(StrictVerify int) (results []*Oauth2ClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("strict_verify = ?", StrictVerify).Find(&results).Error

	return
}

// GetBatchFromStrictVerify 批量唯一主键查找 是否强制验证码:0：用户自定义，1：强制
func (obj *_Oauth2ClientTblMgr) GetBatchFromStrictVerify(StrictVerifys []int) (results []*Oauth2ClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("strict_verify IN (?)", StrictVerifys).Find(&results).Error

	return
}

// GetFromTokenExpireTime 通过token_expire_time获取内容 token过期时间
func (obj *_Oauth2ClientTblMgr) GetFromTokenExpireTime(TokenExpireTime int) (results []*Oauth2ClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("token_expire_time = ?", TokenExpireTime).Find(&results).Error

	return
}

// GetBatchFromTokenExpireTime 批量唯一主键查找 token过期时间
func (obj *_Oauth2ClientTblMgr) GetBatchFromTokenExpireTime(TokenExpireTimes []int) (results []*Oauth2ClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("token_expire_time IN (?)", TokenExpireTimes).Find(&results).Error

	return
}

// GetFromAaa 通过aaa获取内容
func (obj *_Oauth2ClientTblMgr) GetFromAaa(Aaa string) (results []*Oauth2ClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("aaa = ?", Aaa).Find(&results).Error

	return
}

// GetBatchFromAaa 批量唯一主键查找
func (obj *_Oauth2ClientTblMgr) GetBatchFromAaa(Aaas []string) (results []*Oauth2ClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("aaa IN (?)", Aaas).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_Oauth2ClientTblMgr) FetchByPrimaryKey(ID int) (result Oauth2ClientTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}
