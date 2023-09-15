package repairModel

// WorkList 报修任务完成信息
type WorkList struct {
	// AdminId 外键 接取当前预约的管理员主键
	AdminId uint `json:"adminId" gorm:"not null"`
	// Admin 接取当前预约的管理员对象
	Admin AdminInfo `gorm:"foreignKey:AdminId"`
	// OtherAdmin 其他参与人员
	OtherAdmin string `json:"otherAdmin"`
	// ApplyId 外键 当前预约的主键
	ApplyId uint `json:"applyId" gorm:"not null"`
	// Apply 当前预约的对象
	Apply ApplyInfo `gorm:"foreignKey:ApplyId"`
	// Time 预约完成时间
	Time string `json:"time" gorm:"not null"`
	// Location 地点
	Location string `json:"location" gorm:"not null"`
	// Duration 时长
	Duration string `json:"duration" gorm:"not null"`
	// 	WorkContent 工作内容
	WorkContent string `json:"workContent" gorm:"not null"`
}
