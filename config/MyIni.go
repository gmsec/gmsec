package config

import (
	"fmt"
)

// Config custom config struct
type Config struct {
	CfgBase   `yaml:"base"`
	MySQLInfo MysqlDbInfo `yaml:"mysql_info"`
}

// MysqlDbInfo mysql database information. mysql 数据库信息
type MysqlDbInfo struct {
	Host     string `validate:"required"` // Host. 地址
	Port     int    `validate:"required"` // Port 端口号
	Username string `validate:"required"` // Username 用户名
	Password string // Password 密码
	Database string `validate:"required"` // Database 数据库名
}

// SetMysqlDbInfo Update MySQL configuration information
func SetMysqlDbInfo(info *MysqlDbInfo) {
	_map.MySQLInfo = *info
}

// GetMysqlDbInfo Get MySQL configuration information .获取mysql配置信息
func GetMysqlDbInfo() MysqlDbInfo {
	return _map.MySQLInfo
}

// GetMysqlConStr Get MySQL connection string.获取mysql 连接字符串
func GetMysqlConStr() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&interpolateParams=True",
		_map.MySQLInfo.Username,
		_map.MySQLInfo.Password,
		_map.MySQLInfo.Host,
		_map.MySQLInfo.Port,
		_map.MySQLInfo.Database,
	)
}
