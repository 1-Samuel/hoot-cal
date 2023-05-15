package mongodb

import (
	"context"
	"github.com/1-samuel/hoot-cal/owl"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type respoitoryMongo struct {
	db *mongo.Database
}

func NewRepositoryMongo(db *mongo.Database) owl.Repository {
	return respoitoryMongo{db}
}

func (r respoitoryMongo) Get() ([]owl.Match, error) {
	opts := options.Find().SetSort(bson.D{{"start", 1}})
	cur, err := r.db.Collection("matches").Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	var results []owl.Match
	if err = cur.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r respoitoryMongo) GetActive() ([]owl.ActiveMatch, error) {
	cur, err := r.db.Collection("activeMatches").Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	var results []owl.ActiveMatch
	if err = cur.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return results, nil
}
