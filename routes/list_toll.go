package routes

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/gominima/minima"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/mgo.v2/bson"
)



func TollListHandler() minima.Handler {
	return func(res *minima.Response, req *minima.Request) {
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("DATABASE_URL")))
	if err != nil {
		fmt.Print("ERROR 1")
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			fmt.Print("ERROR 2")
			panic(err)
		}
	}()
    
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		fmt.Print("ERROR 3")
		panic(err)
	}
	cursor, err := client.Database("Data").Collection("prod-2").Find(context.TODO(), bson.M{})
if err != nil {
    log.Fatal(err)
}
var data []*Toll
if err = cursor.All(context.TODO(), &data); err != nil {
    log.Fatal(err)
}
	res.OK().JSON(data)
	}
}
