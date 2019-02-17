package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

func NewClient(driver string, host string, port string, dbname string, user string, pass string) (interface{}, error) {
	var client interface{}
	var err error
	switch driver {
	case "mongodb":
        client, err = NewMongoClient(host, port)
	case "mysql" :
		client, err = sql.Conn{}, errors.New("mysql not implement")
	default:
		client, err = NewMongoClient(host, port)
	}
	
	return client, err
}

func NewDefaultClient() (interface{}, error) {
	dbDriver := viper.GetString(`database.driver`)
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)

	client, err := NewClient(dbDriver, dbHost, dbPort, dbName, dbUser, dbPass)
	return client, err
}

func NewMongoClient(host string, port string) (*mongo.Client, error) {
	url := fmt.Sprintf("mongodb://%s:%s", host, port)
	client, err := mongo.NewClient(url)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		logrus.Error("连接mongodb出问题了")
	}

	return  client, err
}

func NewMongoDefaultClient() (*mongo.Client, error) {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)

	client, err := NewMongoClient(dbHost, dbPort)
	return client, err
}
