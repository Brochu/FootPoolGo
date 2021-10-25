package services

import (
    "crypto/rand"
    "encoding/json"
    "encoding/base64"
    "io/ioutil"
    "log"

	beego "github.com/beego/beego/v2/server/web"

    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
)

var conf *oauth2.Config
var queryURL string

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

    queryURL = "https://www.googleapis.com/oauth2/v3/userinfo"
}

func GetAuthURLFromConf() (string, string) {

    state := randToken()
    opts := oauth2.SetAuthURLParam("prompt", "select_account")

    return conf.AuthCodeURL(state, opts), state
}

func FetchEmail(gotState string, genState, code string) (bool, string) {

    // State check
    if gotState != genState {
        return false, "[OAUTH2] Could not fetch email: Mismatched states";
    }

    // Code check
    tok, err := conf.Exchange(oauth2.NoContext, code)
    if err != nil {
        return false, "[OAUTH2] Could not fetch email: Invalid code"
    }

    // Fetch request
    client := conf.Client(oauth2.NoContext, tok)
    info, err := client.Get(queryURL)
    if err != nil {
        return false, "[OAUTH2] Could not fetch email: Query failed"
    }

    // Read response
    defer info.Body.Close()
    data, _ := ioutil.ReadAll(info.Body)

    // Parse response
    var parsed map[string]interface{}
    json.Unmarshal([]byte(data), &parsed)

    return true, parsed["email"].(string)
}

func randToken() string {

    b := make([]byte, 32)
    rand.Read(b)

    return base64.StdEncoding.EncodeToString(b)
}

