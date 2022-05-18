package messages

import "time"

type Product struct {
	Id         string     `json:"Id"`
	Name       *string    `json:"Name,omitempty"`
	Dollars    *uint64    `json:"Dollars,omitempty"`
	Cents      *uint32    `json:"Cents,omitempty"`
	Amount     *uint32    `json:"Amount,omitempty"`
	IsImported *bool      `json:"IsImported,omitempty"`
	CreatedAt  *time.Time `json:"CreatedAt,omitempty"`
	UpdatedAt  *time.Time `json:"UpdatedAt,omitempty"`
}
