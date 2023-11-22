package repairController

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"golang.org/x/crypto/bcrypt"
	"zstuCA/server/database"
	repairModel "zstuCA/server/repair/model"
	"zstuCA/server/repair/tool"
)

// BackstageController 硬件部后端控制器
type BackstageController struct {
}

// Get 硬件部后端登录界面,Get请求方式
func (c *BackstageController) Get() mvc.Result {
	return mvc.View{
		Name: "repair/RepairBackstageLogin.html",
		Data: iris.Map{
			"IsWrong": false,
		},
	}
}

// Post 硬件部后端主界面,Post请求方式
func (c *BackstageController) Post(ctx iris.Context) mvc.Result {
	// 获取session
	sess := sessions.Get(ctx)
	// 获取POST请求参数中的密码
	username, password := ctx.FormValue("username"), ctx.FormValue("password")
	// 如果请求中没有,再从session中取
	if username == "" || password == "" {
		username = sess.GetString("username")
		password = sess.GetString("password")
	}
	fmt.Println("admin try to login,username:", username, " password:", password)

	// 获取管理员对象
	admin, err := tool.GetAdmin(username, password)
	// 不存在用户名或密码错误
	if err != nil {
		fmt.Println("admin wrote wrong username or password")
		return mvc.View{
			Name: "repair/RepairBackstageLogin.html",
			Data: iris.Map{
				"IsWrong": true,
			},
		}
	}

	// 将管理员信息存入session
	sess.Set("username", username)
	sess.Set("password", password)
	fmt.Println("success to set admin info in session")

	// 返回主界面
	return mvc.View{
		Name: "repair/RepairBackstage.html",
		Data: iris.Map{
			"IsRootAdmin": admin.IsRootAdmin,
		},
	}
}

// GetDivBy 管理员获取指定页面,Get请求方式,相对路径./div/{divName: string}
func (c *BackstageController) GetDivBy(ctx iris.Context, divName string) mvc.Result {
	// 获取当前管理员对象
	sess := sessions.Get(ctx)
	admin, err := tool.GetAdmin(sess.GetString("username"), sess.GetString("password"))
	if err != nil {
		fmt.Println("can't get admin info from session")
		// 当前管理员不存在,需要重新登录
		return mvc.Response{
			Code: iris.StatusSeeOther,
			Path: "/repair/bs",
		}
	}

	// 返回相应界面
	if divName == "pendingTasks" { // 接取新的委托
		// 读取未接取的预约信息
		var applyInfo []repairModel.ApplyInfo
		database.Get().Where("(admin_id IS NULL OR admin_id = 0) AND is_abandoned = 0").Find(&applyInfo)
		fmt.Println("admin ", admin.ID, " change div to ", divName)
		return mvc.View{
			Name: "repair/PendingTasks.html",
			Data: iris.Map{
				"ApplyInfo": applyInfo,
			},
		}
	} else if divName == "receiptedTasks" { // 查看已接取的委托
		// 读取当前管理员已接取的预约信息
		var applyInfo []repairModel.ApplyInfo
		// 获取查询参数
		showFinish, _ := ctx.URLParamBool("showFinish")
		showAbandoned, _ := ctx.URLParamBool("showAbandoned")
		if showFinish && showAbandoned {
			// 全部显示
			database.Get().Where("admin_id = ?", admin.ID).Find(&applyInfo)
		} else if showFinish {
			// 不显示已放弃
			database.Get().Where("admin_id = ? AND is_abandoned = false", admin.ID).Find(&applyInfo)
		} else if showAbandoned {
			// 不显示已完成
			database.Get().Where("admin_id = ? AND is_finish = false", admin.ID).Find(&applyInfo)
		} else {
			// 均不显示
			database.Get().Where("admin_id = ? AND is_abandoned = false AND is_finish = false", admin.ID).Find(&applyInfo)
		}
		fmt.Println("admin ", admin.ID, " change div to ", divName)
		return mvc.View{
			Name: "repair/ReceiptedTasks.html",
			Data: iris.Map{
				"ApplyInfo": applyInfo,
			},
		}
	} else if !admin.IsRootAdmin { // 剩下的界面只有管理员可以访问
		return mvc.Response{
			Code: iris.StatusBadRequest,
		}
	} else if divName == "allTasks" { // 查看所有委托信息
		fmt.Println("admin ", admin.ID, " change div to ", divName)
		var applyInfo []repairModel.ApplyInfo
		// 获取查询参数
		showFinish, _ := ctx.URLParamBool("showFinish")
		showAbandoned, _ := ctx.URLParamBool("showAbandoned")
		if showFinish && showAbandoned {
			// 全部显示
			database.Get().Preload("Admin").Find(&applyInfo)
		} else if showFinish {
			// 不显示已放弃
			database.Get().Preload("Admin").Where("is_abandoned = false").Find(&applyInfo)
		} else if showAbandoned {
			// 不显示已完成
			database.Get().Preload("Admin").Where("is_finish = false").Find(&applyInfo)
		} else {
			// 均不显示
			database.Get().Preload("Admin").Where("is_abandoned = false AND is_finish = false").Find(&applyInfo)
		}
		return mvc.View{
			Name: "repair/AllTasks.html",
			Data: iris.Map{
				"ApplyInfo": applyInfo,
			},
		}
	} else if divName == "adminRegister" { // 注册硬件部成员
		fmt.Println("admin ", admin.ID, " change div to ", divName)
		return mvc.View{
			Name: "repair/AdminRegister.html",
			Data: iris.Map{},
		}
	}
	return mvc.Response{
		Code: iris.StatusBadRequest,
	}
}

