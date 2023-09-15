package database

import (
	"github.com/jinzhu/gorm"
	repairModel "zstuCA/server/repair/model"
)

var DB *gorm.DB

// DbInit 初始化(自动迁移)数据库
func DbInit(db *gorm.DB) {
	DB = db

	// 自动迁移
	db.AutoMigrate(repairModel.ApplyInfo{}) // 报修申请信息模型数据表
	db.AutoMigrate(repairModel.AdminInfo{}) // 硬件部成员账号数据表
	db.AutoMigrate(repairModel.WorkList{})  // 报修任务完成信息数据表
}
