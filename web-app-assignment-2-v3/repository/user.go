package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
	"errors"
)

type UserRepository interface {
	GetUserByEmail(email string) (model.User, error)
	CreateUser(user model.User) (model.User, error)
	GetUserTaskCategory() ([]model.UserTaskCategory, error)
}

type userRepository struct {
	filebasedDb *filebased.Data
}

func NewUserRepo(filebasedDb *filebased.Data) *userRepository {
	return &userRepository{filebasedDb}
}

var ErrNotFound = errors.New("user not found")

func (r *userRepository) GetUserByEmail(email string) (model.User, error) {
	// Gunakan fungsi GetUserByEmail dari filebased.Data
	user, err := r.filebasedDb.GetUserByEmail(email)
	if err != nil {
		// Jika user tidak ditemukan, kembalikan error ErrNotFound
		if user == (model.User{}) {
			return model.User{}, ErrNotFound
		}
		return model.User{}, err
	}
	return user, nil
}

func (r *userRepository) CreateUser(user model.User) (model.User, error) {
	// Gunakan fungsi CreateUser dari filebased.Data
	createdUser, err := r.filebasedDb.CreateUser(user)
	if err != nil {
		return model.User{}, err
	}
	return createdUser, nil
}

func (r *userRepository) GetUserTaskCategory() ([]model.UserTaskCategory, error) {
	// Gunakan fungsi GetUserTaskCategory dari filebased.Data
	userTaskCategories, err := r.filebasedDb.GetUserTaskCategory()
	if err != nil {
		return nil, err
	}
	return userTaskCategories, nil
}