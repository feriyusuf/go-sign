package models_mongo

import (
	"context"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Session struct {
	ID        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Username  string        `json:"username" bson:"username"`
	IsActive  bool          `json:"is_active" bson:"is_active"`
	ExpiredAt time.Time     `json:"expired_at" bson:"expired_at"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
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
