/**
 *@Title 初始化配置模块
 *@Desc 程序初始化将从这里开始，configs包的初始化操作将在main方法之前执行
 *@Author Bearki
 *@DateTime 2021/09/20 17:55
 */

package configs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-ini/ini"
)

// configFilePath 默认配置文件路径
var configFilePath = "./config.ini"

// 配置文件结构
type ConfigFile struct {
	HttpServe struct {
		Host string `ini:"http_serve_host" comment:"HTTP服务监听IP"`
		Port uint   `ini:"http_serve_port" comment:"HTTP服务监听端口"`
	} `ini:"http_serve" comment:"HTTP服务配置"`
	Log struct {
		Path    string `ini:"log_path" comment:"日志文件储存路径（不含文件名）"`
		SaveDay uint   `ini:"log_save_day" comment:"日志最大保存天数"`
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
	Config.HttpServe.Host = "0.0.0.0" // HTTP服务默认监听全部地址
	Config.HttpServe.Port = 18018     // HTTP服务默认监听18018端口
	Config.Log.Path = "./logs/bedisk.log"
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
				fmt.Printf("create `%s` dir error: %s\n", filepath.Dir(configFilePath), err.Error())
				return nil, err
			}
			// 创建文件
			_, err = os.Create(configFilePath)
			if err != nil {
				fmt.Printf("create `%s` file error: %s\n", configFilePath, err.Error())
				return nil, err
			}
		} else {
			fmt.Printf("get `%s` file info error: %s\n", configFilePath, err.Error())
			return nil, err
		}
	}
	// 加载配置文件到私有全局配置对象，程序不死，它不毁灭
	cfg, err := ini.Load(configFilePath)
	if err != nil {
		fmt.Printf("open `%s` file error: %s\n", configFilePath, err.Error())
		return nil, err
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
		fmt.Printf("config file map to struct error: %s\n", err.Error())
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
