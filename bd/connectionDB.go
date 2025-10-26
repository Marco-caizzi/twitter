package bd

import (
	"context"
	"fmt"

	"github.com/Marco-caizzi/twitter/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN *mongo.Client
var DataBaseName string

func ConnectDB(ctx context.Context) error {
	user := ctx.Value(models.Key("user")).(string)
	password := ctx.Value(models.Key("password")).(string)
	host := ctx.Value(models.Key("host")).(string)
	connStr := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", user, password, host)

	var clientOptions = options.Client().ApplyURI(connStr)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(" > Successfully connected to MongoDB")
	MongoCN = client
	DataBaseName = ctx.Value(models.Key("database")).(string)
	return nil
}

func CheckConnection() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	return err == nil
}
