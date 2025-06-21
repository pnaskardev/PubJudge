package submit

import (
	"github.com/pnaskardev/pubjudge/gateway/pkg/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	CreateSubmit(Submission *entities.CreateSubmissionInput, userId primitive.ObjectID) (*entities.Submission, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) CreateSubmit(Submission *entities.CreateSubmissionInput, userId primitive.ObjectID) (*entities.Submission, error) {
	return s.repository.CreateSubmit(Submission, userId)
}
