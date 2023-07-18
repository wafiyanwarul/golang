package rental

import (
	rentalEntity "golang/src/rental/entity"

	"gorm.io/gorm"

)

type RentalRepository interface {
	FindAll() []rentalEntity.Rental
	FindById(id int) (*rentalEntity.Rental, error)
	Create(rental rentalEntity.Rental) (*rentalEntity.Rental, error)
	Update(rental rentalEntity.Rental) (*rentalEntity.Rental, error)
}

type RentalRepositoryImpl struct {
	db *gorm.DB
}

// Create implements RentalRepository.
func (rentalRepository *RentalRepositoryImpl) Create(rental rentalEntity.Rental) (*rentalEntity.Rental, error) {
	dataRental := rentalRepository.db.Create(&rental)

	if dataRental.Error != nil {
		return nil, dataRental.Error
	}

	return &rental, nil
}

// FindAll implements RentalRepository.
func (rentalRepository *RentalRepositoryImpl) FindAll() []rentalEntity.Rental {
	var rentals []rentalEntity.Rental

	rentalRepository.db.Find(&rentals)

	return rentals
}

// FindById implements RentalRepository.
func (rentalRepository *RentalRepositoryImpl) FindById(id int) (*rentalEntity.Rental, error) {
	var rental rentalEntity.Rental

	dataRental := rentalRepository.db.Preload("User").Preload("Book").First(&rental, id)

	if dataRental.Error != nil {
		return nil, dataRental.Error
	}

	return &rental, nil
}

// Update implements RentalRepository.
func (rentalRepository *RentalRepositoryImpl) Update(rental rentalEntity.Rental) (*rentalEntity.Rental, error) {
	updateRental := rentalRepository.db.Save(&rental)

	if updateRental.Error != nil {
		return nil, updateRental.Error
	}

	return &rental, nil
}

func NewRentalRepository(db *gorm.DB) RentalRepository {
	return &RentalRepositoryImpl{db}
}