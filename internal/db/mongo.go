package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewMongoDBClient(ctx context.Context, uri string) (*MongoDB, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	db := client.Database("flim")

	// Create indexes
	createIndexes(ctx, db)

	return &MongoDB{
		client: client,
		db:     db,
	}, nil
}

func (m *MongoDB) GetUsersCollection() *mongo.Collection {
	return m.db.Collection("users")
}

func (m *MongoDB) GetMessagesCollection() *mongo.Collection {
	return m.db.Collection("messages")
}

func (m *MongoDB) Disconnect(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}

func createIndexes(ctx context.Context, db *mongo.Database) {
	// Index for users collection
	usersCollection := db.Collection("users")
	usersCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: "username",
	})
	usersCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: "email",
	})

	// Indexes for messages collection
	messagesCollection := db.Collection("messages")
	messagesCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: "from_user_id",
	})
	messagesCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: "to_user_id",
	})
}
