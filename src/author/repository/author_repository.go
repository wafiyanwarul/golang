package author

import (
	authorEntity "golang/src/author/entity"

	"gorm.io/gorm"
)

type AuthorRepository interface {
	FindAll() []authorEntity.Author
	FindById(id int) (*authorEntity.Author, error)
	Create(author authorEntity.Author) (*authorEntity.Author, error)
	Update(author authorEntity.Author) (*authorEntity.Author, error)
	Delete(id int) error
}

type AuthorRepositoryImpl struct {
	db *gorm.DB
}

// Create implements AuthorRepository.
func (authorRepository *AuthorRepositoryImpl) Create(author authorEntity.Author) (*authorEntity.Author, error) {
	// create new user
	dataAuthor := authorRepository.db.Create(&author)

	// return error if any
	if dataAuthor.Error != nil {
		return nil, dataAuthor.Error
	}

	// return user
	return &author, nil
}

// Delete implements AuthorRepository.
func (authorRepository *AuthorRepositoryImpl) Delete(id int) error {
	panic("unimplemented")
}

// FindAll implements AuthorRepository.
func (authorRepository *AuthorRepositoryImpl) FindAll() []authorEntity.Author {
	var authors []authorEntity.Author

	authorRepository.db.Find(&authors)

	return authors
}

// FindById implements AuthorRepository.
func (authorRepository *AuthorRepositoryImpl) FindById(id int) (*authorEntity.Author, error) {
	var author authorEntity.Author

	dataAuthor := authorRepository.db.First(&author, id)

	if dataAuthor.Error != nil {
		return nil, dataAuthor.Error
	}

	return &author, nil
}

// Update implements AuthorRepository.
func (authorRepository *AuthorRepositoryImpl) Update(author authorEntity.Author) (*authorEntity.Author, error) {
	panic("unimplemented")
}

func NewAuthorRepository(db *gorm.DB) AuthorRepository {
	return &AuthorRepositoryImpl{db}
}
