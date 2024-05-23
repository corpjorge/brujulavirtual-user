package repositories

import (
	"brujulavirtual-auth/src/register/domain/models"
	"brujulavirtual-auth/src/register/domain/ports"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type Mongo struct {
	collection *mongo.Collection
}

func Register() ports.Repository {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("brujulavirtual").Collection("users")
	return &Mongo{
		collection: collection,
	}
}

func (r *Mongo) Save(auth models.Register) (models.Register, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"user": auth.User}

	log.Default().Printf("MONGO DATA: %v\n", filter)

	var result models.Register
	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Default().Println("No documents found for the given filter")
			return models.Register{}, errors.New("no user found")
		}
		log.Default().Printf("Error finding document: %v\n", err)
		return models.Register{}, err
	}

	return result, nil
}
