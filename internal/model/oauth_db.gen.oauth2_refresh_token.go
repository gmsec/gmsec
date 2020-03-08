package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type _Oauth2RefreshTokenMgr struct {
	*_BaseMgr
}

// Oauth2RefreshTokenMgr open func
func Oauth2RefreshTokenMgr(db *gorm.DB) *_Oauth2RefreshTokenMgr {
	if db == nil {
		panic(fmt.Errorf("Oauth2RefreshTokenMgr need init by db"))
	}
	return &_Oauth2RefreshTokenMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: gloabIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Oauth2RefreshTokenMgr) GetTableName() string {
	return "oauth2_refresh_token"
}

// Get 获取
func (obj *_Oauth2RefreshTokenMgr) Get() (result Oauth2RefreshToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_Oauth2RefreshTokenMgr) Gets() (results []*Oauth2RefreshToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_Oauth2RefreshTokenMgr) WithID(ID int) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithRefreshToken refresh_token获取
func (obj *_Oauth2RefreshTokenMgr) WithRefreshToken(RefreshToken string) Option {
	return optionFunc(func(o *options) { o.query["refresh_token"] = RefreshToken })
}

// WithTokenType token_type获取
func (obj *_Oauth2RefreshTokenMgr) WithTokenType(TokenType string) Option {
	return optionFunc(func(o *options) { o.query["token_type"] = TokenType })
}

// WithAppKey app_key获取
func (obj *_Oauth2RefreshTokenMgr) WithAppKey(AppKey string) Option {
	return optionFunc(func(o *options) { o.query["app_key"] = AppKey })
}

// WithUsername username获取
func (obj *_Oauth2RefreshTokenMgr) WithUsername(Username string) Option {
	return optionFunc(func(o *options) { o.query["username"] = Username })
}

// WithExpires expires获取
func (obj *_Oauth2RefreshTokenMgr) WithExpires(Expires time.Time) Option {
	return optionFunc(func(o *options) { o.query["expires"] = Expires })
}

// WithTokenExpireTime token_expire_time获取
func (obj *_Oauth2RefreshTokenMgr) WithTokenExpireTime(TokenExpireTime int) Option {
	return optionFunc(func(o *options) { o.query["token_expire_time"] = TokenExpireTime })
}

// GetByOption 功能选项模式获取
func (obj *_Oauth2RefreshTokenMgr) GetByOption(opts ...Option) (result Oauth2RefreshToken, err error) {
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
func (obj *_Oauth2RefreshTokenMgr) GetByOptions(opts ...Option) (results []*Oauth2RefreshToken, err error) {
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
func (obj *_Oauth2RefreshTokenMgr) GetFromID(ID int) (result Oauth2RefreshToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_Oauth2RefreshTokenMgr) GetBatchFromID(IDs []int) (results []*Oauth2RefreshToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

// GetFromRefreshToken 通过refresh_token获取内容
func (obj *_Oauth2RefreshTokenMgr) GetFromRefreshToken(RefreshToken string) (result Oauth2RefreshToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("refresh_token = ?", RefreshToken).Find(&result).Error

	return
}

// GetBatchFromRefreshToken 批量唯一主键查找
func (obj *_Oauth2RefreshTokenMgr) GetBatchFromRefreshToken(RefreshTokens []string) (results []*Oauth2RefreshToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("refresh_token IN (?)", RefreshTokens).Find(&results).Error

	return
}

// GetFromTokenType 通过token_type获取内容
func (obj *_Oauth2RefreshTokenMgr) GetFromTokenType(TokenType string) (results []*Oauth2RefreshToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("token_type = ?", TokenType).Find(&results).Error

	return
}

// GetBatchFromTokenType 批量唯一主键查找
func (obj *_Oauth2RefreshTokenMgr) GetBatchFromTokenType(TokenTypes []string) (results []*Oauth2RefreshToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("token_type IN (?)", TokenTypes).Find(&results).Error

	return
}

// GetFromAppKey 通过app_key获取内容
func (obj *_Oauth2RefreshTokenMgr) GetFromAppKey(AppKey string) (results []*Oauth2RefreshToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_key = ?", AppKey).Find(&results).Error

	return
}

// GetBatchFromAppKey 批量唯一主键查找
func (obj *_Oauth2RefreshTokenMgr) GetBatchFromAppKey(AppKeys []string) (results []*Oauth2RefreshToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_key IN (?)", AppKeys).Find(&results).Error

	return
}

// GetFromUsername 通过username获取内容
func (obj *_Oauth2RefreshTokenMgr) GetFromUsername(Username string) (results []*Oauth2RefreshToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("username = ?", Username).Find(&results).Error

	return
}

// GetBatchFromUsername 批量唯一主键查找
func (obj *_Oauth2RefreshTokenMgr) GetBatchFromUsername(Usernames []string) (results []*Oauth2RefreshToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("username IN (?)", Usernames).Find(&results).Error

	return
}

// GetFromExpires 通过expires获取内容
func (obj *_Oauth2RefreshTokenMgr) GetFromExpires(Expires time.Time) (results []*Oauth2RefreshToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("expires = ?", Expires).Find(&results).Error

	return
}

// GetBatchFromExpires 批量唯一主键查找
func (obj *_Oauth2RefreshTokenMgr) GetBatchFromExpires(Expiress []time.Time) (results []*Oauth2RefreshToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("expires IN (?)", Expiress).Find(&results).Error

	return
}

// GetFromTokenExpireTime 通过token_expire_time获取内容
func (obj *_Oauth2RefreshTokenMgr) GetFromTokenExpireTime(TokenExpireTime int) (results []*Oauth2RefreshToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("token_expire_time = ?", TokenExpireTime).Find(&results).Error

	return
}

// GetBatchFromTokenExpireTime 批量唯一主键查找
func (obj *_Oauth2RefreshTokenMgr) GetBatchFromTokenExpireTime(TokenExpireTimes []int) (results []*Oauth2RefreshToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("token_expire_time IN (?)", TokenExpireTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_Oauth2RefreshTokenMgr) FetchByPrimaryKey(ID int) (result Oauth2RefreshToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// FetchByUnique primay or index 获取唯一内容
func (obj *_Oauth2RefreshTokenMgr) FetchByUnique(RefreshToken string) (result Oauth2RefreshToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("refresh_token = ?", RefreshToken).Find(&result).Error

	return
}