// PostReceive 管理员接取预约,Post请求方式，相对路径./receive
func (c *BackstageController) PostReceive(ctx iris.Context, applyInfo repairModel.ApplyInfo) mvc.Result {
	// 获取请求参数中的预约信息ID
	fmt.Println("admin try to receive the ", applyInfo.ID, " apply")

	// 获取当前管理员对象
	sess := sessions.Get(ctx)
	admin, err := tool.GetAdmin(sess.GetString("username"), sess.GetString("password"))
	if err != nil {
		fmt.Println("can't get admin info from session")
		// 当前管理员不存在,需要重新登录
		return mvc.Response{
			Code: iris.StatusSeeOther,
			Path: "/repair/bs",
		}
	}

	// 取出当前接取的预约信息
	if err := database.Get().Where(&applyInfo).Take(&applyInfo).Error; err != nil {
		fmt.Println("this ID doesn't right")
		return mvc.Response{
			Code:        iris.StatusBadRequest,
			ContentType: "text/html",
		}
	}

	if applyInfo.AdminId != 0 {
		// 该预约已被其他管理员领取
		fmt.Println("this ", applyInfo.ID, " apply had been received yet")
		return mvc.Response{
			Code: iris.StatusOK,
			Text: "该委托已被接取",
		}
	} else if applyInfo.IsAbandoned {
		// 当前预约已被放弃
		fmt.Println("this ", applyInfo.ID, " apply had been abandoned")
		return mvc.Response{
			Code:        iris.StatusBadRequest,
			ContentType: "text/html",
		}
	}

	// 绑定外键
	applyInfo.Admin = admin
	// 改入数据库
	database.Get().Save(&applyInfo)
	fmt.Println("admin ", admin.ID, " successfully received the ", applyInfo.ID, " apply")

	// 利用模板引擎生成通知用户申请已被接取的邮件
	go tool.SendInfoEmail(applyInfo.Email, tool.MessageForReception, iris.Map{
		"ApplyInfo": applyInfo,
		"Admin":     admin,
	})

	// 将详细信息页面发送到该管理员的邮箱中
	go tool.SendInfoEmail(admin.Email, tool.ReminderForReception, iris.Map{
		"ApplyInfo": applyInfo,
		"Admin":     admin,
	})

	return mvc.Response{
		Code: iris.StatusOK,
		Text: "成功接取委托,已将委托人详细个人信息发送至你的邮箱,也可在已接取委托界面查看详细信息.感谢你为硬件部做出的贡献",
	}
}

