/*
  viper : Go 应用程序配置文件解决方案
  特性 :
	* 支持默认值
	* 可以读取 JSON、TOML、YAML、HCL 和 java 属性的配置文件
	* 可以从环境变量读取值
	* 可以从远程配置系统读取（etcd 、consul）, 并可以持续监控更改
	* 可以从命令行读取 flags
	* 可以从 buffer 读取值
	* 可以设置明确的值
*/

package main

import (
	"flag"
	"os"
	"reflect"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func viperDefault() {
	viper.SetDefault("content", 323)
	value3 := viper.Get("content")
	logrus.Infoln("test:", value3)
}

func viperJSON() {

}

func viperEnv() {
	viper.BindEnv("id")
	os.Setenv("ID", "123")
	id := viper.Get("id")
	ID := viper.Get("ID")
	logrus.Infoln(id, ID)
}

func viperRemote() {

}

func viperFlags() {
	flag.Int("flagname", 1234, "help message for flagname")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	value := viper.GetInt("flagname")
	logrus.Infoln(value)
}

func viperBuffer() {

}

func viperExplict() {

}

func watchViper() {
	ch := make(chan struct{}, 1)

	viper.SetConfigName("config") // name of config file (without extension)
	// viper.AddConfigPath("/")  // path to look for the config file in
	// viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath(".")    // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		logrus.Errorf("Fatal error config file: %s \n", err)
	}
	value := viper.Get("test")
	logrus.Infof("%s value is %v", reflect.TypeOf(value), value)
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logrus.Infoln(e.Name, e.Op, e.String())
		value := viper.Get("test")
		logrus.Infof("%s value is %v", reflect.TypeOf(value), value)
		// logrus.Infoln("Config file changed:", e.Name)
		value2 := viper.AllSettings()
		logrus.Infoln("all settings:", value2)
	})

	<-ch
}

func main() {
	// viperEnv()
	viperFlags()
	watchViper()
}
