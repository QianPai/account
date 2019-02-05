package main

import (
	"context"
	"fmt"
	"github.com/GordonChen13/QianPaiRen/account/repository/mongodb"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
	"github.com/sirupsen/logrus"
	"time"
)

func main()  {
	client, err := mongo.NewClient("mongodb://localhost:27017")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
        logrus.Error("连接mongodb出问题了")
	}
	userRepository := mongodb.New(client)
	user, err := userRepository.GetByName(ctx, "Gordon")
	fmt.Println(user.Phone)
	//collection := client.Database("users").Collection("users")
	//ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	//res, err := collection.InsertOne(ctx, bson.M{"name" : "Gordon", "phone" : "13420986973"})
	//id := res.InsertedID
	//println(id)
}
