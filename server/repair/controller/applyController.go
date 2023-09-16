package repairController

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

// ApplyController 报修申请控制器
type ApplyController struct {
	ctx iris.Context
}

// PostApply 前端提交报修申请的接口,Post请求方式,相对路径./apply
func (c *ApplyController) PostApply(apply repairModel.ApplyInfo) int {
	fmt.Println("apply get input:", apply)

	// 写入当前时间
	apply.CreateAt = time.Now()

	// 存入数据库
	database.DB.Create(&apply)
	fmt.Println("apply has written it in the database")

	// 利用模板引擎生成已收到申请的邮件
	go func() {
		tpl, _ := template.ParseFiles("./webapp/template/repair/ReplyUponReceipt.html")
		var body bytes.Buffer
		err := tpl.Execute(&body, iris.Map{
			"Name":    apply.Name,
			"Gender":  apply.Gender,
			"Problem": apply.Problem,
		})
		if err != nil {
			fmt.Println("template for ReplyUponReceipt.html error:", err)
		}

		// 发送申请已提交通知邮件
		err = email.Send(apply.Email, "浙理计协硬件部——硬件报修申请已提交通知", body.Bytes())
		if err != nil {
			fmt.Println("email send error:", err)
		}
		fmt.Println("replay email for apply has sent")
	}()

	// 利用模板引擎生成内部通知的邮件
	go func() {
		tpl, _ := template.ParseFiles("./webapp/template/repair/AlertForNewApplications.html")
		var body bytes.Buffer

		// 电脑型号可能为空,为了美观特殊处理
		computerType := ""
		if apply.ComputerType != "" {
			computerType = "(电脑型号:" + apply.ComputerType + ")"
		}

		err := tpl.Execute(&body, iris.Map{
			"Academy":      apply.Academy,
			"Name":         apply.Name,
			"Gender":       apply.Gender,
			"Problem":      apply.Problem,
			"ComputerType": computerType,
		})
		if err != nil {
			fmt.Println("template for AlertForNewApplications.html error:", err)
		}

		// 发送内部通知邮件
		err = email.Send("hardwaregroup@zstuca.club", "浙理计协硬件部——有新申请已提交通知", body.Bytes())
		if err != nil {
			fmt.Println("email send error:", err)
		}
		fmt.Println("alert email for apply has sent")
	}()

	return iris.StatusCreated
}
