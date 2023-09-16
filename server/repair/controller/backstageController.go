package repairController

import (
	"bytes"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"strconv"
	"strings"
	"time"
	"zstuCA/server/database"
	"zstuCA/server/email"
	repairModel "zstuCA/server/repair/model"
)

// BackstageController 硬件部后端控制器
type BackstageController struct {
}

// Get 硬件部后端默认界面,Get请求方式
func (c *BackstageController) Get() mvc.Result {
	return mvc.View{
		Name: "repair/RepairBackstageLogin.html",
		Data: iris.Map{
			"IsWrong": false,
		},
	}
}

func (c *BackstageController) RemindStayRequest() {
	fmt.Println("reminder for stay request start")
	limit := time.Now().Add(-48 * time.Hour)
	var stayApplys []repairModel.ApplyInfo
	database.DB.Where("create_at < ? AND (admin_id = 0 OR admin_id IS NULL)", limit).Find(&stayApplys)

	// 没有滞留预约信息,终止执行
	if len(stayApplys) == 0 {
		return
	}

	// 向管理员发送提醒接取滞留预约信息的邮件
	go func() {
		tpl, _ := template.ParseFiles("./webapp/template/repair/AlertForStayApplications.html")
		var body bytes.Buffer
		err := tpl.Execute(&body, iris.Map{
			"StayApplys": stayApplys,
		})
		if err != nil {
			fmt.Println("template for AlertForStayApplications.html error:", err)
		}

		// 发送滞留预约信息的邮件
		err = email.Send("hardwaregroup@zstuca.club", "浙理计协硬件部——接收滞留申请提醒", body.Bytes())
		if err != nil {
			fmt.Println("email send error:", err)
		}
		fmt.Println("alert email for stay applications has sent")
	}()

}

// GetPanel 硬件部后端登录接口,Get请求方式,相对路径./panel
func (c *BackstageController) GetPanel(ctx iris.Context) mvc.Result {
	var admin repairModel.AdminInfo

	// 获取GET请求参数中的密码c.Session.Get("admin")
	password := ctx.URLParam("password")
	// 根据GET请求参数中的用户名创建对象
	admin = repairModel.AdminInfo{
		Username: ctx.URLParam("username"),
	}
	fmt.Println("admin try to login,username:", admin.Username, " password:", password)

	// 判断管理员信息是否正确
	rows := database.DB.Where(&admin).Take(&admin).RowsAffected // 在数据库中查找当前用户名
	if rows == 0 || bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)) != nil {
		// 不存在用户名或密码错误
		fmt.Println("admin wrote wrong username or password")
		return mvc.View{
			Name: "repair/RepairBackstageLogin.html",
			Data: iris.Map{
				"IsWrong": true,
			},
		}
	}

	// 读取未接取的预约信息
	var newApplyInfo []repairModel.ApplyInfo
	database.DB.Where("admin_id IS NULL OR admin_id = 0").Find(&newApplyInfo)

	// 读取当前管理员已接取的预约信息
	var receiveApplyInfo []repairModel.ApplyInfo
	database.DB.Where("admin_id = ?", admin.ID).Find(&receiveApplyInfo)

	return mvc.View{
		Name: "repair/RepairBackstage.html",
		Data: iris.Map{
			"IsRootAdmin":      admin.IsRootAdmin,
			"NewApplyInfo":     newApplyInfo,
			"ReceiveApplyInfo": receiveApplyInfo,
		},
	}
}

