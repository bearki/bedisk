/**
 *@Title 初始化配置模块
 *@Desc 程序初始化将从这里开始，configs包的初始化操作将在main方法之前执行
 *@Author Bearki
 *@DateTime 2021/09/20 17:55
 */

package conf

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Bearki/BeDisk/tools"
	"github.com/go-ini/ini"
)

// configFilePath 默认配置文件路径
var configFilePath string

// 配置文件结构
type ConfigFile struct {
	HttpServe struct {
		Host string `ini:"http_serve_host" comment:"HTTP服务监听IP"`
		Port uint   `ini:"http_serve_port" comment:"HTTP服务监听端口"`
	} `ini:"http_serve" comment:"HTTP服务配置"`
	Log struct {
		DirPath string `ini:"log_path" comment:"日志文件储存文件夹"`
		SaveDay uint16 `ini:"log_save_day" comment:"日志最大保存天数"`
		MaxSize uint16 `ini:"max_size" comment:"单文件最大保存容量（单位：MB）"`
	}
}

// 配置文件全局对象
var Config ConfigFile

// initConfig 初始化配置模块
func initConfig() {
	// 赋予配置信息默认值
	setConfigDefaultVal()
	// 读取配置文件
	readConfigFile()
}

// setConfigDefaultVal 赋予配置信息默认值
func setConfigDefaultVal() {
	configFilePath = tools.JoinPath(App.WorkPath, "bedisk.conf") // 默认的配置文件路径
	Config.HttpServe.Host = "0.0.0.0"                            // HTTP服务默认监听全部地址
	Config.HttpServe.Port = 18018                                // HTTP服务默认监听18018端口
	Config.Log.DirPath = tools.JoinPath(App.WorkPath, "logs")    // 默认日志文件夹
	Config.Log.MaxSize = 128                                     // 默认单日志文件最大储存128MB
	Config.Log.SaveDay = 30                                      // 默认日志保存30天
}

// openConfigFile 打开配置文件
// @return *ini.File 配置信息
// @return error     错误信息
func openConfigFile() (*ini.File, error) {
	// 获取配置文件头信息
	_, err := os.Stat(configFilePath)
	if err != nil {
		// 判断文件是否不存在
		if os.IsNotExist(err) {
			// 创建文件夹部分
			err = os.MkdirAll(filepath.Dir(configFilePath), 0755)
			if err != nil {
				e := fmt.Errorf("create `%s` dir error: %s", filepath.Dir(configFilePath), err.Error())
				return nil, e
			}
			// 创建文件
			_, err = os.Create(configFilePath)
			if err != nil {
				e := fmt.Errorf("create `%s` file error: %s", configFilePath, err.Error())
				return nil, e
			}
		} else {
			e := fmt.Errorf("get `%s` file info error: %s", configFilePath, err.Error())
			return nil, e
		}
	}
	// 加载配置文件到私有全局配置对象，程序不死，它不毁灭
	cfg, err := ini.Load(configFilePath)
	if err != nil {
		e := fmt.Errorf("open `%s` file error: %s", configFilePath, err.Error())
		return nil, e
	}
	// 返回配置信息
	return cfg, nil
}

// readConfigFile 读取配置文件内容
func readConfigFile() {
	// 打开配置文件
	cfg, err := openConfigFile()
	if err != nil {
		panic(err.Error())
	}

	// 映射配置信息到结构体
	err = cfg.MapTo(&App)
	if err != nil {
		err = fmt.Errorf("config file map to struct error: %s", err.Error())
		panic(err.Error())
	}
}

// SaveConfigFile 保存配置信息到文件(全量保存)
// @return         error       错误信息
func SaveConfigFile() error {
	// 打开配置文件
	cfg, err := openConfigFile()
	if err != nil {
		return fmt.Errorf("the `%s` profile is not open: %s", configFilePath, err.Error())
	}
	// 将结构体反射到配置文件上
	err = cfg.ReflectFrom(&Config)
	if err != nil {
		return fmt.Errorf("struct reflect to ini.file error: %s", err.Error())
	}
	// 执行保存，并返回错误信息
	return cfg.SaveTo(configFilePath)
}
