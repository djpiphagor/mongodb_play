package repo

import (
	"context"
	"log/slog"
	"mongodb_play/internal/domain/entities"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repo struct {
	collection *mongo.Collection
}

func NewRepo(mongoDB *mongo.Client, db string, col string) *Repo {
	collection := mongoDB.Database(db).Collection(col)
	return &Repo{
		collection: collection,
	}
}

func (r *Repo) AddCar(ctx context.Context, c entities.Car) (string, error) {
	_, err := r.FindCarByNumberPlate(ctx, c.NumberPlate)
	if errors.As(err, mongo.ErrNoDocuments) {
		return "", errors.Wrap(err, "can't add, the car with this number plate already exists")
	}
	res, err := r.collection.InsertOne(ctx, c)
	if err != nil {
		return "", errors.Wrap(err, "error inserting")
	}
	v, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.Wrap(err, "error getting _id")
	}
	return v.String(), nil
}

func (r *Repo) DeleteCar(ctx context.Context, nPlate string) error {
	filter := bson.D{{"number_plate", nPlate}}
	res, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return errors.Wrapf(err, "error deleting car with number plate: %s", nPlate)
	}
	slog.With(slog.Int64("count", res.DeletedCount)).Info("the number of documents deleted")
	return nil
}

func (r *Repo) ModifyCar(ctx context.Context, c entities.Car) error {
	opts := options.Replace().SetUpsert(true)
	filter := bson.D{{"number_plate", c.NumberPlate}}
	replacement := bson.D{
		{"brand", c.Brand},
		{"model", c.Model},
		{"engine", c.Engine},
		{"number_plate", c.NumberPlate},
		{"engine_power", c.EnginePower},
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
		slog.With(slog.Any("id", res.UpsertedID)).Info("inserted a new document")
	}
	return nil
}

func (r *Repo) FindCarByNumberPlate(ctx context.Context, nPlate string) (entities.Car, error) {
	filter := bson.D{{"number_plate", nPlate}}
	res := entities.Car{}
	err := r.collection.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return res, errors.Wrapf(err, "there is no car with number plate: %s", nPlate)
		}
		return res, errors.Wrapf(err, "unexpected error while trying to find car with number plate: %s", nPlate)
	}
	return res, nil
}
