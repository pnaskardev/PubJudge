package submit

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/pnaskardev/pubjudge/gateway/pkg/entities"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const (
	stream = "submission"
	group  = "submission_group"
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
	var values map[string]interface{}
	if err := json.Unmarshal(payload, &values); err != nil {
		panic(err)
	}
	// _, streamError := r.Redis.XAdd(context.Background(), &redis.XAddArgs{Stream: "submission", Values: values}).Result()

	streamError := r.Redis.XGroupCreateMkStream(context.Background(), stream, group, "$").Err()

	if streamError != nil && !strings.Contains(streamError.Error(), "BUSYGROUP") {
		panic(streamError)
	}
	// I am already creating a stream and a group
	_, addErr := r.Redis.XAdd(context.Background(), &redis.XAddArgs{Stream: stream, NoMkStream: true, Approx: true, Values: values}).Result()

	if addErr != nil {
		panic(addErr)
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
