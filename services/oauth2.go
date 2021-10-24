package services

import (
    "log"
    "encoding/json"

	beego "github.com/beego/beego/v2/server/web"

    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
)

var conf *oauth2.Config

func InitOAuthConfig() {
    test, err := beego.AppConfig.String("oauth2GoogleConfig")

    if err != nil {
        log.Print("Could not parse OAuth2 config\n")
        log.Printf("%v\n", err)
        return
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
}

func GetAuthURLFromConf(state string) string {
    opts := oauth2.SetAuthURLParam("sldkfj", "dfd")
    conf.AuthCodeURL(state, opts)

    return conf.AuthCodeURL(state)
}

