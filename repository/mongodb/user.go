package mongodb

import (
	"context"
	"github.com/QianPai/account/model"
	"github.com/QianPai/account/repository"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
)

type user struct {
	Client *mongo.Client
}

func NewUserRepository(client *mongo.Client) repository.User  {
 	return &user{client}
}

func (u *user) GetById(ctx context.Context, id interface{}) (res *model.User, err error){
	collection := u.Client.Database("users").Collection("users")
	filter := bson.M{"_id" : id}
	err = collection.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		log.Printf("can't get user by id:  %v, err: %v", id, err)
	}
	return
}

func (u *user) GetByPhone(ctx context.Context, phone string) (res *model.User, err error){
	collection := u.Client.Database("users").Collection("users")
	filter := bson.M{"phone" : phone}
	err = collection.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		log.Printf("can't get user by phone:  %v, err: %v", phone, err)
	}
	return
}

func (u *user) GetByName(ctx context.Context, name string) (res *model.User, err error){
	collection := u.Client.Database("users").Collection("users")
	filter := bson.M{"name" : name}
	err = collection.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		log.Printf("can't get user by name:  %v, err: %v", name, err)
	}
	return
}

func (u *user) Update(ctx context.Context, user *model.User) (err error){
	collection := u.Client.Database("users").Collection("users")
	filter := bson.M{"_id" : user.ID}
	res, err := collection.UpdateOne(ctx, filter, user)
	if err != nil {
		log.Printf("can't get user by id:  %v, err: %v", '2', err)
	}
	return
}

func (u *user) Store(ctx context.Context, user *model.User) (err error){
	collection := u.Client.Database("users").Collection("users")
	data := bson.M{
		"name": user.Name,
		"phone": user.Phone,
		"password": user.Password,
		"created_at": user.CreatedAt,
		"updated_at": user.CreatedAt,
	}
	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		log.Printf("insert new user fail:  %v, err: %v", user.Phone, err)
	}
	user.ID = res.InsertedID
	return
}

func (u *user) Delete(ctx context.Context, id interface{}) (err error){
	collection := u.Client.Database("users").Collection("users")
	filter := bson.M{"id" : id}
	var res model.User
	err = collection.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		log.Printf("can't get user by id:  %v, err: %v", id, err)
	}
	return
}


