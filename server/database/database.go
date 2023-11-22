package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"sync"
	repairModel "zstuCA/server/repair/model"
)

// database 数据库连接单例对象
var database *gorm.DB

// Get 获取数据库连接单例对象
func Get() *gorm.DB {
	var once sync.Once

	// 初始化数据库
	dbInit := func() {
		if db, err := gorm.Open(repairModel.GetConf().DatabaseName, repairModel.GetConf().DataSourceName); err != nil {
			fmt.Println("Failed to connect to database:", err)
		} else {
			database = db
		}
		// 自动迁移
		database.AutoMigrate(repairModel.ApplyInfo{}) // 报修申请信息模型数据表
		database.AutoMigrate(repairModel.AdminInfo{}) // 硬件部成员账号数据表
		database.AutoMigrate(repairModel.WorkList{})  // 报修任务完成信息数据表
	}

	once.Do(dbInit)

	return database
}

// Close 关闭数据库连接
func Close() {
	if err := database.Close(); err != nil {
		fmt.Println("Failed to close database:", err)
	}
}
