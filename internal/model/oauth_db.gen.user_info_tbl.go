package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type _UserInfoTblMgr struct {
	*_BaseMgr
}

// UserInfoTblMgr open func
func UserInfoTblMgr(db *gorm.DB) *_UserInfoTblMgr {
	if db == nil {
		panic(fmt.Errorf("UserInfoTblMgr need init by db"))
	}
	return &_UserInfoTblMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: gloabIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserInfoTblMgr) GetTableName() string {
	return "user_info_tbl"
}

// Get 获取
func (obj *_UserInfoTblMgr) Get() (result UserInfoTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserInfoTblMgr) Gets() (results []*UserInfoTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 Primary key
func (obj *_UserInfoTblMgr) WithID(ID int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithCreatedAt created_at获取 created time
func (obj *_UserInfoTblMgr) WithCreatedAt(CreatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = CreatedAt })
}

// WithUpdatedAt updated_at获取 updated at
func (obj *_UserInfoTblMgr) WithUpdatedAt(UpdatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = UpdatedAt })
}

// WithDeletedAt deleted_at获取 deleted time
func (obj *_UserInfoTblMgr) WithDeletedAt(DeletedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = DeletedAt })
}

// WithNickname nickname获取
func (obj *_UserInfoTblMgr) WithNickname(Nickname string) Option {
	return optionFunc(func(o *options) { o.query["nickname"] = Nickname })
}

// WithHeadurl headurl获取
func (obj *_UserInfoTblMgr) WithHeadurl(Headurl string) Option {
	return optionFunc(func(o *options) { o.query["headurl"] = Headurl })
}

// GetByOption 功能选项模式获取
func (obj *_UserInfoTblMgr) GetByOption(opts ...Option) (result UserInfoTbl, err error) {
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
func (obj *_UserInfoTblMgr) GetByOptions(opts ...Option) (results []*UserInfoTbl, err error) {
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

// GetFromID 通过id获取内容 Primary key
func (obj *_UserInfoTblMgr) GetFromID(ID int64) (result UserInfoTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 Primary key
func (obj *_UserInfoTblMgr) GetBatchFromID(IDs []int64) (results []*UserInfoTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 created time
func (obj *_UserInfoTblMgr) GetFromCreatedAt(CreatedAt time.Time) (result UserInfoTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("created_at = ?", CreatedAt).Find(&result).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 created time
func (obj *_UserInfoTblMgr) GetBatchFromCreatedAt(CreatedAts []time.Time) (results []*UserInfoTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("created_at IN (?)", CreatedAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 updated at
func (obj *_UserInfoTblMgr) GetFromUpdatedAt(UpdatedAt time.Time) (result UserInfoTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("updated_at = ?", UpdatedAt).Find(&result).Error

	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找 updated at
func (obj *_UserInfoTblMgr) GetBatchFromUpdatedAt(UpdatedAts []time.Time) (results []*UserInfoTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("updated_at IN (?)", UpdatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 deleted time
func (obj *_UserInfoTblMgr) GetFromDeletedAt(DeletedAt time.Time) (result UserInfoTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("deleted_at = ?", DeletedAt).Find(&result).Error

	return
}

// GetBatchFromDeletedAt 批量唯一主键查找 deleted time
func (obj *_UserInfoTblMgr) GetBatchFromDeletedAt(DeletedAts []time.Time) (results []*UserInfoTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("deleted_at IN (?)", DeletedAts).Find(&results).Error

	return
}

// GetFromNickname 通过nickname获取内容
func (obj *_UserInfoTblMgr) GetFromNickname(Nickname string) (results []*UserInfoTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("nickname = ?", Nickname).Find(&results).Error

	return
}

// GetBatchFromNickname 批量唯一主键查找
func (obj *_UserInfoTblMgr) GetBatchFromNickname(Nicknames []string) (results []*UserInfoTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("nickname IN (?)", Nicknames).Find(&results).Error

	return
}

// GetFromHeadurl 通过headurl获取内容
func (obj *_UserInfoTblMgr) GetFromHeadurl(Headurl string) (results []*UserInfoTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("headurl = ?", Headurl).Find(&results).Error

	return
}

// GetBatchFromHeadurl 批量唯一主键查找
func (obj *_UserInfoTblMgr) GetBatchFromHeadurl(Headurls []string) (results []*UserInfoTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("headurl IN (?)", Headurls).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_UserInfoTblMgr) FetchByPrimaryKey(ID int64) (result UserInfoTbl, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}
