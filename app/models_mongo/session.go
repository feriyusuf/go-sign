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
	Token     string        `json:"token" bson:"token"`
	ExpiredAt time.Time     `json:"expired_at" bson:"expired_at"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}

func CreateSession(username string, expiredAt time.Time, token string) error {
	var sessionCollection = OpenCollection(Client, "session")
	var ctx, _ = context.WithTimeout(context.Background(), 100*time.Second)
	var session Session

	session.ExpiredAt = expiredAt
	session.Username = username
	session.IsActive = true
	session.CreatedAt = time.Now()
	session.UpdatedAt = time.Now()
	session.Token = token

	_, err := sessionCollection.InsertOne(ctx, session)

	return err
}

func DestroySession(username string) error {
	var sessionCollection = OpenCollection(Client, "session")
	var ctx, _ = context.WithTimeout(context.Background(), 100*time.Second)
	filter := bson.M{"username": bson.M{"$eq": username}}
	update := bson.M{"$set": bson.M{"is_active": false}}

	_, err := sessionCollection.UpdateMany(ctx, filter, update)

	return err
}

func IsActiveSession(token string) (bool, error) {
	var sessionCollection = OpenCollection(Client, "session")
	var ctx, _ = context.WithTimeout(context.Background(), 100*time.Second)
	var session Session

	filter := bson.M{
		"is_active": bson.M{"$eq": true},
		"token":     bson.M{"$eq": token},
	}

	err := sessionCollection.FindOne(ctx, filter).Decode(&session)

	return session.IsActive, err
}