// GetFinishBy 获取提交任务申报界面,Get请求方式,相对路径:./finish/{id: int}
func (c *BackstageController) GetFinishBy(ctx iris.Context, id int) mvc.Result {
	// 获取请求参数中的预约信息ID
	fmt.Println("admin try to finish the ", id, " apply")

	// 获取当前管理员对象
	sess := sessions.Get(ctx)
	_, err := tool.GetAdmin(sess.GetString("username"), sess.GetString("password"))
	if err != nil {
		fmt.Println("can't get admin info from session")
		// 当前管理员不存在,需要重新登录
		return mvc.Response{
			Code: iris.StatusSeeOther,
			Path: "/repair/bs",
		}
	}

	return mvc.View{
		Name: "repair/CompleteInfo.html",
		Data: iris.Map{
			"ID": id,
		},
	}
}

// PostFinish 管理员提交任务完成申报,Post请求方式,相对路径./finish
func (c *BackstageController) PostFinish(ctx iris.Context, workList repairModel.WorkList) mvc.Result {
	// 取出当前登录的管理员对象
	sess := sessions.Get(ctx)
	admin, err := tool.GetAdmin(sess.GetString("username"), sess.GetString("password"))
	if err != nil {
		fmt.Println("can't get admin info from session")
		// 当前管理员不存在,需要重新登录
		return mvc.Response{
			Code: iris.StatusSeeOther,
			Path: "/repair/bs",
		}
	}

	// 取出当前预约信息对象
	var applyInfo repairModel.ApplyInfo
	err = database.Get().Where("id = ?", workList.ApplyId).Take(&applyInfo).Error
	// 该预约不属于当前管理员或业已完成
	if err != nil || applyInfo.AdminId != admin.ID || applyInfo.IsFinish || applyInfo.IsAbandoned {
		fmt.Println("this ", applyInfo.ID, " apply not received by the ", admin.ID, " admin or had been finished or had been abandoned")
		return mvc.Response{
			Code:        iris.StatusBadRequest,
			ContentType: "text/html",
		}
	}

	// 创建当前报修任务完成信息
	workList.Admin = admin
	workList.Apply = applyInfo
	database.Get().Create(&workList)

	// 标记当前任务已完成
	applyInfo.IsFinish = true
	database.Get().Save(&applyInfo)

	fmt.Println("admin ", admin.ID, " finish the ", applyInfo.ID, " apply")

	// 利用模板引擎生成向用户发送的以便用户核查具体情况的邮件
	go tool.SendInfoEmail(applyInfo.Email, tool.MessageForFinish, iris.Map{
		"WorkList": workList,
	})

	return mvc.Response{
		Code: iris.StatusOK,
		Text: "恭喜你成功完成当前任务,再次感谢您作为浙江理工大学计算机协会硬件部的一份子为浙江理工大学做出的贡献!",
	}
}

