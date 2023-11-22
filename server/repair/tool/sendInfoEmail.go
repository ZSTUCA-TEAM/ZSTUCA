package tool

import (
	"bytes"
	"fmt"
	"github.com/kataras/iris/v12"
	"html/template"
	"time"
	"zstuCA/server/database"
	"zstuCA/server/email"
	repairModel "zstuCA/server/repair/model"
)

const (
	// MessageForSubmission 报修申请已提交至硬件部后台通知
	MessageForSubmission = iota
	// ReminderForSubmission 有新提交的申请通知
	ReminderForSubmission
	// ReminderForDelaying 接收滞留申请提醒
	ReminderForDelaying
	// MessageForReception 报修申请已被硬件部成员领取通知
	MessageForReception
	// ReminderForReception 已接取预约任务通知
	ReminderForReception
	// MessageForFinish 硬件报修申请已完成通知
	MessageForFinish
	// MessageForAbandoned 报修申请已被放弃通知
	MessageForAbandoned
	// MessageCustomize 自定义内容通知
	MessageCustomize
)

// SendInfoEmails 群发通知邮件
func SendInfoEmails(to []string, emailType int, content iris.Map) {
	// 添加标识通知类型的参数
	content["Type"] = emailType
	content["Title"] = repairModel.GetConf().Repair.InfoEmailTitle[emailType]

	// 将模板解析到字符串中
	tpl, _ := template.ParseFiles("./webapp/template/repair/InfoEmail.html")
	var body bytes.Buffer
	if err := tpl.Execute(&body, content); err != nil {
		fmt.Println("template for ", repairModel.GetConf().Repair.InfoEmailForm[emailType], " error:", err)
		return
	}

	// 发送邮件
	if err := email.SendAll(repairModel.GetConf().Repair.InfoEmailForm, to, repairModel.GetConf().Repair.EmailAddr, repairModel.GetConf().Repair.EmailPort, repairModel.GetConf().Repair.EmailPassword, repairModel.GetConf().Repair.InfoEmailTitle[emailType], body.Bytes()); err != nil {
		fmt.Println("email send error:", err)
		return
	}
	fmt.Println("email for ", repairModel.GetConf().Repair.InfoEmailTitle[emailType], " has sent")
}

// SendInfoEmail 发送通知邮件
func SendInfoEmail(to string, emailType int, content iris.Map) {
	SendInfoEmails([]string{to}, emailType, content)
}

// RemindStayRequest 提醒接取滞留委托
func RemindStayRequest() {
	fmt.Println("reminder for stay request start")
	limit := time.Now().Add(-48 * time.Hour)
	var stayApplies []repairModel.ApplyInfo
	database.Get().Where("create_at < ? AND (admin_id = 0 OR admin_id IS NULL)", limit).Find(&stayApplies)

	// 没有滞留预约信息,终止执行
	if len(stayApplies) == 0 {
		return
	}

	// 向管理员发送提醒接取滞留预约信息的邮件
	emails := GetAllAdminsEmail()
	go SendInfoEmails(emails, ReminderForDelaying, iris.Map{
		"StayApplies": stayApplies,
	})
}
