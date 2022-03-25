package repository

import (
	"context"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Selahattinn/bitaksi-driver/internal/model"
	"github.com/Selahattinn/bitaksi-driver/internal/repository/driver"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
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
	DataPath string `yaml:"dataPath"`
}

// dbConn opens connection with Mongo driver
func dbConn(cfg *MongoConfig) (*mongo.Client, error) {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://" + cfg.Username + ":" + cfg.Password + "@" + cfg.Addr + "/myFirstDatabase?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}
	return client, nil
}

// NewMongoRepository creates a new Mongo Repository
func NewMongoRepository(cfg *MongoConfig) (*MongoRepository, error) {
	client, err := dbConn(cfg)
	if err != nil {
		return nil, err
	}

	// initlize with data
	if cfg.DataPath != "" {
		err := initalizeWithData(cfg.DataPath, client)
		if err != nil {
			return nil, err
		}

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

func initalizeWithData(dataPath string, client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := client.Database("bitaksi").Drop(ctx)
	if err != nil {
		logrus.WithError(err).Error("failed to drop database")
		return err
	}
	err = client.Database("bitaksi").CreateCollection(ctx, "driver")
	if err != nil {
		logrus.WithError(err).Error("failed to create collection")
		return err
	}

	// create index 2dsphere
	index := mongo.IndexModel{
		Keys: bson.M{
			"location": "2dsphere",
		},
	}
	_, err = client.Database("bitaksi").Collection("driver").Indexes().CreateOne(context.TODO(), index)
	if err != nil {
		return err
	}
	drivers, err := readCSV(dataPath)
	if err != nil {
		logrus.WithError(err).Error("failed to read csv")
		return err
	}
	for id, driver := range drivers {

		// convert ID to int64
		driver.ID = int64(id)

		// create driver object
		d := &model.Driver{
			ID: driver.ID,
			Location: model.GeoJson{
				Type:        "Point",
				Coordinates: []float64{driver.Location.Coordinates[0], driver.Location.Coordinates[1]},
			},
		}

		// insertOne To DB
		_, err := client.Database("bitaksi").Collection("driver").InsertOne(ctx, d)
		if err != nil {
			logrus.WithError(err).Error("failed to insert driver when intilizing")
		}

	}
	return nil

}

func readCSV(path string) ([]*model.Driver, error) {

	f, err := os.Open(path)
	if err != nil {
		logrus.WithError(err).Error("failed to open initial data file")
		return nil, err
	}
	defer f.Close()

	drivers := []*model.Driver{}
	csvReader := csv.NewReader(f)
	csvReader.Read() // skip header
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			logrus.WithError(err).Error("failed to read csv")
		}
		if len(rec) != 2 {
			logrus.WithError(err).Error("Invalid csv format")
			return nil, err
		}
		// string to float64

		lat, err := strconv.ParseFloat(rec[0], 64)
		if err != nil {

			logrus.WithError(err).Error("failed to parse lat")
			return nil, err
		}
		long, err := strconv.ParseFloat(rec[1], 64)
		if err != nil {
			logrus.WithError(err).Error("failed to parse Long")
			return nil, err
		}
		driver := &model.Driver{
			Location: model.GeoJson{
				Type:        "Point",
				Coordinates: []float64{lat, long},
			},
		}
		err = driver.Validate()
		if err != nil {
			logrus.WithError(err).Error("failed to validate driver")
			continue
		}
		drivers = append(drivers, driver)

	}
	return drivers, nil
}
