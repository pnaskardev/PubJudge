package user

import (
	"context"
	"fmt"
	"time"

	"github.com/pnaskardev/pubjudge/gateway/api/presenter"
	"golang.org/x/crypto/bcrypt"

	"github.com/pnaskardev/pubjudge/gateway/pkg/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// Repository interface allows us to access the CRUD Operations in mongo here.
type Repository interface {
	FetchUser(Credentials *entities.LoginInput) (*entities.User, error)
	CreateUser(User *entities.User) (*entities.User, error)
	ReadUser() (*[]presenter.User, error)
	UpdateUser(User *entities.User) (*entities.User, error)
	DeleteUser(ID string) error
}
type repository struct {
	Collection *mongo.Collection
}

// NewRepo is the single instance repo that is being created.
func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}

func (r *repository) FetchUser(Credentials *entities.LoginInput) (*entities.User, error) {

	var user entities.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := r.Collection.FindOne(ctx, bson.M{"username": Credentials.Username}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("User not found")
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Credentials.Password))
	if err != nil {
		return nil, fmt.Errorf("Invalid credentials")
	}
	return &user, nil
}

// CreateUser is a mongo repository that helps to create Submits
func (r *repository) CreateUser(User *entities.User) (*entities.User, error) {
	User.ID = primitive.NewObjectID()
	User.CreatedAt = time.Now()
	User.UpdatedAt = time.Now()
	_, err := r.Collection.InsertOne(context.Background(), User)
	if err != nil {
		return nil, err
	}
	return User, nil
}

// ReadUser is a mongo repository that helps to fetch Submits
func (r *repository) ReadUser() (*[]presenter.User, error) {
	var Submits []presenter.User
	cursor, err := r.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var User presenter.User
		_ = cursor.Decode(&User)
		Submits = append(Submits, User)
	}
	return &Submits, nil
}

// UpdateUser is a mongo repository that helps to update Submits
func (r *repository) UpdateUser(User *entities.User) (*entities.User, error) {
	User.UpdatedAt = time.Now()
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": User.ID}, bson.M{"$set": User})
	if err != nil {
		return nil, err
	}
	return User, nil
}

// DeleteUser is a mongo repository that helps to delete Submits
func (r *repository) DeleteUser(ID string) error {
	SubmitID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": SubmitID})
	if err != nil {
		return err
	}
	return nil
}
