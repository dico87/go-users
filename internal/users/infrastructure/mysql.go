package infrastructure

import (
	"errors"
	"github.com/dico87/users/internal/common"
	"github.com/dico87/users/internal/users/domain"
	"gorm.io/gorm"
)

func NewMySqlRepository(db *gorm.DB) domain.Repository {
	return MySqlRepository{
		db: db,
	}
}

type MySqlRepository struct {
	db *gorm.DB
}

func (repo MySqlRepository) Create(user domain.User) (domain.User, error) {
	db := repo.db.Create(&user)
	return user, db.Error
}

func (repo MySqlRepository) Update(user domain.User) (domain.User, error) {
	db := repo.db.Save(&user)
	return user, db.Error
}

func (repo MySqlRepository) FindById(id uint) (domain.User, error) {
	user := domain.User{}
	db := repo.db.Preload("DocumentType").Take(&user, id)

	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return domain.User{}, common.ErrNotFoundRecord
	}
	return user, db.Error
}

func (repo MySqlRepository) FindByDocument(documentTypeID uint, document string) (domain.User, error) {
	user := domain.User{}
	db := repo.db.Preload("DocumentType").Where(&domain.User{DocumentTypeID: documentTypeID, Document: document}).First(&user)
	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return domain.User{}, common.ErrNotFoundRecord
	}
	return user, db.Error
}
