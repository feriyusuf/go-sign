package models_mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Session struct {
	ID        primitive.ObjectID `json:"_id,omitempty"`
	Username  string             `json:"username"`
	IsActive  bool               `json:"is_active"`
	ExpiredAt time.Time          `json:"expired_at"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

func CreateSession(username string, expiredAt time.Time) error {
	var sessionCollection = OpenCollection(Client, "session")
	var ctx, _ = context.WithTimeout(context.Background(), 100*time.Second)
	var session Session

	session.ExpiredAt = expiredAt
	session.Username = username
	session.IsActive = true
	session.CreatedAt = time.Now()
	session.UpdatedAt = time.Now()

	_, err := sessionCollection.InsertOne(ctx, session)

	return err
}
