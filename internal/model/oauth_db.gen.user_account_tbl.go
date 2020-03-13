package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type _UserAccountTblMgr struct {
	*_BaseMgr
}

// UserAccountTblMgr open func
func UserAccountTblMgr(db *gorm.DB) *_UserAccountTblMgr {
	if db == nil {
		panic(fmt.Errorf("UserAccountTblMgr need init by db"))
	}
	return &_UserAccountTblMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: gloabIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserAccountTblMgr) GetTableName() string {
	return "user_account_tbl"
}

// Get 获取
func (obj *_UserAccountTblMgr) Get() (result UserAccountTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info UserInfoTbl // 用户信息
			err = obj.DB.New().Table("user_info_tbl").Where("id = ?", result.UserInfoTblID).Find(&info).Error
			if err != nil {
				return
			}
			result.UserInfoTbl = info
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_UserAccountTblMgr) Gets() (results []*UserAccountTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info UserInfoTbl // 用户信息
				err = obj.DB.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoTblID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].UserInfoTbl = info
			}
		}
	}
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UserAccountTblMgr) WithID(ID int) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithAccount account获取
func (obj *_UserAccountTblMgr) WithAccount(Account string) Option {
	return optionFunc(func(o *options) { o.query["account"] = Account })
}

// WithPassword password获取
func (obj *_UserAccountTblMgr) WithPassword(Password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = Password })
}

// WithAccountType account_type获取 帐号类型:0手机号，1邮件
func (obj *_UserAccountTblMgr) WithAccountType(AccountType int) Option {
	return optionFunc(func(o *options) { o.query["account_type"] = AccountType })
}

// WithAppKey app_key获取 authbucket_oauth2_client表的id
func (obj *_UserAccountTblMgr) WithAppKey(AppKey string) Option {
	return optionFunc(func(o *options) { o.query["app_key"] = AppKey })
}

// WithUserInfoTblID user_info_tbl_id获取
func (obj *_UserAccountTblMgr) WithUserInfoTblID(UserInfoTblID int) Option {
	return optionFunc(func(o *options) { o.query["user_info_tbl_id"] = UserInfoTblID })
}

// WithRegTime reg_time获取
func (obj *_UserAccountTblMgr) WithRegTime(RegTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["reg_time"] = RegTime })
}

// WithRegIP reg_ip获取
func (obj *_UserAccountTblMgr) WithRegIP(RegIP string) Option {
	return optionFunc(func(o *options) { o.query["reg_ip"] = RegIP })
}

// WithBundleID bundle_id获取
func (obj *_UserAccountTblMgr) WithBundleID(BundleID string) Option {
	return optionFunc(func(o *options) { o.query["bundle_id"] = BundleID })
}

// WithDescrib describ获取
func (obj *_UserAccountTblMgr) WithDescrib(Describ string) Option {
	return optionFunc(func(o *options) { o.query["describ"] = Describ })
}

// GetByOption 功能选项模式获取
func (obj *_UserAccountTblMgr) GetByOption(opts ...Option) (result UserAccountTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info UserInfoTbl // 用户信息
			err = obj.DB.New().Table("user_info_tbl").Where("id = ?", result.UserInfoTblID).Find(&info).Error
			if err != nil {
				return
			}
			result.UserInfoTbl = info
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UserAccountTblMgr) GetByOptions(opts ...Option) (results []*UserAccountTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info UserInfoTbl // 用户信息
				err = obj.DB.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoTblID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].UserInfoTbl = info
			}
		}
	}
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_UserAccountTblMgr) GetFromID(ID int) (result UserAccountTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info UserInfoTbl // 用户信息
			err = obj.DB.New().Table("user_info_tbl").Where("id = ?", result.UserInfoTblID).Find(&info).Error
			if err != nil {
				return
			}
			result.UserInfoTbl = info
		}
	}

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_UserAccountTblMgr) GetBatchFromID(IDs []int) (results []*UserAccountTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info UserInfoTbl // 用户信息
				err = obj.DB.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoTblID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].UserInfoTbl = info
			}
		}
	}
	return
}

