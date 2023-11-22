package tool

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"zstuCA/server/database"
	repairModel "zstuCA/server/repair/model"
)

// GetAllAdminsEmail 获取所有管理员的邮箱
func GetAllAdminsEmail() (emails []string) {
	database.Get().Model(&repairModel.AdminInfo{}).Pluck("email", &emails) // 获取所有管理员的邮件
	return
}

// GetAdmin 根据管理员用户名和密码取出管理员
func GetAdmin(username, password string) (admin repairModel.AdminInfo, err error) {
	// 根据用户名创建管理员对象
	admin = repairModel.AdminInfo{
		Username: username,
	}

	// 判断管理员信息是否正确
	rows := database.Get().Where(&admin).Take(&admin).RowsAffected // 在数据库中查找当前用户名
	if rows == 0 || bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)) != nil {
		// 不存在用户名或密码错误
		err = errors.New("admin wrote wrong username or password")
	}
	return
}
