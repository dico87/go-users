package service

import (
	"errors"
	"github.com/dico87/users/model"
	"github.com/dico87/users/repository"
	"gorm.io/gorm"
)

type UserService interface {
	Create(user model.User) (model.User, error)
	Update(id uint, user model.User) (model.User, error)
	FindById(id uint) (model.User, error)
	FindByDocument(documentTypeId uint, document string) (model.User, error)
}

type UserServiceImpl struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return UserServiceImpl{repository: repository}
}

func (service UserServiceImpl) Create(user model.User) (model.User, error) {
	_, err := service.repository.FindByDocument(user.DocumentType.ID, user.Document)

	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return model.User{}, err
		}

		createdUser, err := service.repository.Create(user)

		if err != nil {
			return model.User{}, err
		}

		return createdUser, nil
	}

	return model.User{}, errors.New("user already exists")
}

func(service UserServiceImpl) Update(id uint, user model.User) (model.User, error) {
	_, err := service.repository.FindById(id)

	if err != nil {
		return model.User{}, err
	}

	user.ID = id
	updatedUser, err := service.repository.Update(user)

	if err != nil {
		return model.User{}, err
	}

	return updatedUser, nil

}

func (service UserServiceImpl) FindById(id uint) (model.User, error) {
	return service.repository.FindById(id)
}

func (service UserServiceImpl) FindByDocument(documentTypeID uint, document string) (model.User, error) {
	return service.repository.FindByDocument(documentTypeID, document)
}