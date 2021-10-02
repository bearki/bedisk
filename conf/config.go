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

	"github.com/bearki/bedisk/tools"
	"github.com/go-ini/ini"
)

// configFilePath 默认配置文件路径
var configFilePath string

// 配置文件结构
type ConfigFile struct {
	// 程序配置
	App struct {
		DataDir string `ini:"data_dir" comment:"数据储存文件夹"`
	} `ini:"app" comment:"程序配置"`
	// HTTP服务配置
	HttpServe struct {
		Host string `ini:"host" comment:"HTTP服务监听IP"`
		Port uint   `ini:"port" comment:"HTTP服务监听端口"`
	} `ini:"http_serve" comment:"HTTP服务配置"`
	// Log 日志配置
	Log struct {
		DirPath string `ini:"path" comment:"日志文件储存文件夹"`
		SaveDay uint16 `ini:"save_day" comment:"日志最大保存天数"`
		MaxSize uint16 `ini:"max_size" comment:"单文件最大保存容量（单位：MB）"`
	} `ini:"log" comment:"日志配置"`
	// MySQL数据库配置
	MySQL struct {
		Host    string `ini:"host" comment:"数据库地址"`
		Port    uint   `ini:"port" comment:"数据库端口"`
		Name    string `ini:"name" comment:"数据库名称"`
		User    string `ini:"username" comment:"数据库用户名"`
		Passwd  string `ini:"password" comment:"数据库密码"`
		CharSet string `ini:"charset" comment:"数据库字符集"`
	} `ini:"mysql" comment:"MySQL数据库配置"`
}

// 配置文件全局对象
var Config ConfigFile

// initConfig 初始化配置模块
// @params configfile string 配置文件路径
func initConfig(configfile string) {
	// 赋值配置文件路径
	configFilePath = configfile
	// 赋予配置信息默认值
	Config.setConfigDefaultVal()
	// 读取配置文件
	Config.readConfigFile()
}

// setConfigDefaultVal 赋予配置信息默认值
func (c *ConfigFile) setConfigDefaultVal() {
	c.HttpServe.Host = "0.0.0.0"                         // HTTP服务默认监听全部地址
	c.HttpServe.Port = 18018                             // HTTP服务默认监听18018端口
	c.Log.DirPath = tools.JoinPath(App.WorkPath, "logs") // 默认日志文件夹
	c.Log.MaxSize = 128                                  // 默认单日志文件最大储存128MB
	c.Log.SaveDay = 30                                   // 默认日志保存30天
	c.MySQL.Host = "127.0.0.1"                           // 默认为本机数据库
	c.MySQL.Port = 3306                                  // 默认端口3306
	c.MySQL.Name = "bedisk"                              // 默认数据库名称
	c.MySQL.User = "root"                                // 默认数据库用户
	c.MySQL.Passwd = "root"                              // 默认数据库密码
	c.MySQL.CharSet = "utf8mb4"                          // 默认数据库字符集
}

// openConfigFile 打开配置文件
// @return *ini.File 配置信息
// @return error     错误信息
func (c *ConfigFile) openConfigFile() (*ini.File, error) {
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
func (c *ConfigFile) readConfigFile() {
	// 打开配置文件
	cfg, err := c.openConfigFile()
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
func (c *ConfigFile) SaveConfigFile() error {
	// 初始化一个空配置
	cfg := ini.Empty()
	// 将结构体反射到配置文件上
	err := cfg.ReflectFrom(&Config)
	if err != nil {
		return fmt.Errorf("struct reflect to ini.file error: %s", err.Error())
	}
	// 执行保存，并返回错误信息
	return cfg.SaveTo(configFilePath)
}
