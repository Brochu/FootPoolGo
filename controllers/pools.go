package controllers

import (
    "log"
    "FootPoolGo/services"

	beego "github.com/beego/beego/v2/server/web"
)

type PoolsController struct {
	beego.Controller
}

func (c *PoolsController) Get() {

    res := services.FetchMathes(2021, 1)
    for _, m := range res {
        log.Printf("* %v\n", m)
    }

    teams := services.GetTeams()
    for _, t := range teams {
        log.Printf("* %v = %v\n", t, services.GetLonganme(t))
    }

    weeks := []int { 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22 }
    for _, w := range weeks {
        log.Printf("* %v : %v\n", w, services.GetWeekLongName(w))
    }

    //log.Printf("info: %v; %v", c.GetSession("userId"), c.GetSession("poolerId"))

    c.TplName = "listPoolsTemplate.tpl"
    c.Layout = "layout.html"

    c.Render();
}

