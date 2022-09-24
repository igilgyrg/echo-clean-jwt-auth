package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init(config *mongoConfig) (db *mongo.Database, err error) {
	url := fmt.Sprintf("mongodb://%s:%s@%s:%s", config.username, config.password, config.host, config.port)

	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect mongo db: %v", err)
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, fmt.Errorf("failed to ping mongo db: %v", err)
	}

	db = client.Database(config.database)

	return
}
