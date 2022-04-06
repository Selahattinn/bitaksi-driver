package driver

import (
	"context"
	"errors"

	"github.com/Selahattinn/bitaksi-driver/internal/model"
	"github.com/sirupsen/logrus"
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
	logrus.Debug("QUERY: Get Driver DriverID: ", id, " Driver: ", d)
	return d, nil
}

func (r *MongoRepository) CreateDriver(driver *model.Driver) (int64, error) {
	// find sutiable id
	id, err := r.collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return -1, err
	}
	driver.ID = id + 1

	opt, err := r.collection.InsertOne(context.TODO(), driver)
	logrus.Debug("QUERY: Create a driver Driver: ", driver.ID, " Inserted Driver Count: ", opt)
	if err != nil {
		return -1, err
	}
	logrus.Debug("QUERY: Create a driver Driver: ", driver, " Inserted Driver Count: ", opt)
	return opt.InsertedID.(int64), nil
}

func (r *MongoRepository) UpdateDriver(driver *model.Driver) (*model.Driver, error) {
	update := bson.M{"$set": bson.M{"location": driver.Location}}
	_, err := r.collection.UpdateOne(context.TODO(), bson.M{"_id": driver.ID}, update)
	if err != nil {
		return nil, err
	}
	logrus.Debug("QUERY: Update a driver Driver: ", driver, " Updated Driver Count: ", 1)

	return driver, nil
}

func (r *MongoRepository) DeleteDriver(driver *model.Driver) (int64, error) {
	opt, err := r.collection.DeleteOne(context.TODO(), bson.M{"_id": driver.ID})
	if err != nil {
		return -1, err
	}
	logrus.Debug("QUERY: Delete a driver DriverID: ", driver.ID, " Deleted Driver Count: ", opt)
	return opt.DeletedCount, nil
}

func (r *MongoRepository) GetAllDrivers() ([]*model.Driver, error) {
	cursor, err := r.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	logrus.Debug("QUERY: Get All Drivers")
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
		return nil, err
	}
	logrus.Debug("QUERY: Find Suitable Drivers, Rider:", RiderPoint, " ,Max Distance: ", maxDistance)

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
