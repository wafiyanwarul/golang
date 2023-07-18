package rental

import (
	"database/sql"
	fileupload "golang/pkg/fileupload/cloudinary"
	rentalDto "golang/src/rental/dto"
	rentalEntity "golang/src/rental/entity"
	rentalRepository "golang/src/rental/repository"
	"time"
)

type RentalService interface {
	FindAll() []rentalEntity.Rental
	FindById(id int) (*rentalEntity.Rental, error)
	Create(rentalRequest rentalDto.CreateRentalRequest) (*rentalEntity.Rental, error)
	Update(id int, rentalRequest rentalDto.UpdateRentalRequest) (*rentalEntity.Rental, error)
	Return(id int, rentalRequest rentalDto.ReturnRentalRequest) (*rentalEntity.Rental, error)
}

type RentalServiceImpl struct {
	rentalRepository rentalRepository.RentalRepository
	fileupload       fileupload.FileUpload
}

// Return implements RentalService.
func (rentalService *RentalServiceImpl) Return(id int, rentalRequest rentalDto.ReturnRentalRequest) (*rentalEntity.Rental, error) {
	rental, err := rentalService.rentalRepository.FindById(id)

	if err != nil {
		return nil, err
	}

	rental.Status = "returned"
	rental.ReturnDate = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	if rentalRequest.Image != nil {

		image, err := rentalService.fileupload.Upload(*rentalRequest.Image)

		if err != nil {
			return nil, err
		}

		rental.Image = image
	}

	updatedRental, err := rentalService.rentalRepository.Update(*rental)

	if err != nil {
		return nil, err
	}

	return updatedRental, nil
}

// Create implements RentalService.
func (rentalService *RentalServiceImpl) Create(rentalRequest rentalDto.CreateRentalRequest) (*rentalEntity.Rental, error) {
	newRental := rentalEntity.Rental{
		Status: "borrowed",
		RentDate: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		ReturnDate: sql.NullTime{},
		BookID:     rentalRequest.BookID,
		UserID:     rentalRequest.UserID,
	}

	rental, err := rentalService.rentalRepository.Create(newRental)

	if err != nil {
		return nil, err
	}

	return rental, nil
}

// FindAll implements RentalService.
func (rentalService *RentalServiceImpl) FindAll() []rentalEntity.Rental {
	return rentalService.rentalRepository.FindAll()
}

// FindById implements RentalService.
func (rentalService *RentalServiceImpl) FindById(id int) (*rentalEntity.Rental, error) {
	return rentalService.rentalRepository.FindById(id)
}

// Update implements RentalService.
func (rentalService *RentalServiceImpl) Update(id int, rentalRequest rentalDto.UpdateRentalRequest) (*rentalEntity.Rental, error) {
	rental, err := rentalService.rentalRepository.FindById(id)

	if err != nil {
		return nil, err
	}

	rental.BookID = rentalRequest.BookID
	rental.UserID = rentalRequest.UserID

	updatedRental, err := rentalService.rentalRepository.Update(*rental)

	if err != nil {
		return nil, err
	}

	return updatedRental, nil
}

func NewRentalService(
	rentalRepository rentalRepository.RentalRepository,
	fileupload fileupload.FileUpload,
) RentalService {
	return &RentalServiceImpl{
		rentalRepository,
		fileupload,
	}
}
