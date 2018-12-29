package main

import (
	_ "myproj/routers"
	"strings"

	_ "myproj/models"

	"github.com/astaxie/beego"
)

func main() {
	initSession()
	initTemplate()
	beego.Run()
}
func initTemplate() {
	beego.AddFuncMap("equrl", func(x, y string) bool {
		s1 := strings.Trim(x, "/")
		s2 := strings.Trim(y, "/")
		return strings.Compare(s1, s2) == 0
	})
	beego.AddFuncMap("add", func(x, y int) int {
		return x + y
	})
}

func initSession() {
	//beego的session序列号是用gob的方式，因此需要将注册models.User
	gob.Register(models.User{})
	//https://beego.me/docs/mvc/controller/session.md
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "liteblog-key"
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "data/session"
}