// GetFromAccount 通过account获取内容
func (obj *_UserAccountTblMgr) GetFromAccount(Account string) (result UserAccountTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("account = ?", Account).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info UserInfoTbl // 用户信息
			err = obj.DB.New().Table("user_info_tbl").Where("id = ?", result.UserInfoTblID).Find(&info).Error
			if err != nil {
				return
			}
			result.UserInfoTbl = info
		}
	}

	return
}

// GetBatchFromAccount 批量唯一主键查找
func (obj *_UserAccountTblMgr) GetBatchFromAccount(Accounts []string) (results []*UserAccountTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("account IN (?)", Accounts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info UserInfoTbl // 用户信息
				err = obj.DB.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoTblID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].UserInfoTbl = info
			}
		}
	}
	return
}

// GetFromPassword 通过password获取内容
func (obj *_UserAccountTblMgr) GetFromPassword(Password string) (results []*UserAccountTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("password = ?", Password).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info UserInfoTbl // 用户信息
				err = obj.DB.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoTblID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].UserInfoTbl = info
			}
		}
	}
	return
}

// GetBatchFromPassword 批量唯一主键查找
func (obj *_UserAccountTblMgr) GetBatchFromPassword(Passwords []string) (results []*UserAccountTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("password IN (?)", Passwords).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info UserInfoTbl // 用户信息
				err = obj.DB.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoTblID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].UserInfoTbl = info
			}
		}
	}
	return
}

// GetFromAccountType 通过account_type获取内容 帐号类型:0手机号，1邮件
func (obj *_UserAccountTblMgr) GetFromAccountType(AccountType int) (results []*UserAccountTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("account_type = ?", AccountType).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info UserInfoTbl // 用户信息
				err = obj.DB.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoTblID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].UserInfoTbl = info
			}
		}
	}
	return
}

// GetBatchFromAccountType 批量唯一主键查找 帐号类型:0手机号，1邮件
func (obj *_UserAccountTblMgr) GetBatchFromAccountType(AccountTypes []int) (results []*UserAccountTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("account_type IN (?)", AccountTypes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info UserInfoTbl // 用户信息
				err = obj.DB.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoTblID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].UserInfoTbl = info
			}
		}
	}
	return
}

// GetFromAppKey 通过app_key获取内容 authbucket_oauth2_client表的id
func (obj *_UserAccountTblMgr) GetFromAppKey(AppKey string) (results []*UserAccountTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_key = ?", AppKey).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info UserInfoTbl // 用户信息
				err = obj.DB.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoTblID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].UserInfoTbl = info
			}
		}
	}
	return
}

// GetBatchFromAppKey 批量唯一主键查找 authbucket_oauth2_client表的id
func (obj *_UserAccountTblMgr) GetBatchFromAppKey(AppKeys []string) (results []*UserAccountTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_key IN (?)", AppKeys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info UserInfoTbl // 用户信息
				err = obj.DB.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoTblID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].UserInfoTbl = info
			}
		}
	}
	return
}

// GetFromUserInfoTblID 通过user_info_tbl_id获取内容
func (obj *_UserAccountTblMgr) GetFromUserInfoTblID(UserInfoTblID int) (results []*UserAccountTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_info_tbl_id = ?", UserInfoTblID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info UserInfoTbl // 用户信息
				err = obj.DB.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoTblID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].UserInfoTbl = info
			}
		}
	}
	return
}

// GetBatchFromUserInfoTblID 批量唯一主键查找
func (obj *_UserAccountTblMgr) GetBatchFromUserInfoTblID(UserInfoTblIDs []int) (results []*UserAccountTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_info_tbl_id IN (?)", UserInfoTblIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info UserInfoTbl // 用户信息
				err = obj.DB.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoTblID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].UserInfoTbl = info
			}
		}
	}
	return
}

// GetFromRegTime 通过reg_time获取内容
func (obj *_UserAccountTblMgr) GetFromRegTime(RegTime time.Time) (results []*UserAccountTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("reg_time = ?", RegTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info UserInfoTbl // 用户信息
				err = obj.DB.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoTblID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].UserInfoTbl = info
			}
		}
	}
	return
}

