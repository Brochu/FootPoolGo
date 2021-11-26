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
    Result map[primitive.ObjectID]int
}

func (c *PoolsController) Get() {

    /*
    * Could use some channels to parallelize the work here.
    */

    //poolerId := c.GetSession("poolerId").(primitive.ObjectID)
    poolId := c.GetSession("poolId").(primitive.ObjectID)

    // Could be done together?
    //season, week := services.DB.FetchCurrentWeek(poolerId)
    season, week := 2021, 10
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

        poolData[i].Picks = poolPicks[match.EventId]
        poolData[i].Result = ComputeScores(poolData[i].Picks, winner, isTie, week);
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

func ComputeScores(picks map[primitive.ObjectID]string, winner string, isTie bool, week int) map[primitive.ObjectID]int {
    scores := make(map[primitive.ObjectID]int)

    //TODO: Check for unique pick

    for poolerId, pick := range picks {
        score := GetScore(winner == pick, isTie, false, week)
        log.Printf("[%v] > %v (%v)\n", poolerId, pick, score)

        scores[poolerId] = score
    }

    return scores;
}

func GetScore(isGood bool, isTied bool, isUnique bool, week int) int {
    //TODO: Scale points by week num
    right := 2;

    switch {

    case isUnique:
        return (int)(float64(right) * 1.5);
    case isGood:
        return right;

    case isTied:
        return 1;

    default:
        return 0;
    }
}

