package dao

import (
	"time"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"app/dao/daoModel"
)

var dbPool*gorm.DB

func ConnectDB() {
	// 配置数据库连接字符串
	dsn := "root:123456@tcp(127.0.0.1:3306)/MyDataBase?charset=utf8mb4&parseTime=True&loc=Local"

	//建立连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // 获取通用数据库对象 sql.DB ，以便配置连接池
    sqlDB, err := db.DB()
    if err != nil {
        panic("failed to get generic database object")
    }

    // 设置连接池参数
    sqlDB.SetMaxIdleConns(10)   // 设置空闲连接池中连接的最大数量
    sqlDB.SetMaxOpenConns(100)  // 设置数据库的最大打开连接数
    sqlDB.SetConnMaxLifetime(time.Hour) // 设置连接可复用的最长时间

    // 自动迁移模式
    db.AutoMigrate(&daoModel.UserModel{})

	dbPool = db
}
