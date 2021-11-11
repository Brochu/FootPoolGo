package controllers

import (
    "log"
    "FootPoolGo/services"

	beego "github.com/beego/beego/v2/server/web"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type PoolsController struct {
	beego.Controller
}

type MatchData struct {
    Match services.Match
    Picks map[primitive.ObjectID]string
}

func (c *PoolsController) Get() {

    /*
    * Could use some channels to parallelize the work here.
    */

    //poolerId := c.GetSession("poolerId").(primitive.ObjectID)
    poolId := c.GetSession("poolId").(primitive.ObjectID)

    // Could be done together?
    //season, week := services.DB.FetchCurrentWeek(poolerId)
    season, week := 2021, 9
    poolers := services.DB.FetchPoolersFromPool(poolId)
    // ------------------------------

    // Could be done together?
    matches := services.NflAPI.FetchMatches(season, week)
    poolPicks := services.DB.FetchPicks(poolers, season, week)
    // ------------------------------

    poolData := make([]MatchData, len(matches))
    for i, match := range matches {
        poolData[i].Match = match

        winner, isTie := GetWinner(&match)
        log.Printf("[%v] > winner = %v [%v]\n", match.EventId, winner, isTie)

        for _, pooler := range poolers {
            log.Printf("[%v] > %v\n",
                pooler.Name,
                poolPicks[match.EventId][pooler.ID],
            )
        }
        //TODO: Fill in results data
        poolData[i].Picks = poolPicks[match.EventId]
    }

    c.Layout = "layout.html"
    c.TplName = "Pools-index.tpl"

    c.Data["season"] = season
    c.Data["week"] = week

    c.Data["poolers"] = poolers;
    c.Data["pooldata"] = poolData

    c.Render()
}

func GetWinner(m* services.Match) (string, bool) {
    if m.AwayScore == m.HomeScore {
        return "", true

    } else if m.AwayScore > m.HomeScore {
        return services.NflAPI.GetShortname(m.AwayTeam), false;

    } else {
        return services.NflAPI.GetShortname(m.HomeTeam), false;
    }
}

