package repository

import (
	publisherEntity "golang/src/publisher/entity"

	"gorm.io/gorm"

)

type PublisherRepository interface {
	FindAll() []publisherEntity.Publisher
	FindById(id int) (*publisherEntity.Publisher, error)
	Create(publisher publisherEntity.Publisher) (*publisherEntity.Publisher, error)
}

type PublisherRepositoryImpl struct {
	db *gorm.DB
}

// Create implements PublisherRepository.
func (publisherRepository *PublisherRepositoryImpl) Create(publisher publisherEntity.Publisher) (*publisherEntity.Publisher, error) {
	dataPublisher := publisherRepository.db.Create(&publisher)

	if dataPublisher.Error != nil {
		return nil, dataPublisher.Error
	}

	return &publisher, nil
}

// FindAll implements PublisherRepository.
func (publisherRepository *PublisherRepositoryImpl) FindAll() []publisherEntity.Publisher {
	var publishers []publisherEntity.Publisher

	publisherRepository.db.Find(&publishers)

	return publishers
}

// FindById implements PublisherRepository.
func (publisherRepository *PublisherRepositoryImpl) FindById(id int) (*publisherEntity.Publisher, error) {
	var publisher publisherEntity.Publisher

	dataPublisher := publisherRepository.db.First(&publisher, id)

	if dataPublisher.Error != nil {
		return nil, dataPublisher.Error
	}

	return &publisher, nil
}

func NewPublisherRepository(db *gorm.DB) PublisherRepository {
	return &PublisherRepositoryImpl{db}
}
