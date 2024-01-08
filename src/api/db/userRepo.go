package db

import (
	"context"
	"fmt"
	"server/src/api/db/models"
	"server/src/helper"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client

	User *mongo.Collection
}

func New() (*DB, error) {
	u := helper.GetEnvVar("MONGOURI")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(u).SetServerAPIOptions(serverAPI)

	ctx, cancel_func := context.WithTimeout(context.Background(), 10*time.Second)
	cancel_func()
	c, e := mongo.Connect(ctx, opts)
	if e != nil {
		return nil, e
	}

	userCollection := c.Database("rustDB").Collection("User")

	return &DB{
		client: c,
		User:   userCollection,
	}, nil
}

//add user

func (db *DB) AddUser(name string, email string, password string) (*mongo.InsertOneResult, error) {
	newUser := models.CreateUser(name, email, password, "123")
	result, err := db.User.InsertOne(context.TODO(), newUser)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (db *DB) GetUser(email string) (models.User, error) {
	fmt.Print("Getting user with Email: ")

	usercoll := db.User
	filter := bson.D{{Key: "email", Value: email}}
	cursor, err := usercoll.Find(context.TODO(), filter)
	if err != nil {

		return models.User{}, err
	}

	var users []models.User
	if err := cursor.All(context.TODO(), &users); err != nil {
		println(err)
		return models.User{}, err

	}
	return users[0], nil
}

//get all user

func (db *DB) GetAllUsers() ([]models.User, error) {
	usercoll := db.User

	filter := bson.D{} // Empty filter to select all documents
	cursor, err := usercoll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var users []models.User
	if err := cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}
	fmt.Println(users)
	return users, nil
}

//TODO check if user name is avalibe

func (db *DB) AvalabileEmail(email string) (bool, error) {
	usercoll := db.User
	filter := bson.D{{Key: "email", Value: email}}
	cursor, err := usercoll.Find(context.TODO(), filter)
	if err != nil {

		return false, err
	}

	var users []models.User
	if err := cursor.All(context.TODO(), &users); err != nil {
		println(err)
		return false, err

	}
	if len(users) != 0 {
		return true, nil
	}
	return false, nil
}

func (db *DB) RemoveUser() {

}
