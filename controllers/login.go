package controllers

import (
    "log"
    "encoding/json"

	beego "github.com/beego/beego/v2/server/web"

    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
)

var conf *oauth2.Config

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
    test, err := beego.AppConfig.String("oauth2GoogleConfig")

    if err != nil {
        log.Printf("%v", err)
        c.Abort("Could not parse config!")
    }

    var result map[string]interface{}
    json.Unmarshal([]byte(test), &result)

    // handle login request
    conf = &oauth2.Config{
        ClientID: result["client_id"].(string),
        ClientSecret: result["client_secret"].(string),
        RedirectURL: result["redirect_uri"].(string),
        Scopes: []string{
            "https://www.googleapis.com/auth/userinfo.email",
        },
        Endpoint: google.Endpoint,
    }

    c.Data["loginURL"] = conf.AuthCodeURL("abc")
	c.TplName = "loginTemplate.tpl"
    c.Render()
}

func (c *LoginController) Callback() {
    // Handle results of the login
    c.Redirect("/pools", 302)
}

