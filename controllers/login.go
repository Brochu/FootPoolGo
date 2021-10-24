package controllers

import (
    "FootPoolGo/services"
	beego "github.com/beego/beego/v2/server/web"
)


type LoginController struct {
	beego.Controller
}

func (c *LoginController) Login() {

    c.Data["loginURL"] = services.GetAuthURLFromConf("abc")
	c.TplName = "loginTemplate.tpl"
    c.Render()
}

func (c *LoginController) Callback() {

    // Handle results of the login
    c.Redirect("/pools", 302)
}

