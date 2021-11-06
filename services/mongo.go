package services

import (
    "context"
    "fmt"
    "log"
    "time"
    "encoding/json"

	beego "github.com/beego/beego/v2/server/web"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
    prim "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type DBContext struct {
    Client *mongo.Client
    Data *mongo.Database
}

type User struct {
    ID prim.ObjectID `bson:"_id"`
    Email string `bson:"email"`
    AccessLevel string `bson:"accesslevel"`
    Token string `bson:"token"`
    RefreshToken string `bson:"refreshtoken"`
}

type Pooler struct {
    ID prim.ObjectID `bson:"_id"`
    Name string `bson:"name"`
    FavTeam string `bson:"favTeam"`
    PoolId prim.ObjectID `bson:"pool_id"`
    UserId prim.ObjectID `bson:"user_id"`
}

type Pool struct {
    ID prim.ObjectID `bson:"_id"`
    Name string `bson:"name"`
    Motp string `bson:"motp"`
}

type Pick struct {
    ID prim.ObjectID `bson:"_id"`
    Season int `bson:"season"`
    Week int `bson:"week"`
    Pooler prim.ObjectID `bson:"pooler_id"`
    PickString string `bson:"pickstring"`
}

var DB DBContext

func (db* DBContext) InitDBContext() {
    dbUrlformat, _ := beego.AppConfig.String("DB_URL")
    dbUser, _ := beego.AppConfig.String("DB_USERNAME")
    dbPass, _ := beego.AppConfig.String("DB_PASSWORD")
    dbName, _ := beego.AppConfig.String("DB_NAME")

    dbUrl := fmt.Sprintf(dbUrlformat, dbUser, dbPass, dbName)
    clientOptions := options.Client().ApplyURI(dbUrl)

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatalf("[MONGO][InitDBContext] - %v", err)
    }

    DB.Client = client
    DB.Data = client.Database(dbName)
}

// Database queries
func (db* DBContext) FetchPooler(email string) (prim.ObjectID, prim.ObjectID, prim.ObjectID) {
    userCollection, _ := beego.AppConfig.String("DB_USER_COLLECT")
    filter := bson.M {
        "email": email,
    }

    var u User
    err := db.Data.Collection(userCollection).FindOne(context.TODO(), filter).Decode(&u)
    if err != nil {
        log.Fatalf("[MONGO][FetchPooler] find user - %v", err)
    }

    poolerCollection, _ := beego.AppConfig.String("DB_POOLER_COLLECT")
    poolerFilter := bson.M {
        "user_id": u.ID,
    }

    var p Pooler
    err = db.Data.Collection(poolerCollection).FindOne(context.TODO(), poolerFilter).Decode(&p)
    if err != nil {
        log.Fatalf("[MONGO][FetchPooler] find pooler - %v", err)
    }

    return u.ID, p.ID, p.PoolId
}

func (db* DBContext) FetchCurrentWeek(pooler prim.ObjectID) (int, int) {
    picksCollection, _ := beego.AppConfig.String("DB_PICK_COLLECT")

    poolerFilter := bson.M {
        "pooler_id": pooler,
    }
    opts := options.FindOneOptions{}
    opts.SetSort(bson.D{{Key:"season",Value:-1}})

    // Find latest season pooler's picks
    var ps Pick
    err := DB.Data.Collection(picksCollection).FindOne(context.TODO(), poolerFilter, &opts).Decode(&ps)
    if err != nil {
        log.Fatalf("[MONGO][FetchAllPicksCurrentWeek] find season - %v", err)
    }

    seasonFiter := bson.M {
        "pooler_id": pooler,
        "season": ps.Season,
    }
    weekOpts := options.FindOneOptions{}
    weekOpts.SetSort(bson.D{{Key:"week",Value:-1}})

    // Find latest week pooler's picks
    var pw Pick
    err = DB.Data.Collection(picksCollection).FindOne(context.TODO(), seasonFiter, &weekOpts).Decode(&pw)
    if err != nil {
        log.Fatalf("[MONGO][FetchAllPicksCurrentWeek] find week - %v", err)
    }

    return ps.Season, pw.Week
}

func (db* DBContext) FetchPoolPicks(pool prim.ObjectID, season int, week int) map[prim.ObjectID]map[string]string {
    //TODO: Get all poolers from pool
    //TODO: Fetch picks for each poolers, aggregate all into one map

    result := make(map[prim.ObjectID]map[string]string)
    result[pool] = make(map[string]string)
    result[pool]["test"] = "match"

    return result
}

func (db* DBContext) FetchPoolerPicks(pooler prim.ObjectID, season int, week int) map[string]string {
    picksCollection, _ := beego.AppConfig.String("DB_PICK_COLLECT")

    picksFilter := bson.M {
        "pooler_id": pooler,
        "season": season,
        "week": week,
    }
    log.Printf("[MongoService] Fetching all picks for season:%v; week:%v", season, week)

    var p Pick
    err := DB.Data.Collection(picksCollection).FindOne(context.TODO(), picksFilter).Decode(&p)
    if err != nil {
        log.Fatalf("[MONGO][FetchAllPicks] season |%v| week |%v| - %v", season, week, err)
    }

    var picks map[string]string
    json.Unmarshal([]byte(p.PickString), &picks)
    return picks
}

