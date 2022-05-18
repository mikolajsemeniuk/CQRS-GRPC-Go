package messages

import "time"

type Product struct {
	Id         string    `json:"Id"`
	Name       string    `json:"Name"`
	Dollars    uint64    `json:"Dollars"`
	Cents      uint32    `json:"Cents"`
	Amount     uint32    `json:"Amount"`
	IsImported bool      `json:"IsImported"`
	CreatedAt  time.Time `json:"CreatedAt"`
	UpdatedAt  time.Time `json:"UpdatedAt,omitempty"`
}
