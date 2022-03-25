package driver

import (
	"context"
	"reflect"
	"testing"

	"github.com/Selahattinn/bitaksi-driver/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func TestMongoRepository_GetDriver(t *testing.T) {

	uri := "mongodb://localhost:27017/?maxPoolSize=20&w=majority"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		t.Fatal(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		t.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		t.Fatal(err)
	}
	// Create a mock collection
	collection := client.Database("bitaksi").Collection("driver")
	collection.Drop(context.TODO())
	// Create a mock driver
	driver := &model.Driver{
		ID:       1,
		Location: model.GeoJson{Type: "Point", Coordinates: []float64{1, 1}},
	}

	// Insert the mock document
	_, err = collection.InsertOne(context.TODO(), driver)
	if err != nil {
		t.Fatal(err)
	}
	// Create a mock repository
	repository := NewMongoRepository(client)

	// Get the mock document
	got, err := repository.GetDriver(1)
	if err != nil {
		t.Fatal(err)
	}

	// Check if the mock document is equal to the mock driver
	if !reflect.DeepEqual(got, driver) {
		t.Fatalf("got %v, want %v", got, driver)
	}

}

func TestMongoRepository_CreateDriver(t *testing.T) {
	uri := "mongodb://localhost:27017/?maxPoolSize=20&w=majority"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		t.Fatal(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		t.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		t.Fatal(err)
	}
	// Create a mock collection
	collection := client.Database("bitaksi").Collection("driver")
	collection.Drop(context.TODO())
	// Create a mock driver
	driver := &model.Driver{
		ID:       1,
		Location: model.GeoJson{Type: "Point", Coordinates: []float64{1, 1}},
	}

	// Create a mock repository
	repository := NewMongoRepository(client)

	id, err := repository.CreateDriver(driver)
	if err != nil {
		t.Fatal(err)
	}
	d := &model.Driver{}
	err = collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(d)
	if err != nil {
		t.Fatal(err)
	}

	// Check if the mock document is equal to the mock driver
	if !reflect.DeepEqual(d, driver) {
		t.Fatalf("got %v, want %v", d, driver)
	}

}

func TestMongoRepository_DeleteDriver(t *testing.T) {
	uri := "mongodb://localhost:27017/?maxPoolSize=20&w=majority"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		t.Fatal(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		t.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		t.Fatal(err)
	}
	// Create a mock collection
	collection := client.Database("bitaksi").Collection("driver")
	collection.Drop(context.TODO())
	// Create a mock driver
	driver := &model.Driver{
		ID:       1,
		Location: model.GeoJson{Type: "Point", Coordinates: []float64{1, 1}},
	}

	// Insert the mock document
	_, err = collection.InsertOne(context.TODO(), driver)
	if err != nil {
		t.Fatal(err)
	}
	// Create a mock repository
	repository := NewMongoRepository(client)

	// Get the mock document
	got, err := repository.DeleteDriver(driver)
	if err != nil {
		t.Fatal(err)
	}

	// Check if the mock document is equal to the mock driver
	if !reflect.DeepEqual(got, int64(1)) {
		t.Fatalf("got %v, want %v", got, 1)
	}
}

func TestMongoRepository_UpdateDriver(t *testing.T) {
	uri := "mongodb://localhost:27017/?maxPoolSize=20&w=majority"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		t.Fatal(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		t.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		t.Fatal(err)
	}
	// Create a mock collection
	collection := client.Database("bitaksi").Collection("driver")
	collection.Drop(context.TODO())
	// Create a mock driver
	driver := &model.Driver{
		ID:       1,
		Location: model.GeoJson{Type: "Point", Coordinates: []float64{1, 1}},
	}
	updatedDriver := &model.Driver{
		ID:       1,
		Location: model.GeoJson{Type: "Point", Coordinates: []float64{1, 2}},
	}
	// Insert the mock document
	_, err = collection.InsertOne(context.TODO(), driver)
	if err != nil {
		t.Fatal(err)
	}
	// Create a mock repository
	repository := NewMongoRepository(client)

	d, err := repository.UpdateDriver(updatedDriver)
	if err != nil {
		t.Fatal(err)
	}

	// Check if the mock document is equal to the mock driver
	if !reflect.DeepEqual(d, updatedDriver) {
		t.Fatalf("got %v, want %v", d, 1)
	}
}

func TestMongoRepository_GetAllDrivers(t *testing.T) {
	uri := "mongodb://localhost:27017/?maxPoolSize=20&w=majority"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		t.Fatal(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		t.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		t.Fatal(err)
	}
	// Create a mock collection
	collection := client.Database("bitaksi").Collection("driver")
	collection.Drop(context.TODO())
	// Create a mock driver
	driver1 := &model.Driver{
		ID:       1,
		Location: model.GeoJson{Type: "Point", Coordinates: []float64{1, 1}},
	}
	driver2 := &model.Driver{
		ID:       2,
		Location: model.GeoJson{Type: "Point", Coordinates: []float64{1, 2}},
	}
	// Insert the mock document
	_, err = collection.InsertOne(context.TODO(), driver1)
	if err != nil {
		t.Fatal(err)
	}
	// Insert the mock document
	_, err = collection.InsertOne(context.TODO(), driver2)
	if err != nil {
		t.Fatal(err)
	}

	mockDrivers := []*model.Driver{driver1, driver2}
	// Create a mock repository
	repository := NewMongoRepository(client)
	drivers, err := repository.GetAllDrivers()
	if err != nil {
		t.Fatal(err)
	}

	// Check if the mock document is equal to the mock driver
	if !reflect.DeepEqual(drivers, mockDrivers) {
		t.Fatalf("got %v, want %v", drivers, mockDrivers)
	}

}
