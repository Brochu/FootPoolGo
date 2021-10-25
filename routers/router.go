package routers

import (
	"FootPoolGo/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/auth/login", &controllers.LoginController{}, "get:Login")
    beego.Router("/auth/google_oauth2/callback", &controllers.LoginController{}, "*:Callback")
    beego.Router("/auth/goodbye", &controllers.LoginController{}, "*:Logout")

    beego.Router("/pools", &controllers.PoolsController{})
}
