package repository

import (
	"github.com/dico87/users/model"
	"gorm.io/gorm"
)

func NewMySqlRepository(db *gorm.DB) UserRepository {
	return MySqlRepository{
		db: db,
	}
}

type MySqlRepository struct {
	db *gorm.DB
}


func (repo MySqlRepository) Create(user model.User) (model.User, error) {
	db := repo.db.Create(&user)
	return user, db.Error
}

func (repo MySqlRepository) Update(user model.User) (model.User, error) {
	db := repo.db.Save(&user)
	return user, db.Error
}

func (repo MySqlRepository) FindById(id uint) (model.User, error) {
	user := model.User{}
	db := repo.db.Take(&user, id)
	return user, db.Error
}

func (repo MySqlRepository) FindByDocument(documentTypeID uint, document string) (model.User, error) {
	user := model.User{}
	db := repo.db.Where(&model.User{DocumentTypeID: documentTypeID, Document: document}).First(&user)
	return user, db.Error
}