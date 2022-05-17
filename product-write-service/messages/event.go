package messages

import "time"

type Event struct {
	Method    string    `json:"method"`
	Data      Product   `json:"data"`
	Timestamp time.Time `json:"timestamp"`
}
