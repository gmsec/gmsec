package config

import (
	"fmt"
	"os"
	"path"
	"sync"

	"github.com/xxjwxc/consult"
	"github.com/xxjwxc/consult/consulkv"
	"github.com/xxjwxc/public/dev"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/tools"
	"gopkg.in/yaml.v3"
)

// CfgBase base config struct
type CfgBase struct {
	SerialNumber       string `json:"serial_number" yaml:"serial_number" consul:"serial_number" `                  // version.版本号
	ServiceName        string `json:"service_name" yaml:"service_name" consul:"service_name"`                      // service name .service名字
	ServiceDisplayname string `json:"service_displayname" yaml:"service_displayname" consul:"service_displayname"` // display name .显示名
	SerciceDesc        string `json:"sercice_desc" yaml:"sercice_desc" consul:"sercice_desc"`                      // sercice desc .service描述
	IsDev              bool   `json:"is_dev" yaml:"is_dev" consul:"is_dev"`                                        // Is it a development version?是否是开发版本
}

var (
	// _map 会被 consult.AutoLoadConfig/AutoSetConfig 通过反射直接写入(热更新),
	// 同时被各个 GetXxx 并发读取,必须加锁,否则 -race 下必然 data race。
	mu         sync.RWMutex
	_map       = Config{}
	configPath string
	// initErr 保存包 init 阶段的配置错误:init 里不能 os.Exit,由 main 决定退出码。
	initErr error
)

func init() {
	configPath = path.Join(tools.GetCurrentDirectory(), "conf/config.yml")
	initErr = onInit()
}

// ConfigError 返回配置初始化阶段的错误,由 main 决定退出码。
func ConfigError() error {
	return initErr
}

func onInit() error {
	if err := InitFile(configPath); err != nil {
		return err
	}

	mu.RLock()
	dev.OnSetDev(_map.IsDev)
	mu.RUnlock()

	// consul 热更新是可选能力:连不上只记日志,不阻断服务启动。
	if err := initConsul(); err != nil {
		mylog.Error(err)
	}
	return nil
}

// InitFile default value from file .
func InitFile(filename string) error {
	if _, err := os.Stat(filename); err != nil {
		mylog.Info("未找到配置文件,生成默认配置: ", filename)
		if e := SaveToFile(); e != nil {
			mylog.Error(e)
			return fmt.Errorf("生成默认配置文件失败: %w", e)
		}
		// 默认配置需要人工填写,返回 error 交给 main 以非 0 退出;
		// 历史实现在这里 fmt.Println + os.Exit(0) 会以"成功"掩盖启动失败。
		return fmt.Errorf("已生成默认配置文件 %s,请填写后重启", filename)
	}

	bs, err := os.ReadFile(filename)
	if err != nil {
		mylog.Error(err)
		return err
	}

	mu.Lock()
	defer mu.Unlock()
	if err := yaml.Unmarshal(bs, &_map); err != nil {
		mylog.Error(err)
		return err
	}
	// 敏感项允许用环境变量覆盖(存在才覆盖),生产环境由编排平台注入。
	applyEnvOverridesLocked()
	return nil
}

// applyEnvOverridesLocked 调用方必须已持有 mu 写锁。
func applyEnvOverridesLocked() {
	if v, ok := os.LookupEnv("GMSEC_DB_PASSWORD"); ok {
		_map.MySQLInfo.Password = v
	}
	if v, ok := os.LookupEnv("GMSEC_REDIS_PASSWORD"); ok {
		_map.RedisDbInfo.Password = v
	}
	if v, ok := os.LookupEnv("GMSEC_CONSUL_ADDR"); ok {
		_map.ConsulAddr = v
	}
	if v, ok := os.LookupEnv("GMSEC_JAEGER_ADDR"); ok {
		_map.Jaeger.Addr = v
	}
}

// GetServiceConfig Get service configuration information
func GetServiceConfig() (name, displayName, desc string) {
	mu.RLock()
	defer mu.RUnlock()
	name = _map.ServiceName
	displayName = _map.ServiceDisplayname
	desc = _map.SerciceDesc
	return
}

// GetIsDev is is dev
func GetIsDev() bool {
	mu.RLock()
	defer mu.RUnlock()
	return _map.IsDev
}

// SetIsDev is is dev
func SetIsDev(b bool) {
	mu.Lock()
	defer mu.Unlock()
	_map.IsDev = b
}

// SaveToFile save config info to file
func SaveToFile() error {
	mu.RLock()
	d, err := yaml.Marshal(_map)
	mu.RUnlock()
	if err != nil {
		return err
	}
	// tools.WriteFile 返回 bool,必须检查,否则写失败也会被当成成功。
	if !tools.WriteFile(configPath, []string{string(d)}, true) {
		return fmt.Errorf("写入配置文件失败: %s", configPath)
	}
	return nil
}

func initConsul() error {
	mu.RLock()
	addr := _map.ConsulAddr
	tag := _map.ConsulTag
	serviceName := _map.ServiceName
	mu.RUnlock()

	if len(addr) == 0 {
		return nil
	}
	if len(tag) == 0 {
		tag = "service"
	}

	// 初始化数据
	conf := consulkv.NewConfig(
		consulkv.WithPrefix(tag+"/"+serviceName), // consul kv prefix
		consulkv.WithAddress(addr),               // consul address
	)
	if err := conf.Init(); err != nil {
		return err
	}

	// consult 通过反射直接写 &_map,持写锁才能和 GetXxx 的读锁互斥。
	mu.Lock()
	defer mu.Unlock()
	if err := consult.AutoLoadConfig(conf, &_map); err != nil {
		return err
	}
	return consult.AutoSetConfig(conf, &_map, false) // 执行一次更新
}
