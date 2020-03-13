package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// UserAccountTbl 用户账号ssssss
type UserAccountTbl struct {
	ID            int    `gorm:"primary_key"`
	Account       string `gorm:"unique_index:account"`
	Password      string
	AccountType   int         // 帐号类型:0手机号，1邮件
	AppKey        string      // authbucket_oauth2_client表的id
	UserInfoTblID int         `gorm:"index:user_info_id"`
	UserInfoTbl   UserInfoTbl `gorm:"association_foreignkey:user_info_tbl_id;foreignkey:id"` // 用户信息
	RegTime       time.Time
	RegIP         string `gorm:"unique_index:account;index:user_info_id"`
	BundleID      string
	Describ       string
}

// UserInfoTbl 用户信息
type UserInfoTbl struct {
	gorm.Model
	Nickname string
	Headurl  string
}

// UserPaybillOrder [...]
type UserPaybillOrder struct {
	ID               int   `gorm:"primary_key"`
	PaybillID        int64 `gorm:"index:order_id"` // 二次账单id
	OrderIDMysql     int64 `gorm:"index:order_id"` // MySql中的订单Id
	OrderItemIDMysql int64 `gorm:"index:order_id"` // MySql中的订单ItemId
	OrderIDMssql     int64 // MsSql中的订单Id
}

// Oauth2AccessToken token认证
type Oauth2AccessToken struct {
	ID          int    `gorm:"primary_key"`
	AccessToken string `gorm:"unique"`
	TokenType   string
	AppKey      string    // key
	Username    string    // 用户名
	Expires     time.Time // 过期时间
}

// Oauth2RefreshToken 刷新token
type Oauth2RefreshToken struct {
	ID              int    `gorm:"primary_key"`
	RefreshToken    string `gorm:"unique"`
	TokenType       string
	AppKey          string
	Username        string
	Expires         time.Time
	TokenExpireTime int
}

// SignClientTbl [...]
type SignClientTbl struct {
	ID              int `gorm:"primary_key"`
	AppKey          string
	AppSecret       string
	ExpireTime      time.Time // 超时时间
	StrictSign      int       // 是否强制验签:0：用户自定义，1：强制
	StrictVerify    int       // 是否强制验证码:0：用户自定义，1：强制
	TokenExpireTime int       // token过期时间
}

// Oauth2ClientTbl client key 信息
type Oauth2ClientTbl struct {
	ID              int `gorm:"primary_key"`
	AppKey          string
	AppSecret       string
	ExpireTime      time.Time // 超时时间
	StrictSign      int       // 是否强制验签:0：用户自定义，1：强制
	StrictVerify    int       // 是否强制验证码:0：用户自定义，1：强制
	TokenExpireTime int       // token过期时间
	Aaa             string
}

// Organ [...]
type Organ struct {
	ID       int    `gorm:"primary_key"`
	UserID   int    `gorm:"index"`
	UserList []User `gorm:"association_foreignkey:userId;foreignkey:sex"`
	Type     int
	Score    int
}

// User [...]
type User struct {
	UserID int `gorm:"primary_key"`
	Name   string
	Sex    int `gorm:"index"`
	Job    int
}
