<main role="main" class="container">

    <div class="starter-template">
        <h2>Pour saison {{.season}} - semaine {{.week}}</h2>

        <!-- <div class="match-container">
            <div class="match-table header">
                <div class="match-row">AWAY</div>
                <div class="match-row">AWAY</div>
                <div class="match-row">HOME</div>
                <div class="match-row">HOME</div>
                {{ range .poolers }}
                    <div class="match-row">{{.Name}}</div>
                {{ end }}
            </div>

            {{ range .pooldata }}
            <div class="match-table row">
                <div class="match-row away-team">{{.Match.AwayTeam}}</div>
                <div class="match-row">{{.Match.AwayScore}}</div>

                <div class="match-row">{{.Match.HomeScore}}</div>
                <div class="match-row home-team">{{.Match.HomeTeam}}</div>

                {{ range $k, $v := .Picks }}
                    <div>{{$v}}</div>
                {{ end }}
            </div>
            {{ end }}
        </div> -->

        <table class="table table-hover table-sm">
            <thead>
                <tr>
                    <th scope="col" colspan=2>AWAY</th>
                    <th scope="col" colspan=2>HOME</th>
                    {{ range .poolers }}
                        <th scope="col">{{.Name}}</th>
                    {{ end }}
                </tr>
            </thead>

            <tbody>
            {{ range .pooldata }}
                <tr>
                    <td class="away-team"> {{.Match.AwayTeam }}</td>
                    <td>{{.Match.AwayScore }}</td>

                    <td>{{.Match.HomeScore }}</td>
                    <td class="home-team">{{.Match.HomeTeam }}</td>

                    {{ range $poolerId, $pick := .Picks }}
                        <td>{{$pick}}</td>
                    {{ end }}
                </tr>
            {{ end }}
            </tbody>
        <table>
    </div>

</main>