// PostAbandoned 管理员放弃任务,Post请求方式,相对路径./abandoned
func (c *BackstageController) PostAbandoned(ctx iris.Context, applyInfo repairModel.ApplyInfo) mvc.Result {
	// 取出当前登录的管理员对象
	sess := sessions.Get(ctx)
	admin, err := tool.GetAdmin(sess.GetString("username"), sess.GetString("password"))
	if err != nil {
		fmt.Println("can't get admin info from session")
		// 当前管理员不存在,需要重新登录
		return mvc.Response{
			Code: iris.StatusSeeOther,
			Path: "/repair/bs",
		}
	}

	// 取出当前接取的预约信息
	if err := database.Get().Where(&applyInfo).Take(&applyInfo).Error; err != nil {
		fmt.Println("this ID doesn't right")
		return mvc.Response{
			Code:        iris.StatusBadRequest,
			ContentType: "text/html",
		}
	}

	// 任务信息不正常
	if admin.ID != applyInfo.AdminId || applyInfo.IsFinish || applyInfo.IsAbandoned {
		fmt.Println("apply info is wrong")
		return mvc.Response{
			Code:        iris.StatusBadRequest,
			ContentType: "text/html",
		}
	}

	// 改入数据库
	applyInfo.IsAbandoned = true
	database.Get().Save(&applyInfo)
	fmt.Println("admin ", admin.ID, " successful abandoned the ", applyInfo.ID, " apply")

	// 向用户发送申请被放弃通知邮件
	go tool.SendInfoEmail(applyInfo.Email, tool.MessageForAbandoned, iris.Map{
		"ApplyInfo": applyInfo,
	})

	return mvc.Response{
		Code: iris.StatusOK,
		Text: "成功拒绝委托",
	}
}

// PostAdmin 注册管理员,Post请求方式,相对路径./admin
func (c *BackstageController) PostAdmin(ctx iris.Context, adminInfo repairModel.AdminInfo) mvc.Result {
	sess := sessions.Get(ctx)
	rootAdmin, err := tool.GetAdmin(sess.GetString("username"), sess.GetString("password"))
	if err != nil {
		fmt.Println("can't get admin info from session")
		// 当前管理员不存在,需要重新登录
		return mvc.Response{
			Code: iris.StatusSeeOther,
			Path: "/repair/bs",
		}
	} else if !rootAdmin.IsRootAdmin {
		// 当前管理员不具有足够的权限
		fmt.Println("permission is too low to register admin")
		return mvc.Response{
			Code: iris.StatusBadRequest,
		}
	}

	fmt.Println("admin ", adminInfo.Username, " try to register")

	// 检验用户名是否重复
	var tmp repairModel.AdminInfo
	if rows := database.Get().Where("username = ?", adminInfo.Username).Take(&tmp).RowsAffected; rows != 0 {
		fmt.Println("duplicate admin username")
		// 当前用户名重复
		return mvc.Response{
			Code: iris.StatusOK,
			Text: "用户已存在(用户名重复)",
		}
	}

	// 加密密码
	password, _ := bcrypt.GenerateFromPassword([]byte(adminInfo.Password), bcrypt.DefaultCost)
	adminInfo.Password = string(password)

	// 写入数据库
	database.Get().Create(&adminInfo)

	fmt.Println("Admin register successfully")
	return mvc.Response{
		Code: iris.StatusOK,
		Text: "用户注册成功",
	}
}

// GetSend 管理员向所有未被接取的委托发送自定义信息的通知邮件
func (c *BackstageController) GetSend(ctx iris.Context, text string) mvc.Result {
	sess := sessions.Get(ctx)
	rootAdmin, err := tool.GetAdmin(sess.GetString("username"), sess.GetString("password"))
	if err != nil {
		fmt.Println("can't get admin info from session")
		// 当前管理员不存在,需要重新登录
		return mvc.Response{
			Code: iris.StatusSeeOther,
			Path: "/repair/bs",
		}
	} else if !rootAdmin.IsRootAdmin {
		// 当前管理员不具有足够的权限
		fmt.Println("permission is too low to register admin")
		return mvc.Response{
			Code: iris.StatusBadRequest,
		}
	}

	var emails []string
	database.Get().Model(&repairModel.ApplyInfo{}).Where("admin_id = 0 OR admin_id IS NULL").Pluck("email", &emails) // 获取所有管理员的邮件
	for i := 0; i < len(emails); i += 10 {
		tmpEmails := emails[i : i+10]
		if i+10 >= len(emails) {
			tmpEmails = emails[i:]
		}

		// 网易邮箱有并发限制,所以此处不使用协程
		tool.SendInfoEmails(tmpEmails, tool.MessageCustomize, iris.Map{
			"Text": text,
		})
	}

	return mvc.Response{
		Code: iris.StatusOK,
	}
}
