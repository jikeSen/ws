package library

import (
    "fmt"
    "github.com/spf13/viper"
    "gopkg.in/ini.v1"
    "os"
)

// 初始化连接器中所有的参数
var iniFile *ini.File

// MySQL 设置
type Mysql struct {
    Type     string
    User     string
    Password string
    Host     string
    Name     string
    Port     int
}

var MysqlSet = &Mysql{}

// redis 设置
type Redis struct {
    Host string
    Port string
    Auth string
}

var RedisSet = &Redis{}

func Init() {
    dir, _ := os.Getwd()

    var err error
    iniFile, err = ini.Load(dir + "/app.ini")

    if err != nil {
        panic(fmt.Errorf("加载配置文件错误: %s \n", err))
    }
    LoadMysqlConf()
    LoadRedisConf()

}

func LoadRedisConf() {
    redisSection := iniFile.Section("redis")

    RedisSet.Host = redisSection.Key("HOST").String()
    RedisSet.Port = redisSection.Key("PORT").String()
    RedisSet.Auth = redisSection.Key("PORT").String()
}

func LoadMysqlConf() {
    mysqlSection := iniFile.Section("mysql")

    port, err := mysqlSection.Key("PORT").Int()
    if err != nil {
        panic("mysql端口转化错误")
    }

    MysqlSet.Port = port
    MysqlSet.Host = mysqlSection.Key("HOST").String()
    MysqlSet.User = mysqlSection.Key("USER").String()
    MysqlSet.Password = mysqlSection.Key("PASSWORD").String()
    MysqlSet.Name = mysqlSection.Key("DB").String()
}

func GetSingleConf(section, configname string) interface{} {
    dir, _ := os.Getwd()
    viper.SetConfigName("app.ini")
    viper.AddConfigPath(dir)
    err := viper.ReadInConfig() // Find and read the config file
    if err != nil {             // Handle errors reading the config file
        panic(fmt.Errorf("配置文件读取错误: %s \n", err))
    }
    return viper.Get(section + "." + configname)
}
