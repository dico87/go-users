package domain

type Repository interface {
	Create(user User) (User, error)
	Update(user User) (User, error)
	FindById(id uint) (User, error)
	FindByDocument(documentTypeID uint, document string) (User, error)
}
