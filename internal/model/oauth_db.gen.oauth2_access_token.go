package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type _Oauth2AccessTokenMgr struct {
	*_BaseMgr
}

// Oauth2AccessTokenMgr open func
func Oauth2AccessTokenMgr(db *gorm.DB) *_Oauth2AccessTokenMgr {
	if db == nil {
		panic(fmt.Errorf("Oauth2AccessTokenMgr need init by db"))
	}
	return &_Oauth2AccessTokenMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: gloabIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Oauth2AccessTokenMgr) GetTableName() string {
	return "oauth2_access_token"
}

// Get 获取
func (obj *_Oauth2AccessTokenMgr) Get() (result Oauth2AccessToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_Oauth2AccessTokenMgr) Gets() (results []*Oauth2AccessToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_Oauth2AccessTokenMgr) WithID(ID int) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithAccessToken access_token获取
func (obj *_Oauth2AccessTokenMgr) WithAccessToken(AccessToken string) Option {
	return optionFunc(func(o *options) { o.query["access_token"] = AccessToken })
}

// WithTokenType token_type获取
func (obj *_Oauth2AccessTokenMgr) WithTokenType(TokenType string) Option {
	return optionFunc(func(o *options) { o.query["token_type"] = TokenType })
}

// WithAppKey app_key获取 key
func (obj *_Oauth2AccessTokenMgr) WithAppKey(AppKey string) Option {
	return optionFunc(func(o *options) { o.query["app_key"] = AppKey })
}

// WithUsername username获取 用户名
func (obj *_Oauth2AccessTokenMgr) WithUsername(Username string) Option {
	return optionFunc(func(o *options) { o.query["username"] = Username })
}

// WithExpires expires获取 过期时间
func (obj *_Oauth2AccessTokenMgr) WithExpires(Expires time.Time) Option {
	return optionFunc(func(o *options) { o.query["expires"] = Expires })
}

// GetByOption 功能选项模式获取
func (obj *_Oauth2AccessTokenMgr) GetByOption(opts ...Option) (result Oauth2AccessToken, err error) {
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
func (obj *_Oauth2AccessTokenMgr) GetByOptions(opts ...Option) (results []*Oauth2AccessToken, err error) {
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
func (obj *_Oauth2AccessTokenMgr) GetFromID(ID int) (result Oauth2AccessToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_Oauth2AccessTokenMgr) GetBatchFromID(IDs []int) (results []*Oauth2AccessToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

// GetFromAccessToken 通过access_token获取内容
func (obj *_Oauth2AccessTokenMgr) GetFromAccessToken(AccessToken string) (result Oauth2AccessToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("access_token = ?", AccessToken).Find(&result).Error

	return
}

// GetBatchFromAccessToken 批量唯一主键查找
func (obj *_Oauth2AccessTokenMgr) GetBatchFromAccessToken(AccessTokens []string) (results []*Oauth2AccessToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("access_token IN (?)", AccessTokens).Find(&results).Error

	return
}

// GetFromTokenType 通过token_type获取内容
func (obj *_Oauth2AccessTokenMgr) GetFromTokenType(TokenType string) (results []*Oauth2AccessToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("token_type = ?", TokenType).Find(&results).Error

	return
}

// GetBatchFromTokenType 批量唯一主键查找
func (obj *_Oauth2AccessTokenMgr) GetBatchFromTokenType(TokenTypes []string) (results []*Oauth2AccessToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("token_type IN (?)", TokenTypes).Find(&results).Error

	return
}

// GetFromAppKey 通过app_key获取内容 key
func (obj *_Oauth2AccessTokenMgr) GetFromAppKey(AppKey string) (results []*Oauth2AccessToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_key = ?", AppKey).Find(&results).Error

	return
}

// GetBatchFromAppKey 批量唯一主键查找 key
func (obj *_Oauth2AccessTokenMgr) GetBatchFromAppKey(AppKeys []string) (results []*Oauth2AccessToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_key IN (?)", AppKeys).Find(&results).Error

	return
}

// GetFromUsername 通过username获取内容 用户名
func (obj *_Oauth2AccessTokenMgr) GetFromUsername(Username string) (results []*Oauth2AccessToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("username = ?", Username).Find(&results).Error

	return
}

// GetBatchFromUsername 批量唯一主键查找 用户名
func (obj *_Oauth2AccessTokenMgr) GetBatchFromUsername(Usernames []string) (results []*Oauth2AccessToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("username IN (?)", Usernames).Find(&results).Error

	return
}

// GetFromExpires 通过expires获取内容 过期时间
func (obj *_Oauth2AccessTokenMgr) GetFromExpires(Expires time.Time) (results []*Oauth2AccessToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("expires = ?", Expires).Find(&results).Error

	return
}

// GetBatchFromExpires 批量唯一主键查找 过期时间
func (obj *_Oauth2AccessTokenMgr) GetBatchFromExpires(Expiress []time.Time) (results []*Oauth2AccessToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("expires IN (?)", Expiress).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_Oauth2AccessTokenMgr) FetchByPrimaryKey(ID int) (result Oauth2AccessToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// FetchByUnique primay or index 获取唯一内容
func (obj *_Oauth2AccessTokenMgr) FetchByUnique(AccessToken string) (result Oauth2AccessToken, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("access_token = ?", AccessToken).Find(&result).Error

	return
}