// GetReceive 管理员接取预约,相对路径./receive
func (c *BackstageController) GetReceive(ctx iris.Context) mvc.Result {
	// 获取GET请求参数中的预约信息ID
	applyID, _ := strconv.Atoi(ctx.URLParam("ID"))
	applyInfo := repairModel.ApplyInfo{
		ID: uint(applyID),
	}
	fmt.Println("admin try to receive the ", applyInfo.ID, " apply")

	// 取出当前接取的预约信息
	err := database.DB.Where(&applyInfo).Take(&applyInfo).Error
	if err != nil {
		fmt.Println("this ID doesn't right")
		return mvc.Response{
			Code: iris.StatusBadRequest,
		}
	}

	// 取出当前登录的管理员对象
	var admin repairModel.AdminInfo
	err = database.DB.Where("username = ?", ctx.URLParam("username")).Take(&admin).Error
	// 当前管理员不存在
	if err != nil || bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(ctx.URLParam("password"))) != nil {
		// 不存在用户名或密码错误
		fmt.Println("wrong admin to receive")
		return mvc.Response{
			Code: iris.StatusBadRequest,
		}
	}

	// 该预约已被其他管理员领取
	if applyInfo.AdminId != 0 {
		fmt.Println("this ", applyInfo.ID, " apply had been received yet")
		return mvc.View{
			Name: "repair/OneSentenceAlertForAdmin.html",
			Data: iris.Map{
				"AlertSentence": "该任务已被接取",
			},
		}
	}

	// 绑定外键
	applyInfo.Admin = admin
	// 改入数据库
	database.DB.Save(&applyInfo)
	fmt.Println("admin ", admin.ID, " successfully received the ", applyInfo.ID, " apply")

	// 利用模板引擎生成通知用户申请已被接取的邮件
	go func() {
		tpl, _ := template.ParseFiles("./webapp/template/repair/ReplyUponAccess.html")
		var body bytes.Buffer
		// 根据申请里的联系方式返回不同的联系方式
		var contact string
		switch applyInfo.ContactType {
		case "QQ":
			contact = admin.QQ
		case "微信":
			contact = admin.WeChat
		case "手机":
			contact = admin.Phone
		}
		// 根据管理员性别返回不同的人称
		var adminPronoun string
		switch admin.Gender {
		case "男":
			adminPronoun = "他"
		case "女":
			adminPronoun = "她"
		}
		// 根据管理员身份返回不同的职位
		var adminIdentity string
		if admin.IsRootAdmin {
			adminIdentity = "部长"
		} else {
			adminIdentity = "干事"
		}

		err := tpl.Execute(&body, iris.Map{
			"Name":          applyInfo.Name,
			"Gender":        applyInfo.Gender,
			"Problem":       applyInfo.Problem,
			"AdminName":     admin.Name,
			"AdminIdentity": adminIdentity,
			"AdminGender":   admin.Gender,
			"AdminPronoun":  adminPronoun,
			"ContactType":   applyInfo.ContactType,
			"Contact":       contact,
		})
		if err != nil {
			fmt.Println("template for ReplyUponAccess.html error:", err)
		}

		// 发送用户申请已被接取的邮件
		err = email.Send(applyInfo.Email, "浙理计协硬件部——硬件报修申请已被领取通知", body.Bytes())
		if err != nil {
			fmt.Println("email send error:", err)
		}
		fmt.Println("replay email for access has sent")
	}()

	// 返回详细信息页面,同时也将这个页面发送到该管理员的邮箱中
	// 电脑型号可能为空,为了美观特殊处理
	computerType := ""
	if applyInfo.ComputerType != "" {
		computerType = "(电脑型号:" + applyInfo.ComputerType + ")"
	}

	// 利用模板引擎生成内部通知的邮件
	go func() {
		tpl, _ := template.ParseFiles("./webapp/template/repair/AlertForReceiveApplications.html")
		var body bytes.Buffer

		err := tpl.Execute(&body, iris.Map{
			"AdminName":    admin.Name,
			"Academy":      applyInfo.Academy,
			"Name":         applyInfo.Name,
			"Gender":       applyInfo.Gender,
			"CardId":       applyInfo.CardId,
			"Problem":      applyInfo.Problem,
			"ComputerType": computerType,
			"Email":        applyInfo.Email,
			"ContactType":  applyInfo.ContactType,
			"Contact":      applyInfo.Contact,
			"LocationType": applyInfo.LocationType,
			"Location":     applyInfo.Location,
			"IsEmail":      true,
		})
		if err != nil {
			fmt.Println("template for AlertForReceiveApplications.html error:", err)
		}

		// 发送内部通知邮件
		err = email.Send(admin.Email, "浙理计协硬件部——已接取预约任务通知", body.Bytes())
		if err != nil {
			fmt.Println("email send error:", err)
		}
		fmt.Println("alert email for receive has sent")
	}()

	return mvc.View{
		Name: "repair/AlertForReceiveApplications.html",
		Data: iris.Map{
			"AdminName":    admin.Name,
			"Academy":      applyInfo.Academy,
			"Name":         applyInfo.Name,
			"Gender":       applyInfo.Gender,
			"CardId":       applyInfo.CardId,
			"Problem":      applyInfo.Problem,
			"ComputerType": computerType,
			"Email":        applyInfo.Email,
			"ContactType":  applyInfo.ContactType,
			"Contact":      applyInfo.Contact,
			"LocationType": applyInfo.LocationType,
			"Location":     applyInfo.Location,
			"IsEmail":      false,
		},
	}
}

