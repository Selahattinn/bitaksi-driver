package driver

import (
	"context"
	"errors"

	"github.com/Selahattinn/bitaksi-driver/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrDriverNotFound = errors.New("driver not found")
)

type MongoRepository struct {
	collection *mongo.Collection
}

const (
	dataBaseName   = "bitaksi"
	collectionName = "driver"
)

func NewMongoRepository(client *mongo.Client) *MongoRepository {
	collection := client.Database(dataBaseName).Collection(collectionName)
	return &MongoRepository{
		collection: collection,
	}

}

func (r *MongoRepository) GetDriver(id int64) (*model.Driver, error) {
	d := &model.Driver{}
	err := r.collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(d)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		return nil, ErrDriverNotFound
	} else if err != nil {
		return nil, err
	}
	return d, nil
}

func (r *MongoRepository) CreateDriver(user *model.Driver) (interface{}, error) {
	opt, err := r.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	return opt.InsertedID, nil
}
