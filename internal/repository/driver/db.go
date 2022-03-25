package driver

import (
	"context"
	"errors"
	"fmt"

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

	// create index 2dsphere
	index := mongo.IndexModel{
		Keys: bson.M{
			"location": "2dsphere",
		},
	}
	_, err := collection.Indexes().CreateOne(context.TODO(), index)
	if err != nil {
		return nil
	}

	return &MongoRepository{
		collection: collection,
	}

}

func (r *MongoRepository) GetDriver(id int64) (*model.Driver, error) {
	d := &model.Driver{}
	err := r.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(d)
	if err == mongo.ErrNoDocuments {
		return nil, ErrDriverNotFound
	} else if err != nil {
		return nil, err
	}
	return d, nil
}

func (r *MongoRepository) CreateDriver(driver *model.Driver) (int64, error) {
	opt, err := r.collection.InsertOne(context.TODO(), driver)
	if err != nil {
		return -1, err
	}
	return opt.InsertedID.(int64), nil
}

func (r *MongoRepository) UpdateDriver(driver *model.Driver) (*model.Driver, error) {
	update := bson.M{"$set": bson.M{"location": driver.Location}}
	_, err := r.collection.UpdateOne(context.TODO(), bson.M{"_id": driver.ID}, update)
	if err != nil {
		return nil, err
	}

	return driver, nil
}

func (r *MongoRepository) DeleteDriver(driver *model.Driver) (int64, error) {
	opt, err := r.collection.DeleteOne(context.TODO(), bson.M{"_id": driver.ID})
	if err != nil {
		return -1, err
	}

	return opt.DeletedCount, nil
}

func (r *MongoRepository) GetAllDrivers() ([]*model.Driver, error) {
	cursor, err := r.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var drivers []*model.Driver
	for cursor.Next(context.TODO()) {
		var driver model.Driver
		if err := cursor.Decode(&driver); err != nil {
			return nil, err
		}
		drivers = append(drivers, &driver)
	}

	return drivers, nil
}

func (r *MongoRepository) FindSuitableDrivers(RiderPoint *model.Rider, maxDistance float64) ([]*model.Driver, error) {

	cursor, err := r.collection.Find(context.TODO(), bson.M{"location": bson.M{"$nearSphere": bson.M{"$geometry": bson.M{"type": "Point", "coordinates": []float64{RiderPoint.Lat, RiderPoint.Long}}, "$maxDistance": maxDistance}}})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var drivers []*model.Driver
	for cursor.Next(context.TODO()) {
		var driver model.Driver
		if err := cursor.Decode(&driver); err != nil {
			return nil, err
		}
		drivers = append(drivers, &driver)
	}

	return drivers, nil
}
