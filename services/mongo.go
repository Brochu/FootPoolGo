package services

import (
    "context"
    "fmt"
    "log"
    "time"

	beego "github.com/beego/beego/v2/server/web"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type DBContext struct {
    Client *mongo.Client
    Data *mongo.Database
}

type User struct {
    ID primitive.ObjectID `bson:"_id"`
    Email string `bson:"email"`
    AccessLevel string `bson:"accesslevel"`
    Token string `bson:"token"`
    RefreshToken string `bson:"refreshtoken"`
}

type Pooler struct {
    ID primitive.ObjectID `bson:"_id"`
    Name string `bson:"name"`
    FavTeam string `bson:"favTeam"`
    PoolId primitive.ObjectID `bson:"pool_id"`
    UserId primitive.ObjectID `bson:"user_id"`
}

type Pool struct {
    ID primitive.ObjectID `bson:"_id"`
    Name string `bson:"name"`
    Motp string `bson:"motp"`
}

type Pick struct {
    ID primitive.ObjectID `bson:"_id"`
    Season int `bson:"season"`
    Week int `bson:"week"`
    Pooler primitive.ObjectID `bson:"pooler_id"`
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
        log.Fatal(err)
    }

    DB.Client = client
    DB.Data = client.Database(dbName)
}

// Database queries
func (db* DBContext) FetchPooler(email string) (primitive.ObjectID, primitive.ObjectID) {
    userCollection, _ := beego.AppConfig.String("DB_USER_COLLECT")
    filter := bson.M{
        "email": email,
    }

    var u User
    err := db.Data.Collection(userCollection).FindOne(context.TODO(), filter).Decode(&u)
    if err != nil {
        log.Fatalf("[MONGO] Could not parse results for FetchPooler!")
    }

    poolerCollection, _ := beego.AppConfig.String("DB_POOLER_COLLECT")
    poolerFilter := bson.M {
        "user_id": u.ID,
    }

    var p Pooler
    err = db.Data.Collection(poolerCollection).FindOne(context.TODO(), poolerFilter).Decode(&p)
    if err != nil {
        log.Fatalf("[MONGO] Could not parse results for FetchPooler!")
    }

    return u.ID, p.ID
}

func (db* DBContext) FetchAllPicksCurrentWeek(pooler primitive.ObjectID) []Pick {
    picksCollection, _ := beego.AppConfig.String("DB_PICK_COLLECT")

    poolerFilter := bson.M{
        "pooler_id": pooler,
    }
    opts := options.FindOneOptions{}
    opts.SetSort(bson.D{{Key:"season",Value:-1}})

    // Find latest season pooler's picks
    var ps Pick
    err := DB.Data.Collection(picksCollection).FindOne(context.TODO(), poolerFilter, &opts).Decode(&ps)
    if err != nil {
        log.Fatalf("[MONGO] Could not parse results for getting most recent picks season")
    }

    seasonFiter := bson.M{
        "pooler_id": pooler,
        "season": ps.Season,
    }
    weekOpts := options.FindOneOptions{}
    weekOpts.SetSort(bson.D{{Key:"week",Value:-1}})

    // Find latest week pooler's picks
    var pw Pick
    err = DB.Data.Collection(picksCollection).FindOne(context.TODO(), seasonFiter, &weekOpts).Decode(&pw)
    if err != nil {
        log.Fatalf("[MONGO] Could not parse results for getting most recent picks week")
    }

    season, week := ps.Season, pw.Week

    // Call FetchAllPicks for current pooler and found season and week
    return db.FetchAllPicks(pooler, season, week)
}

func (db* DBContext) FetchAllPicks(pooler primitive.ObjectID, season int, week int) []Pick {
    log.Printf("[MongoService] Fetching all picks for season:%v; week:%v", season, week)

    //TODO: Find all picks for the pooler's pool
    return make([]Pick, 16)
}

