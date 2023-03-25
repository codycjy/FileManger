package mysql

import (
	"io/ioutil"
	"log"
	"sync"

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


var (
	once sync.Once
	db   *gorm.DB
)

func GetDB() *gorm.DB {
	once.Do(func() {
		// 连接数据库
		configFile, err := ioutil.ReadFile("test/MysqlTest/test.yaml")
		if err != nil {
			log.Fatalf("failed to read config file: %v", err)
		}

		// 解析配置文件
		var config Config
		if err := yaml.Unmarshal(configFile, &config); err != nil {
			log.Fatalf("failed to parse config file: %v", err)
		}

		// 构建数据库连接字符串
		dsn :=  config.DB.Username + ":" + config.DB.Password + "@tcp(" +
			config.DB.Host + ":" + config.DB.Port + ")/" +
			config.DB.Database + "?"
		for k, v := range config.DB.Parameters {
			dsn += k + "=" + v + "&"
		}

		// 连接数据库
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect database: %v", err)
		}



	})

	return db
}

func CloseDB() error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
