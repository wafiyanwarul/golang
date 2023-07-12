package book

import (
	"database/sql"
	authorEntity "golang/src/author/entity"
	publisherEntity "golang/src/publisher/entity"
)

type Book struct {
	ID          int64                      `json:"id"`
	Title       string                     `json:"title"`
	Genre       string                     `json:"genre"`
	Stock       int64                      `json:"stock"`
	PublisherID int64                      `json:"publisher_id"`
	Publisher   *publisherEntity.Publisher `gorm:"foreignKey:PublisherID; references:ID" json:"publisher"`
	AuthorID    int64                      `json:"author_id"`
	Author      *authorEntity.Author       `gorm:"foreignKey:AuthorID; references:ID" json:"author"`

	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}
