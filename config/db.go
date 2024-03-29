package config

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

	Transactions = *client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("TRANSACTIONS"))

	Summary = *client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("SUMMARY"))

	Budget = *client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("BUDGET"))

	Settings = *client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("SETTINGS"))

	fmt.Println("Database connected")
}
