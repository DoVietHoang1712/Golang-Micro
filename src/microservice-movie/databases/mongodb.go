package databases

import (
	"Golang-Microservice/src/microservice-movie/common"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client       *mongo.Client
	Databasename string
}

func (db *MongoDB) Init() error {
	db.Databasename = common.Config.MgDbName
	var err error
	client, err := mongo.NewClient(options.Client().ApplyURI(""))
	if err != nil {
		return err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)
	return err
}
