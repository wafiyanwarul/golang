package user

import (
	userEntity "golang/src/user/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() []userEntity.User
	FindById(id int) (*userEntity.User, error)
	FindByEmail(email string) (*userEntity.User, error)
	Update(user userEntity.User) (*userEntity.User, error)
	Create(user userEntity.User) (*userEntity.User, error)
	Delete(id int) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

// Delete implements UserRepository.
func (userRepository *UserRepositoryImpl) Delete(id int) error {
	deleteUser := userRepository.db.Delete(&userEntity.User{}, id)

	if deleteUser.Error != nil {
		return deleteUser.Error
	}

	return nil

}

// Update implements UserRepository.
func (userRepository *UserRepositoryImpl) Update(user userEntity.User) (*userEntity.User, error) {
	updateUser := userRepository.db.Save(&user)

	if updateUser.Error != nil {
		return nil, updateUser.Error
	}

	return &user, nil
}

// FindAll implements UserRepository.
func (userRepository *UserRepositoryImpl) FindAll() []userEntity.User {
	var users []userEntity.User

	userRepository.db.Find(&users)

	return users
}

// FindById implements UserRepository.
func (userRepository *UserRepositoryImpl) FindById(id int) (*userEntity.User, error) {
	var user userEntity.User

	dataUser := userRepository.db.First(&user, id)

	if dataUser.Error != nil {
		return nil, dataUser.Error
	}

	return &user, nil
}

// FindByEmail implements UserRepository.
func (userRepository *UserRepositoryImpl) FindByEmail(email string) (*userEntity.User, error) {
	var user userEntity.User

	dataUser := userRepository.db.Where("email = ?", email).First(&user)

	if dataUser.Error != nil {
		return nil, dataUser.Error
	}

	return &user, nil
}

// Create implements UserRepository.
func (userRepository *UserRepositoryImpl) Create(user userEntity.User) (*userEntity.User, error) {

	dataUser := userRepository.db.Create(&user)

	if dataUser.Error != nil {
		return nil, dataUser.Error
	}

	return &user, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}
