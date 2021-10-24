package controllers

import (
    "context"
    "log"
    "FootPoolGo/services"

	beego "github.com/beego/beego/v2/server/web"

    "go.mongodb.org/mongo-driver/bson"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Layout = "layout.html"
	c.TplName = "index.tpl"

    test := services.DB.Data.Collection("poolers")
    cur, _ := test.Find(context.TODO(), bson.D{})
    for cur.Next(context.TODO()) {
        var result bson.D
        err := cur.Decode(&result)
        if err != nil { log.Fatal(err) }
        log.Printf("%v", result.Map()["name"])
    }
}

