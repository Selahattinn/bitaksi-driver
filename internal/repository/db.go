package repository

import (
	"context"

	"github.com/Selahattinn/bitaksi-driver/internal/repository/driver"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoRepository defines the Mongo implementation of Repository interface
type MongoRepository struct {
	cfg        *MongoConfig
	client     *mongo.Client
	driverRepo driver.Repository
}

type MongoConfig struct {
	Addr     string `yaml:"addr"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// dbConn opens connection with Mongo driver
func dbConn(cfg *MongoConfig) (*mongo.Client, error) {

	//TODO : Add authentication
	uri := "mongodb://" + cfg.Addr + ":" + cfg.Port + "/?maxPoolSize=20&w=majority"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}
	// Return db object to be used by other functions
	return client, nil
}

// NewMongoRepository creates a new Mongo Repository
func NewMongoRepository(cfg *MongoConfig) (*MongoRepository, error) {
	client, err := dbConn(cfg)
	if err != nil {
		return nil, err
	}
	driverRepo := driver.NewMongoRepository(client)
	return &MongoRepository{
		cfg:        cfg,
		client:     client,
		driverRepo: driverRepo,
	}, nil
}

// Shutdown closes the database connection
func (r *MongoRepository) Shutdown() {
	err := r.client.Disconnect(context.TODO())
	if err != nil {
		logrus.Fatal(err)
	}
}

// GetDriverRepository gets the user repository
func (r *MongoRepository) GetDriverRepository() driver.Repository {
	return r.driverRepo
}
