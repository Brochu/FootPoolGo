package services

import (
    "context"
    "fmt"
    "log"
    "time"

	beego "github.com/beego/beego/v2/server/web"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type DBContext struct {
    Client *mongo.Client
    Data *mongo.Database
}

var DB DBContext

func InitDBContext() {
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