// GetBatchFromRegTime 批量唯一主键查找
func (obj *_UserAccountTblMgr) GetBatchFromRegTime(RegTimes []time.Time) (results []*UserAccountTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("reg_time IN (?)", RegTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info UserInfoTbl // 用户信息
				err = obj.DB.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoTblID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].UserInfoTbl = info
			}
		}
	}
	return
}

// GetFromRegIP 通过reg_ip获取内容
func (obj *_UserAccountTblMgr) GetFromRegIP(RegIP string) (result UserAccountTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("reg_ip = ?", RegIP).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info UserInfoTbl // 用户信息
			err = obj.DB.New().Table("user_info_tbl").Where("id = ?", result.UserInfoTblID).Find(&info).Error
			if err != nil {
				return
			}
			result.UserInfoTbl = info
		}
	}

	return
}

// GetBatchFromRegIP 批量唯一主键查找
func (obj *_UserAccountTblMgr) GetBatchFromRegIP(RegIPs []string) (results []*UserAccountTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("reg_ip IN (?)", RegIPs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info UserInfoTbl // 用户信息
				err = obj.DB.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoTblID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].UserInfoTbl = info
			}
		}
	}
	return
}

// GetFromBundleID 通过bundle_id获取内容
func (obj *_UserAccountTblMgr) GetFromBundleID(BundleID string) (results []*UserAccountTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("bundle_id = ?", BundleID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info UserInfoTbl // 用户信息
				err = obj.DB.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoTblID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].UserInfoTbl = info
			}
		}
	}
	return
}

// GetBatchFromBundleID 批量唯一主键查找
func (obj *_UserAccountTblMgr) GetBatchFromBundleID(BundleIDs []string) (results []*UserAccountTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("bundle_id IN (?)", BundleIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info UserInfoTbl // 用户信息
				err = obj.DB.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoTblID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].UserInfoTbl = info
			}
		}
	}
	return
}

// GetFromDescrib 通过describ获取内容
func (obj *_UserAccountTblMgr) GetFromDescrib(Describ string) (results []*UserAccountTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("describ = ?", Describ).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info UserInfoTbl // 用户信息
				err = obj.DB.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoTblID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].UserInfoTbl = info
			}
		}
	}
	return
}

// GetBatchFromDescrib 批量唯一主键查找
func (obj *_UserAccountTblMgr) GetBatchFromDescrib(Describs []string) (results []*UserAccountTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("describ IN (?)", Describs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info UserInfoTbl // 用户信息
				err = obj.DB.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoTblID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].UserInfoTbl = info
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_UserAccountTblMgr) FetchByPrimaryKey(ID int) (result UserAccountTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info UserInfoTbl // 用户信息
			err = obj.DB.New().Table("user_info_tbl").Where("id = ?", result.UserInfoTblID).Find(&info).Error
			if err != nil {
				return
			}
			result.UserInfoTbl = info
		}
	}

	return
}

// FetchByAccountUniqueIndex primay or index 获取唯一内容
func (obj *_UserAccountTblMgr) FetchByAccountUniqueIndex(Account string, RegIP string) (result UserAccountTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("account = ? AND reg_ip = ?", Account, RegIP).Find(&result).Error
	if err == nil && obj.isRelated {
		{
			var info UserInfoTbl // 用户信息
			err = obj.DB.New().Table("user_info_tbl").Where("id = ?", result.UserInfoTblID).Find(&info).Error
			if err != nil {
				return
			}
			result.UserInfoTbl = info
		}
	}

	return
}

// FetchByUserInfoIDIndex  获取多个内容
func (obj *_UserAccountTblMgr) FetchByUserInfoIDIndex(UserInfoTblID int, RegIP string) (results []*UserAccountTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_info_tbl_id = ? AND reg_ip = ?", UserInfoTblID, RegIP).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			{
				var info UserInfoTbl // 用户信息
				err = obj.DB.New().Table("user_info_tbl").Where("id = ?", results[i].UserInfoTblID).Find(&info).Error
				if err != nil {
					return
				}
				results[i].UserInfoTbl = info
			}
		}
	}
	return
}
