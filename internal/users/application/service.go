package application

import (
	"errors"
	"github.com/dico87/users/internal/common"
	"github.com/dico87/users/internal/users/domain"
)

type Service struct {
	repository domain.Repository
}

func New(repository domain.Repository) Service {
	return Service{repository: repository}
}

func (service Service) Create(user domain.User) (domain.User, error) {
	_, err := service.repository.FindByDocument(user.DocumentType.ID, user.Document)

	if err != nil {
		if err != common.ErrNotFoundRecord {
			return domain.User{}, err
		}

		createdUser, err := service.repository.Create(user)

		if err != nil {
			return domain.User{}, err
		}

		return createdUser, nil
	}

	return domain.User{}, errors.New("user already exists")
}

func (service Service) Update(id uint, user domain.User) (domain.User, error) {
	_, err := service.repository.FindById(id)

	if err != nil {
		return domain.User{}, err
	}

	user.ID = id
	updatedUser, err := service.repository.Update(user)

	if err != nil {
		return domain.User{}, err
	}

	return updatedUser, nil

}

func (service Service) FindById(id uint) (domain.User, error) {
	foundUser, err := service.repository.FindById(id)
	if err != nil {
		return domain.User{}, err
	}

	return foundUser, nil
}

func (service Service) FindByDocument(documentTypeID uint, document string) (domain.User, error) {
	foundUser, err := service.repository.FindByDocument(documentTypeID, document)
	if err != nil {
		return domain.User{}, err
	}

	return foundUser, nil
}
