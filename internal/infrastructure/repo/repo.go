package repo

import (
	"context"
	"fmt"
	"log/slog"
	"mongodb_play/internal/domain/entities"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	fieldBrand       = "brand"
	fieldModel       = "model"
	fieldEngine      = "engine"
	fieldEnginePower = "engine_power"
	fieldNumberPlate = "number_plate"
)

type Filter map[string]any

type Repo struct {
	garage     *mongo.Database
	collection *mongo.Collection
}

func NewRepo(mongoDB *mongo.Client, dbName string) *Repo {
	db := mongoDB.Database(dbName)
	return &Repo{
		garage: db,
	}
}

func (r *Repo) SetCollection(ctx context.Context, colName string) error {
	colls, err := r.garage.ListCollectionNames(ctx, bson.M{"name": colName})
	if err != nil {
		return errors.Wrap(err, "failed to get the list of collections")
	}
	if len(colls) == 1 {
		r.collection = r.garage.Collection(colName)
		return nil
	}
	slog.With(slog.String("coll name", colName)).Error("collection is not found")
	return errors.New("collection is not found")
}

func (r *Repo) FindCarByNumberPlate(ctx context.Context, nPlate string) (entities.Car, error) {
	if r.collection == nil {
		return entities.Car{}, errors.New("collection is not set")
	}
	// filter := bson.D{{fieldNumberPlate, nPlate}}
	filter := bson.M{fieldNumberPlate: nPlate}
	res := entities.Car{}
	err := r.collection.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r *Repo) AddCar(ctx context.Context, c entities.Car) (string, error) {
	if r.collection == nil {
		return "", errors.New("collection is not set")
	}
	snap := "repo.AddCar"
	_, err := r.FindCarByNumberPlate(ctx, c.NumberPlate)
	if errors.Is(err, mongo.ErrNoDocuments) {
		res, e := r.collection.InsertOne(ctx, c)
		if e != nil {
			return "", errors.Wrap(err, "error inserting")
		}
		v, ok := res.InsertedID.(primitive.ObjectID)
		if !ok {
			return "", errors.Wrap(err, "error getting _id")
		}
		return v.String(), nil
	} else if err != nil {
		slog.With(slog.String("snap", snap)).ErrorContext(ctx, err.Error())
	}
	return "", fmt.Errorf("can't add, the car with \"%s\" number plate already exists", c.NumberPlate)
}

func (r *Repo) DeleteCar(ctx context.Context, nPlate string) error {
	if r.collection == nil {
		return errors.New("collection is not set")
	}
	filter := bson.M{fieldNumberPlate: nPlate}
	res, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return errors.Wrapf(err, "error deleting car with number plate: %s", nPlate)
	}
	slog.With(slog.Int64("count", res.DeletedCount)).InfoContext(ctx, "the number of documents deleted")
	return nil
}

func (r *Repo) ModifyCar(ctx context.Context, c entities.Car) error {
	if r.collection == nil {
		return errors.New("collection is not set")
	}
	opts := options.Replace().SetUpsert(false)
	filter := Filter{fieldNumberPlate: c.NumberPlate}
	replacement := bson.D{
		primitive.E{Key: fieldBrand, Value: c.Brand},
		primitive.E{Key: fieldModel, Value: c.Model},
		primitive.E{Key: fieldEngine, Value: c.Engine},
		primitive.E{Key: fieldNumberPlate, Value: c.NumberPlate},
		primitive.E{Key: fieldEnginePower, Value: c.EnginePower},
	}
	res, err := r.collection.ReplaceOne(ctx, filter, replacement, opts)
	if err != nil {
		return errors.Wrap(err, "error modifying")
	}
	if res.MatchedCount != 0 {
		slog.Info("matched and replaced an existing document")
		return nil
	}
	if res.UpsertedCount != 0 {
		slog.With(slog.Any("id", res.UpsertedID)).InfoContext(ctx, "inserted a new document")
	}
	return nil
}
