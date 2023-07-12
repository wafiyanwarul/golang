package author

import (
	authorDto "golang/src/author/dto"
	authorEntity "golang/src/author/entity"
	authorRepository "golang/src/author/repository"
)

type AuthorService interface {
	FindAll() []authorEntity.Author
	FindById(id int) (*authorEntity.Author, error)
	Create(authorDto authorDto.CreateAuthorRequest) (*authorEntity.Author, error)
	Update(id int, authorDto authorDto.CreateAuthorRequest) (*authorEntity.Author, error)
	Delete(id int) error
}

type AuthorServiceImpl struct {
	authorRepository authorRepository.AuthorRepository
}

// Create implements AuthorService.
func (authorService *AuthorServiceImpl) Create(authorDto authorDto.CreateAuthorRequest) (*authorEntity.Author, error) {
	var author authorEntity.Author

	author.Name = authorDto.Name

	// create new user to database

	dataAuthor, err := authorService.authorRepository.Create(author)

	if err != nil {
		return nil, err
	}

	return dataAuthor, nil
}

// Delete implements AuthorService.
func (authorService *AuthorServiceImpl) Delete(id int) error {
	panic("unimplemented")
}

// FindAll implements AuthorService.
func (authorService *AuthorServiceImpl) FindAll() []authorEntity.Author {
	return authorService.authorRepository.FindAll()
}

// FindById implements AuthorService.
func (authorService *AuthorServiceImpl) FindById(id int) (*authorEntity.Author, error) {
	return authorService.authorRepository.FindById(id)
}

// Update implements AuthorService.
func (authorService *AuthorServiceImpl) Update(id int, authorDto authorDto.CreateAuthorRequest) (*authorEntity.Author, error) {
	panic("unimplemented")
}

func NewAuthorService(authorRepository authorRepository.AuthorRepository) AuthorService {
	return &AuthorServiceImpl{authorRepository}
}
