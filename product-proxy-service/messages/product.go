package messages

import "time"

type Product struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Dollars    uint64    `json:"dollars"`
	Cents      uint8     `json:"cents"`
	Amount     uint32    `json:"amount"`
	IsImported bool      `json:"is_imported"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}
