package config

import (
	rlog "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"sync"
	"time"
)

var (
	gloablConfig GlobalConfig
	once         sync.Once //用单例来使配置文件（GlobalConfig）只需初始化一次
)

type GlobalConfig struct {
	AppConfig AppConf `yaml:"app" mapstructure:"app"`
	LogConfig LogConf `yaml:"log" mapstructure:"log"`
	DbConfig  DbConf  `yaml:"db" mapstructure:"db"`
}

type AppConf struct {
	AppName string `yaml:"app_name" mapstructure:"app_name"` //yaml表示是在哪种类型文件中，mapstructure才是真正起对应关系,其实只需要后一个
	Version string `yaml:"version" mapstructure:"version"`
	Port    int    `yaml:"port" mapstructure:"port"`
	RunMod  string `yaml:"run_mod" mapstructure:"run_mod"`
}

type LogConf struct {
	LogPattern string `yaml:"log_pattern" mapstructure:"log_pattern"`
	LogPath    string `yaml:"log_path" mapstructure:"log_path"`
	SaveDays   uint   `yaml:"save_days" mapstructure:"save_days"`
	Level      string `yaml:"level" mapstructure:"level"`
}

type DbConf struct {
	Host        string `yaml:"host" mapstructure:"host"`
	Port        int    `yaml:"port" mapstructure:"port"`
	User        string `yaml:"user" mapstructure:"user"`
	PassWord    string `yaml:"password" mapstructure:"password"`
	DbName      string `yaml:"dbname" mapstructure:"dbname"`
	MaxIdleConn int    `yaml:"max_idle_conn" mapstructure:"host"mapstructure:"max_idle_conn"`
	MaxOpenConn int    `yaml:"max_open_conn" mapstructure:"host"mapstructure:"max_open_conn"`
	MaxIdleTime int    `yaml:"max_idle_time" mapstructure:"host"mapstructure:"max_idle_time"`
}

func GetGlobalConfig() *GlobalConfig { //对外暴露配置文件调用接口
	once.Do(readConf) //外部调用GetGlobalConfig()方法时，第一次GlobalConfig里啥都没有，需进行初始化，第二次调用就会跳过该命令直接返回&gloablConfig
	return &gloablConfig
}

func readConf() {
	viper.SetConfigName("config")   //当前文件名
	viper.SetConfigType("yaml")     //文件类型
	viper.AddConfigPath("./config") //从main()路径下开始找配置文件路径

	err := viper.ReadInConfig() //读取已加载的配置文件(yaml)
	if err != nil {
		panic("read config file err" + err.Error()) //解析配置出错，读取原因
	}
	err = viper.Unmarshal(&gloablConfig) //将已读取的配置文件映射到gloablConfig里
	if err != nil {
		panic("Unmarshal config file err")
	}
}

// 日志
func InitGlobalConfig() {
	config := GetGlobalConfig()
	level, err := log.ParseLevel(config.LogConfig.Level)
	if err != nil {
		panic("parse log level err")
	}
	log.SetFormatter(&logFormatter{
		log.TextFormatter{
			DisableColors:   true,
			TimestampFormat: "2006-01-02 15:04:05", //go固定起始时间，不能是其他的
			FullTimestamp:   true,                  //时间时时刻刻打印
		},
	})
	log.SetReportCaller(true)                  //调用打印文件位置
	log.SetLevel(level)                        //打印日志级别
	switch gloablConfig.LogConfig.LogPattern { //打印模式
	case "stdout":
		log.SetOutput(os.Stdout)
	case "stderr":
		log.SetOutput(os.Stderr)
	case "file":
		logger, err := rlog.New( //调用rlog库对日志进行拆分
			config.LogConfig.LogPath+".%Y%m%d", //在日志文件名后面加了时间戳
			rlog.WithRotationTime(time.Hour*24),
			rlog.WithRotationCount(config.LogConfig.SaveDays),
		)
		if err != nil {
			panic("log conf err")
		}
		log.SetOutput(logger)
	default:
		panic("log init err")
	}
}
