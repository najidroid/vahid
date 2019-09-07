package routers

import (
	"newsService/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user", beego.NSInclude(&controllers.UserController{})),
	)
	beego.AddNamespace(ns)
}
