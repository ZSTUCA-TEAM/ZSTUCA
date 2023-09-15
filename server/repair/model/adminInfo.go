package repairModel

type AdminInfo struct {
	// ID 序号
	ID uint `json:"id" gorm:"primary_key"`
	// Username 用户名
	Username string `json:"username" gorm:"type:text;not null"`
	// Password 密码(需加密)
	Password string `json:"password" gorm:"type:text;not null"`
	// IsRootAdmin 是否是顶级管理员(硬件部部长)
	IsRootAdmin bool `json:"isRootAdmin" gorm:"not null"`
	// Name 姓名
	Name string `json:"name" gorm:"type:text;not null"`
	// Gender 性别
	Gender string `json:"gender" gorm:"type:text;not null"`
	// Email 邮箱
	Email string `json:"email" gorm:"type:text;not null"`
	// QQ QQ号
	QQ string `json:"qq" gorm:"type:text;not null"`
	// WeChat 微信号
	WeChat string `json:"WeChat" gorm:"type:text;not null"`
	// Phone 手机号
	Phone string `json:"phone" gorm:"type:text;not null"`
}
