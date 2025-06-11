package user

import (
	"github.com/pnaskardev/pubjudge/gateway/api/presenter"
	"github.com/pnaskardev/pubjudge/gateway/pkg/entities"
)

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	InsertUser(User *entities.User) (*entities.User, error)
	FetchUsers() (*[]presenter.User, error)
	UpdateUser(User *entities.User) (*entities.User, error)
	DeleteUsers(ID string) error
	GetUsersWithSubmissions()
}

type service struct {
	repository Repository
}

// NewService is used to create a single instance of the service
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// InsertSubmit is a service layer that helps insert Submit in SubmitShop
func (s *service) InsertUser(User *entities.User) (*entities.User, error) {
	return s.repository.CreateUser(User)
}

// FetchSubmits is a service layer that helps fetch all Submits in SubmitShop
func (s *service) FetchUsers() (*[]presenter.User, error) {
	return s.repository.ReadUser()
}

// UpdateSubmit is a service layer that helps update Submits in SubmitShop
func (s *service) UpdateUser(User *entities.User) (*entities.User, error) {
	return s.repository.UpdateUser(User)
}

// RemoveSubmit is a service layer that helps remove Submits from SubmitShop
func (s *service) DeleteUsers(ID string) error {
	return s.repository.DeleteUser(ID)
}

func (s *service) GetUsersWithSubmissions() {
	s.repository.GetUsersWithSubmissions()
}
