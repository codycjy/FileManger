package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/url"

	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DB struct {
		Host       string            `yaml:"host"`
		Port       string            `yaml:"port"`
		Username   string            `yaml:"username"`
		Password   string            `yaml:"password"`
		Database   string            `yaml:"database"`
		Parameters map[string]string `yaml:"parameters"`
	} `yaml:"db"`
}

func main() {
	// 读取配置文件
	configFile, err := ioutil.ReadFile("test/MysqlTest/config.yaml")
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	// 解析配置文件
	var config Config
	if err := yaml.Unmarshal(configFile, &config); err != nil {
		log.Fatalf("failed to parse config file: %v", err)
	}

	// 构建数据库连接字符串
	params := url.Values{}
	for k, v := range config.DB.Parameters {
		params.Set(k, v)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", config.DB.Username, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.Database, params.Encode())


	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 使用db进行数据库操作
	// ...

	// 关闭数据库连接
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to close database: %v", err)
	}
	sqlDB.Close()
	fmt.Println("Test Mysql Connection Success!")

}
