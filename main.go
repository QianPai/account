package main

import (
	"context"
	"fmt"
	"github.com/QianPai/account/repository"
	"github.com/QianPai/account/repository/mongodb"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/spf13/viper"
	"time"
)

func init()  {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service Run on DEBUG mode")
	}
}

func main()  {
	dbClient,err := repository.NewDefaultClient()
	userRepository := mongodb.NewUserRepository(dbClient.(*mongo.Client))
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	user, err := userRepository.GetByName(ctx, "Gordon")
	fmt.Println(user.Phone, err)
	//collection := client.Database("users").Collection("users")
	//ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	//res, err := collection.InsertOne(ctx, bson.M{"name" : "Gordon", "phone" : "13420986973"})
	//id := res.InsertedID
	//println(id)
}
