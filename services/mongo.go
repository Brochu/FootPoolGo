package services

import (
    "time"
    "context"

    "github.com/astaxie/beego"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    //"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DBContext struct {
    Client *mongo.Client
}

var DB DBContext

func InitDBContext() {
    dbUrl := beego.AppConfig.String("DB_URL")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, _ := mongo.Connect(ctx, options.Client().ApplyURI(dbUrl))
    DB.Client = client
    // TO TEST THIS
}
