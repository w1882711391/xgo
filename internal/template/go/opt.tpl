package opt

import (
	"flag"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

type redis struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type etcd struct {
	Endpoints      []string `mapstructure:"endpoints"`
	Username       string   `mapstructure:"username"`
	Password       string   `mapstructure:"password"`
	EnableResolver bool     `mapstructure:"enable_resolver"`
}

type mysql struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DataBase string `mapstructure:"database"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type config struct {
	Mysql    mysql   `mapstructure:"mysql"`
	Redis    redis   `mapstructure:"redis"`
}

var (
	configFile string
	Cfg        = new(config)
)

func init() {
	time.Local = time.FixedZone("CST", 8*3600)
	flag.StringVar(&configFile, "c", "etc/config.json", "")
}

func MustInitConfig() {
	flag.Parse()

	var (
		err error
	)

	viper.SetConfigType("json")
	viper.SetConfigFile(configFile)
	if err = viper.ReadInConfig(); err != nil {
		logrus.Panicf("read in config file err: %v", err)
	}

	if err = viper.Unmarshal(Cfg); err != nil {
		logrus.Panicf("unmarshal config file err: %v", err)
	}

	return
}
