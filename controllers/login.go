package controllers

import (
    "FootPoolGo/services"

	beego "github.com/beego/beego/v2/server/web"
)


type LoginController struct {
	beego.Controller
}

func (c *LoginController) Login() {

    var state = ""
    c.Data["loginURL"], state = services.OAuth.GetAuthURLFromConf()
    c.SetSession("oauthState", state)


	c.TplName = "login.tpl"
    c.Render()
}

func (c *LoginController) Callback() {

    // Handle results of the login

    success, email := services.OAuth.FetchEmail(
        c.GetSession("oauthState").(string),
        c.GetString("state"),
        c.GetString("code"),
    )
    if !success {
        c.Abort("401")
        return
    }

    userId, poolerId, poolId := services.DB.FetchPooler(email)
    c.SetSession("userId", userId)
    c.SetSession("poolerId", poolerId)
    c.SetSession("poolId", poolId)

    c.Redirect("/pools", 302)
}

func (c *LoginController) Logout() {

    err := c.DestroySession()
    if err != nil {
        c.Abort("500")
        return
    }

    c.Redirect("/", 302)
}

