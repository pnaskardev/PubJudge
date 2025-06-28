package submit

import (
	"context"
	"encoding/json"

	"github.com/pnaskardev/pubjudge/gateway/pkg/entities"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository interface {
	CreateSubmit(Submission *entities.CreateSubmissionInput, userId primitive.ObjectID) (*entities.Submission, error)
}

type repository struct {
	Redis      *redis.Client
	Collection *mongo.Collection
}

func NewRepo(collection *mongo.Collection, redis_client *redis.Client) Repository {
	return &repository{Collection: collection, Redis: redis_client}
}

func (r *repository) CreateSubmit(Submission *entities.CreateSubmissionInput, userId primitive.ObjectID) (*entities.Submission, error) {

	// convertts the object to JSON
	payload, parsingErr := json.Marshal(Submission)
	if parsingErr != nil {
		panic(parsingErr)
	}

	// publish to redis
	// submission is the key
	if err := r.Redis.Publish(context.Background(), "submission", string(payload)).Err(); err != nil {
		panic(err)
	}

	submission_instance := entities.Submission{
		ID:       primitive.NewObjectID(),
		UserID:   userId,
		Code:     Submission.Code,
		Language: Submission.Language,
	}

	_, err := r.Collection.InsertOne(context.Background(), submission_instance)

	if err != nil {
		return nil, err
	}

	return &submission_instance, nil

}