// PostAdmin 注册管理员,相对路径./admin
func (c *BackstageController) PostAdmin(ctx iris.Context) mvc.Result {
	admin := repairModel.AdminInfo{
		Username:    ctx.FormValue("username"),
		Password:    ctx.FormValue("password"),
		IsRootAdmin: false,
		Name:        ctx.FormValue("name"),
		Gender:      ctx.FormValue("gender"),
		Email:       ctx.FormValue("email"),
		QQ:          ctx.FormValue("QQ"),
		WeChat:      ctx.FormValue("weChat"),
		Phone:       ctx.FormValue("phone"),
	}
	fmt.Println("admin ", admin.Username, " try to register")

	// 加密密码
	password, _ := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	admin.Password = string(password)

	// 检验用户名是否重复
	var tmp repairModel.AdminInfo
	rows := database.DB.Where("username = ?", admin.Username).Take(&tmp).RowsAffected
	if rows != 0 {
		fmt.Println("duplicate admin username")
		// 当前用户名重复
		return mvc.View{
			Name: "repair/OneSentenceAlertForAdmin.html",
			Data: iris.Map{
				"AlertSentence": "用户已存在(用户名重复)",
			},
		}
	}

	// 写入数据库
	database.DB.Create(&admin)

	fmt.Println("Admin register successfully")
	return mvc.View{
		Name: "repair/OneSentenceAlertForAdmin.html",
		Data: iris.Map{
			"AlertSentence": "用户注册成功",
		},
	}
}

// GetComplete 获取提交任务申报界面,相对路径:./complete
func (c *BackstageController) GetComplete() mvc.Result {
	return mvc.View{
		Name: "repair/CompleteInfo.html",
	}
}

// PostDeclaration 管理员提交任务完成申报,相对路径./declaration
func (c *BackstageController) PostDeclaration(ctx iris.Context) mvc.Result {
	// 取出当前登录的管理员对象
	var admin repairModel.AdminInfo
	err := database.DB.Where("username = ?", ctx.URLParam("username")).Take(&admin).Error
	// 当前管理员不存在
	if err != nil || bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(ctx.URLParam("password"))) != nil {
		// 不存在用户名或密码错误
		fmt.Println("wrong admin to receive")
		return mvc.Response{
			Code: iris.StatusBadRequest,
		}
	}

	// 取出当前预约信息对象
	var applyInfo repairModel.ApplyInfo
	err = database.DB.Where("id = ?", ctx.URLParam("ID")).Take(&applyInfo).Error
	// 该预约不属于当前管理员或业已完成
	if err != nil || applyInfo.AdminId != admin.ID || applyInfo.IsFinish {
		fmt.Println("this ", applyInfo.ID, " apply not received by the ", admin.ID, " admin or had been finished")
		return mvc.Response{
			Code: iris.StatusBadRequest,
		}
	}

	// 创建当前报修任务完成信息
	workList := repairModel.WorkList{
		AdminId:     admin.ID,
		Admin:       admin,
		OtherAdmin:  ctx.FormValue("otherAdmin"),
		ApplyId:     applyInfo.ID,
		Apply:       applyInfo,
		Time:        strings.Replace(ctx.FormValue("time"), "T", " ", 1),
		Location:    ctx.FormValue("location"),
		Duration:    ctx.FormValue("duration"),
		WorkContent: ctx.FormValue("workContent"),
	}
	database.DB.Create(&workList)

	// 标记当前任务已完成
	applyInfo.IsFinish = true
	database.DB.Save(&applyInfo)

	// 利用模板引擎生成向用户发送的以便用户核查具体情况的邮件
	go func() {
		tpl, _ := template.ParseFiles("./webapp/template/repair/ReplyUponFinished.html")
		var body bytes.Buffer

		err := tpl.Execute(&body, iris.Map{
			"Name":           applyInfo.Name,
			"Gender":         applyInfo.Gender,
			"Problem":        applyInfo.Problem,
			"AdminName":      admin.Name,
			"OtherAdminName": workList.OtherAdmin,
			"Time":           workList.Time,
			"Location":       workList.Location,
			"Duration":       workList.Duration,
			"WorkContent":    workList.WorkContent,
		})
		if err != nil {
			fmt.Println("template for ReplyUponFinished.html error:", err)
		}

		// 发送任务已完成通知邮件
		err = email.Send(applyInfo.Email, "浙理计协硬件部——硬件报修申请已完成通知", body.Bytes())
		if err != nil {
			fmt.Println("email send error:", err)
		}
		fmt.Println("reply email for finish has sent")
	}()

	return mvc.View{
		Name: "repair/OneSentenceAlertForAdmin.html",
		Data: iris.Map{
			"AlertSentence": "当前维修任务成功完成申报,再次感谢您作为浙江理工大学计算机协会硬件部的一份子为浙江理工大学做出的贡献!",
		},
	}
}
