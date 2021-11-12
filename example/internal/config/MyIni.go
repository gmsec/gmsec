package config

import (
	"fmt"
)

// Config custom config struct
type Config struct {
	CfgBase     `yaml:"base" consul:"base"`
	MySQLInfo   MysqlDbInfo `yaml:"mysql_info" consul:"mysql_info"`
	EtcdInfo    EtcdInfo    `yaml:"etcd_info" consul:"etcd_info"`
	RedisDbInfo RedisDbInfo `yaml:"redis_info" consul:"redis_info"`
	Oauth2Url   string      `yaml:"oauth2_url" consul:"oauth2_url"`
	Port        string      `yaml:"port" consul:"port"`                // 端口号
	ConsulAddr  string      `yaml:"consul_addr" consul:"consul_addr" ` // consul 地址
	ConsulTag   string      `yaml:"consul_tag" consul:"consul_tag"`
}

// MysqlDbInfo mysql database information. mysql 数据库信息
type MysqlDbInfo struct {
	Host     string `validate:"required" consul:"host"`     // Host. 地址
	Port     int    `validate:"required" consul:"port"`     // Port 端口号
	Username string `validate:"required" consul:"username"` // Username 用户名
	Password string `yaml:"password" consul:"password"`     // Password 密码
	Database string `validate:"required" consul:"database"` // Database 数据库名
	Type     int    `consul:"type"`                         // 数据库类型: 0:mysql , 1:sqlite , 2:mssql
}

// RedisDbInfo redis database information. redis 数据库信息
type RedisDbInfo struct {
	Addrs     []string `yaml:"addrs" consul:"addrs"`           // Host. 地址
	Password  string   `yaml:"password" consul:"password"`     // Password 密码
	GroupName string   `yaml:"group_name" consul:"group_name"` // 分组名字
	DB        int      `yaml:"db" consul:"db"`                 // 数据库序号
}

// EtcdInfo etcd config info
type EtcdInfo struct {
	Addrs   []string `yaml:"addrs" consul:"addrs"`     // Host. 地址
	Timeout int      `yaml:"timeout" consul:"timeout"` // 超时时间(秒)
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

// GetCheckTokenURL checktoken
func GetCheckTokenURL() string {
	return _map.Oauth2Url + "/check_token"
}

// GetLoginURL 登陆用的url
func GetLoginURL() string {
	return _map.Oauth2Url + "/authorize"
}

// GetLoginNoPwdURL token 授权模式登陆
func GetLoginNoPwdURL() string {
	return _map.Oauth2Url + "/authorize_nopwd"
}

// GetPort 获取端口号
func GetPort() string {
	return _map.Port
}

// GetRedisDbInfo Get redis configuration information .获取redis配置信息
func GetRedisDbInfo() RedisDbInfo {
	return _map.RedisDbInfo
}

// GetEtcdInfo get etcd configuration information. 获取etcd 配置信息
func GetEtcdInfo() EtcdInfo {
	return _map.EtcdInfo
}

func GetConsulTag() string {
	if len(_map.ConsulTag) > 0 {
		return _map.ConsulTag
	}

	return "service"
}
