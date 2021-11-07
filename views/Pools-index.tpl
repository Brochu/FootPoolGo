<main role="main" class="container">

  <div class="starter-template">
    <h2>Pour saison {{.season}} - semaine {{.week}}</h2>

    <p class="lead">LISTING POOLS HERE</p>
    <!--
    backed by this struct:
    type Match struct {
        EventId string
        EventName string
        EventNameAlt string

        Season int
        Week int

        HomeTeam string
        HomeScore int
        AwayTeam string
        AwayScore int
    }
    -->
    <div class="match-list">
        {{  range .matches }}
            <div>{{.AwayTeam}}</div>
            <div>{{.AwayScore}}</div>

            <div>{{.HomeScore}}</div>
            <div>{{.HomeTeam}}</div>

            <div>{{.EventId}}</div>
        {{ end }}
    </div>

    <p>{{.picks}}</p>
    <p>for user/pooler: </p>
    <p>{{.user}} / {{.pooler}}</p>
  </div>

</main>
