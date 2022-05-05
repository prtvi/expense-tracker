package config

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Global mongo collection connection
var Transactions mongo.Collection
var CurrentStatus mongo.Collection

// load env variables into the system env
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err.Error())
	}
}

// establish connection & initialize global collection connection
func EstablishConnection() {
	LoadEnv()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("DB_URL")))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Database connected")

	Transactions = *client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("TRANSACTIONS_COLL"))
	CurrentStatus = *client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("CURRENT_STATUS_COLL"))
}
