# mygolang
package controllers

import (
	"log"

	"github.com/astaxie/beego"
)

//定义session中的key值
const SESSION_USER_KEY = "SESSION_USER_KEY"

// 约定：如果子controller 存在NestPrepare()方法，就实现了该接口，
type NestPreparer interface {
	NestPrepare()
}
type BaseController struct {
	beego.Controller
	IsLogin bool        //标识 用户是否登陆
	User    models.User //登陆的用户
}

func (ctx *BaseController) Prepare() {
	log.Println("BaseControll")
	ctx.Data["Path"] = ctx.Ctx.Request.RequestURI
	// 判断子类是否实现了NestPreparer接口，如果实现了就调用接口方法。
	if app, ok := ctx.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}

	// 验证用户是否登陆，判断session中是否存在用户，存在就已经登陆，不存在就没有登陆。
	ctx.IsLogin = false
	tu := ctx.GetSession(SESSION_USER_KEY)
	if tu != nil {
		if u, ok := tu.(models.User); ok {
			ctx.User = u
			ctx.Data["User"] = u
			ctx.IsLogin = true
		}
	}
	ctx.Data["IsLogin"] = ctx.IsLogin
}
