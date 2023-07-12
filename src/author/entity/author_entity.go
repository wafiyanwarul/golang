package author

import "database/sql"

type Author struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`

	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}
