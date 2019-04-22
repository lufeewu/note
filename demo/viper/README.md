# Viper
> 用于 golang 应用程序的配置文件管理
	
* 支持默认值
* 可以读取 JSON、TOML、YAML、HCL 和 java 属性的配置文件
* 可以从环境变量读取值
* 可以从远程配置系统读取（etcd 、consul）, 并可以持续监控更改
* 可以从命令行读取 flags
* 可以从 buffer 读取值
* 可以设置明确的值

## 函数
* WatchConfig()
* OnConfigChange()
* Get()
* Set()
* SetDefault()
* SetConfigName()
* AddConfigPath()
* ReadInConfig()
* WatchConfig()
* OnConfigChange()
* SetConfigType()
* RegisterAlias()
* GetBool()
* AutomaticEnv()
* BindEnv()
* SetEnvPrefix()
* SetEnvKeyReplace()
* AllowEmptyEnvVar()
* AddRemoteProvider()
* SetConfigType()
* ReadRemoteConfig()

## package
    - github.com/fsnotify/fsnotify 提供平台无关的文件系统通知
        - NewWatcher()
        - Event
        - Write、Create
    - path/filepath
        - Clean()
        - Split()
        - EvalSymlinks()