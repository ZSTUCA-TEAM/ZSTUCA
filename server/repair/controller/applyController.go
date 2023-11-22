package repairController

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"time"
	"zstuCA/server/database"
	repairModel "zstuCA/server/repair/model"
	"zstuCA/server/repair/tool"
)

// ApplyController 报修申请控制器
type ApplyController struct {
}

// PostApply 前端提交报修申请的接口,Post请求方式,相对路径./apply
func (c *ApplyController) PostApply(apply repairModel.ApplyInfo) int {
	fmt.Println("apply get input:", apply)

	// 获取当前时间
	apply.CreateAt = time.Now()

	// 确认是否重复提交
	var prevTimes []time.Time
	database.Get().Model(&repairModel.ApplyInfo{}).Where("name=? AND card_id=?", apply.Name, apply.CardId).Order("create_at DESC").Pluck("create_at", &prevTimes)
	if len(prevTimes) >= 1 && int(apply.CreateAt.Sub(prevTimes[0]).Seconds()) <= 10 {
		fmt.Println("apply repeatedly submitted")
		return iris.StatusConflict
	} // 最近一次提交在10s内，判定为重复提交

	// 存入数据库
	database.Get().Create(&apply)
	fmt.Println("apply has written it in the database")

	// 向用户发送收到申请的邮件
	go tool.SendInfoEmail(apply.Email, tool.MessageForSubmission, iris.Map{
		"ApplyInfo": apply,
	})

	// 利用模板引擎生成内部通知的邮件
	go tool.SendInfoEmails(tool.GetAllAdminsEmail(), tool.ReminderForSubmission, iris.Map{
		"ApplyInfo": apply,
	})

	return iris.StatusCreated
}
