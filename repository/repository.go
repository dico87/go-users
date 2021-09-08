package repository

import "github.com/dico87/users/model"

type UserRepository interface {
	Create(user model.User) (model.User, error)
	Update(user model.User) (model.User, error)
	FindById(id uint) (model.User, error)
	FindByDocument(documentTypeID uint, document string) (model.User, error)
}
