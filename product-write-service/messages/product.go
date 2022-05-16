package messages

import "time"

type Product struct {
	Id         string     `json:"id"`
	Name       *string    `json:"name,omitempty"`
	Dollars    *uint64    `json:"dollars,omitempty"`
	Cents      *uint32    `json:"cents,omitempty"`
	Amount     *uint32    `json:"amount,omitempty"`
	IsImported *bool      `json:"is_imported,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
}
