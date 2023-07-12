package service

import (
	publisherDto "golang/src/publisher/dto"
	publisherEntity "golang/src/publisher/entity"
	publisherRepository "golang/src/publisher/repository"
)

type PublisherService interface {
	FindAll() []publisherEntity.Publisher
	FindById(id int) (*publisherEntity.Publisher, error)
	Create(publisherDto publisherDto.CreatePublisherRequest) (*publisherEntity.Publisher, error)
}

type PublisherServiceImpl struct {
	publisherRepository publisherRepository.PublisherRepository
}

// Create implements PublisherService.
func (publisherService *PublisherServiceImpl) Create(publisherDto publisherDto.CreatePublisherRequest) (*publisherEntity.Publisher, error) {
	var publisher publisherEntity.Publisher

	publisher.Name = publisherDto.Name
	publisher.Address = publisherDto.Address
	publisher.Phone = publisherDto.Phone

	// create new user to database

	dataPublisher, err := publisherService.publisherRepository.Create(publisher)

	if err != nil {
		return nil, err
	}

	return dataPublisher, nil
}

// FindAll implements PublisherService.
func (publisherService *PublisherServiceImpl) FindAll() []publisherEntity.Publisher {
	return publisherService.publisherRepository.FindAll()
}

// FindById implements PublisherService.
func (publisherService *PublisherServiceImpl) FindById(id int) (*publisherEntity.Publisher, error) {
	return publisherService.publisherRepository.FindById(id)
}

func NewPublisherService(publisherRepository publisherRepository.PublisherRepository) PublisherService {
	return &PublisherServiceImpl{publisherRepository}
}
