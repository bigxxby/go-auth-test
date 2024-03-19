package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	DB *mongo.Database
}

func InitDb() (*mongo.Database, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://bigxxby:bigxxby@cluster0.frjq8wo.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0").SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}

	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		return nil, err
	}
	log.Println("Connected")
	return client.Database("database"), nil
}

func (db *Database) SaveTokensToMongo(userID, accessToken, refreshToken string) error {
}

func (db *Database) GetUserIDByRefreshTokenFromMongo(refreshToken string) (string, error) {
}

func (db *Database) UpdateTokensInMongo(userID, accessToken, refreshToken string) error {
}
