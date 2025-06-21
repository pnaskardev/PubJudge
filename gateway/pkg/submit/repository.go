package submit

import (
	"context"

	"github.com/pnaskardev/pubjudge/gateway/pkg/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository interface {
	CreateSubmit(Submission *entities.CreateSubmissionInput, userId primitive.ObjectID) (*entities.Submission, error)
}

type repository struct {
	Collection *mongo.Collection
}

func NewRepo(collection *mongo.Collection) Repository {
	return &repository{Collection: collection}
}

func (r *repository) CreateSubmit(Submission *entities.CreateSubmissionInput, userId primitive.ObjectID) (*entities.Submission, error) {

	submission_instance := entities.Submission{
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
