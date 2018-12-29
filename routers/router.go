package routers

import (
	"myproj/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.IndexController{})
	//注解路由 需要调用Include。
	beego.ErrorController(&controllers.ErrorController{})
	beego.Include(
		&controllers.IndexController{},
		&controllers.UserController{},
	)
	beego.AddNamespace(
		beego.NewNamespace(
			"note",
			beego.NSInclude(&controllers.NoteController{}),
		),
		beego.NewNamespace(
			"message",
			beego.NSInclude(&controllers.MessageController{}),
		),
	)
}
