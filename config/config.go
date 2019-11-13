/**
* @file   : config.go
* @descrip: 读取&解析配置文件中的相关配置参数
* @author : ch-yk
* @create : 2018-08-30 下午2:54
* @email  : commonheart.yk@gmail.com
**/

package config


import (
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/lexkong/log"
)

/*最好每次运行的时候有传递*/
var (
	CONFIG_FILE="init.yaml"
	CONFIG_PATH="res/conf"
)

//这里放一个结构体 entity, 方便给出一些成员方法
// 同时方便以后扩展 属性
type Config struct {
	Name string
}


func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	// 初始化日志包
	c.initLog()

	// 监控配置文件变化并热加载程序
	c.watchConfig()

	return nil
}


/*这个函数初始化配置文件的信息*/
func (c *Config) initConfig() error {

	// c.Name就拿到了配置文件路径
	// 如果指定了配置文件，则解析指定的配置文件
	// `./可执行文件 -c xxx.yaml`，没有则读取默认的
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else {
		// 如果没有指定配置文件，则解析默认的配置文件
		//viper.AddConfigPath() //文件夹/目录
		viper.SetConfigFile(CONFIG_PATH + string(os.PathSeparator) + CONFIG_FILE)  //conf.yaml
	}
	// 设置配置文件格式为YAML:
	//viper.SetConfigType("yaml")

	/***********************优先处理环境变量***********************/
	// 匹配的环境变量 (如果有的话，则环境变量优先)
	viper.AutomaticEnv()
	// 环境变量的前缀为 yk_cgi
	viper.SetEnvPrefix("YK_CGI")
	// 多级配置情况(yaml 的两个空格为一级)的替换原则
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// 开始解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

// 监控配置文件变化，并热加载程序
// 这里只是保证了，配置文件变化之后，运行时能拿到变化的值
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("配置文件变化了: %s", e.Name)
	})
}

//读取日志的配置信息
func (c *Config) initLog() {
	passLagerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}

	log.InitWithConfig(&passLagerCfg)
}