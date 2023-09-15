package server

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/basicauth"
	"github.com/kataras/iris/v12/mvc"
	"github.com/robfig/cron/v3"
	repairController "zstuCA/server/repair/controller"
)

// Handle 绑定资源
func Handle(app *iris.Application) {
	// 注册模板引擎
	app.RegisterView(iris.HTML("./webapp/template", ".html"))
	// 静态资源
	handleStatic(app)
	// 错误页面
	handleError(app)
	// 预约报修服务
	handleRepair(app)
}

// handleStatic 绑定静态资源
func handleStatic(app *iris.Application) {
	app.HandleDir("/", "./webapp/static")
	//app.HandleDir("/", "./webapp/static/old") // 兼容旧版
}

// handleError 绑定错误页面
func handleError(app *iris.Application) {
	// 400
	app.OnErrorCode(iris.StatusBadRequest, func(ctx iris.Context) {
		ctx.View("errorPage/400.html")
	})

	// 401
	app.OnErrorCode(iris.StatusUnauthorized, func(ctx iris.Context) {
		ctx.View("errorPage/401.html")
	})

	// 404
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.View("errorPage/404.html")
	})

	// 500
	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		// 渲染自定义的 500 错误页面
		ctx.View("errorPage/500.html")
	})

	// 502
	app.OnErrorCode(iris.StatusBadGateway, func(ctx iris.Context) {
		// 渲染自定义的 502 错误页面
		ctx.View("errorPage/502.html")
	})
}

// handleRepair 预约报修服务
func handleRepair(app *iris.Application) {
	// 绑定用户报修申请入口
	mvc.New(app.Party("/repair")).Handle(new(repairController.ApplyController))

	// 绑定管理员后台
	// Basic Authentication认证中间件
	auth := basicauth.Default(map[string]string{
		"hardware": "Zjlgdx2312",
	})
	bs := app.Party("/repair/bs")
	bs.Use(auth)
	mvc.New(bs).Handle(new(repairController.BackstageController))

	// 创建任务调度器,每天12点调用,提醒滞留请求
	c := cron.New()
	controller := repairController.BackstageController{}
	//controller.RemindStayRequest()
	_, err := c.AddFunc("0 12 * * *", controller.RemindStayRequest)
	if err != nil {
		fmt.Println("cron create err:", err)
	} else {
		fmt.Println("cron successfully create")
	}
	c.Start()
}
