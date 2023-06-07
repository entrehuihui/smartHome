package databases

import (
	"log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB 连接DB
var db *gorm.DB

// GetDB 获取DB
func GetDB() *gorm.DB {
	return db
}

// InitMysql 初始化mysql连接
// autoIncrementList 起始ID设置
// dst 迁移数据库
// LogLevel 日志等级
func InitMysql() (err error) {
	user := viper.GetString("Mysql.user")
	password := viper.GetString("Mysql.password")
	host := viper.GetString("Mysql.host")
	dbname := viper.GetString("Mysql.dbname")
	if user == "" || password == "" || host == "" || dbname == "" {
		log.Fatal("Mysql 配置不齐全")
	}

	dsn := user + ":" + password + "@tcp(" + host + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
	LogLevel := logger.Error
	if viper.GetBool("Mysql.info") {
		LogLevel = logger.Info
	}

loop:
	// 循环重连数据库
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(LogLevel),
	})
	if err != nil {
		log.Println("mysql db isn't connect error :", err)
		time.Sleep(3 * time.Second)
		goto loop
	}
	// 开启打印日志
	// db.LogMode(true)

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 迁移数据表
	// return
	err = db.AutoMigrate(Userinfo{}, Device{}, Groupinfo{}, Typesinfo{}, Linkageinfo{})
	if err != nil {
		log.Fatal(err)
	}
	return
}
