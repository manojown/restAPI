package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	ServerHost    string
	MongoHost     string
	MongoUser     string
	MongoPort     string
	MongoPassword string
}

func (config *Config) initialize() {
	config.ServerHost = os.Getenv("server_host")
	// config.MongoHost = os.Getenv("mongo_host")
	// config.MongoUser = os.Getenv("mongo_user")
	// config.MongoPort = os.Getenv("mongo_port")
	// config.MongoPassword = os.Getenv("mongo_password")
}

func (config *Config) MongoUri() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s",
		config.MongoUser,
		config.MongoPassword,
		config.MongoHost,
		config.MongoPort,
	)
}

func (config *Config) Connect(database string) *mongo.Database {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://manoj:manoj@cluster0.6jvp2.mongodb.net/myFirstDatabase?retryWrites=true&w=majority",
	))
	if err != nil {
		log.Fatal(err)
	}
	return client.Database(database)

}
func NewConfig() *Config {
	config := new(Config)
	config.initialize()
	return config
}
