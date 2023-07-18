package rental

import (
	"database/sql"
	bookEntity "golang/src/book/entity"
	userEntity "golang/src/user/entity"
)

type Rental struct {
	ID         int64            `json:"id"`
	Status     string           `json:"status"`
	Image      *string          `json:"image"`
	RentDate   sql.NullTime     `json:"rent_date"`
	ReturnDate sql.NullTime     `json:"return_date"`
	BookID     int64            `json:"book_id"`
	Book       *bookEntity.Book `gorm:"foreignKey:BookID; references:ID" json:"book"`
	UserID     int64            `json:"user_id"`
	User       *userEntity.User `gorm:"foreignKey:UserID; references:ID" json:"user"`
}
