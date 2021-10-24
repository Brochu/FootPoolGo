package routers

import (
	"FootPoolGo/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    // Login
    beego.Router("/", &controllers.LoginController{})
    beego.Router("/auth/google_oauth2/callback", &controllers.LoginController{}, "*:Callback")

    // Content
    beego.Router("/pools", &controllers.MainController{})
}
