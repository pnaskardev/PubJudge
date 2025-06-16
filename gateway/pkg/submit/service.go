package submit

import "github.com/pnaskardev/pubjudge/gateway/pkg/entities"

type Service interface {
	CreateSubmit(Submission *entities.Submission) (*entities.Submission, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) CreateSubmit(Submission *entities.Submission) (*entities.Submission, error) {
	return s.repository.CreateSubmit(Submission)
}
