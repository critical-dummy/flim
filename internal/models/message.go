package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FromUserID string             `bson:"from_user_id" json:"from_user_id"`
	ToUserID  string             `bson:"to_user_id" json:"to_user_id"`
	Content   string             `bson:"content" json:"content"`
	Timestamp time.Time          `bson:"timestamp" json:"timestamp"`
	IsRead    bool               `bson:"is_read" json:"is_read"`
}
