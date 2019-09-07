package routers

import (
	"newsService/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("https://app-fb9d7beb-1dd0-4872-b746-1ccf2a6f0e69.cleverapps.io/v1",
		beego.NSNamespace("/user", beego.NSInclude(&controllers.UserController{})),
	)
	beego.AddNamespace(ns)
}
