package controllers

import (
    "log"

    "FootPoolGo/services"

	beego "github.com/beego/beego/v2/server/web"
)


type LoginController struct {
	beego.Controller
}

func (c *LoginController) Login() {

    var state = ""
    c.Data["loginURL"], state = services.GetAuthURLFromConf()
    c.SetSession("oauthState", state)


	c.TplName = "loginTemplate.tpl"
    c.Render()
}

func (c *LoginController) Callback() {

    // Handle results of the login

    success, email := services.FetchEmail(
        c.GetSession("oauthState").(string),
        c.GetString("state"),
        c.GetString("code"),
    )
    if !success {
        c.Abort("401")
        return
    }

    log.Printf("[LOGIN] Got the email here => %v\n", email)
    // use email to find userid -> poolerid
    // store poolerid in session

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

