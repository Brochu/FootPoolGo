package main

import (
	_ "FootPoolGo/routers"
    "FootPoolGo/services"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
    services.InitDBContext()
	beego.Run()
}

