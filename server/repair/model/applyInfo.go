package repairModel

import "time"

// ApplyInfo 报修申请信息模型
type ApplyInfo struct {
	// ID 序号
	ID uint `json:"id" gorm:"primary_key"`
	// Name 姓名
	Name string `json:"name" gorm:"type:text;not null"`
	// Gender 性别
	Gender string `json:"gender" gorm:"type:text;not null"`
	// Academy 学院
	Academy string `json:"academy" gorm:"type:text;not null"`
	// CardId 学工号
	CardId string `json:"cardId" gorm:"type:text;not null"`
	// Email 邮箱
	Email string `json:"email" gorm:"type:text;not null"`
	// ContactType 即时通讯联系方式类型
	ContactType string `json:"contactType" gorm:"type:text;not null"`
	// Contact 即时通讯联系方式
	Contact string `json:"contact" gorm:"type:text;not null"`
	// ComputerType 电脑型号
	ComputerType string `json:"computerType" gorm:"type:text"`
	// Problem 遇到的问题
	Problem string `json:"problem" gorm:"type:text;not null"`
	// LocationType 门牌号类型
	LocationType string `json:"locationType" gorm:"type:text;not null"`
	// Location 门牌号
	Location string `json:"location" gorm:"type:text;not null"`
	// AdminId 外键 接取当前预约的管理员主键
	AdminId uint `json:"adminId"`
	// Admin 接取当前预约的管理员对象
	Admin AdminInfo `gorm:"foreignKey:AdminId"`
	// CreateAt 当前预约创建时间
	CreateAt time.Time
	// IsFinish 当前预约任务是否已完成
	IsFinish bool `gorm:"not null"`
	// IsAbandoned 当前预约任务是否被放弃
	IsAbandoned bool `gorm:"not null"`
}
