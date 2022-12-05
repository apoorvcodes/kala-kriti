package routes

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gominima/minima"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func SaveTollHandler() minima.Handler {
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
	fmt.Println("Successfully connected and pinged.")
	w := req.Param("read")
	b,_ := strconv.ParseBool(req.Param("bool"))
	n := req.Param("name")
	q := req.BodyValue("base")[0]
    ocr := req.Param("ocr")

	save := &Toll{
		TollName: n,
		Base64: q,
		Weight: w,
		IsOverweight: b,
		NumberPlate: ocr,
		Timestamp: time.Now().Format("2006.01.02 15:04:05"),
	}
		db := client.Database("Data").Collection(req.Param("name"))
		resp, err := db.InsertOne(context.TODO(),save)
        fmt.Print(resp)
		if err != nil {
			fmt.Print("ERROR 3")

			res.Error(404, err.Error())
			panic(err)
		}

		res.OK().JSON(save)
	}
}
