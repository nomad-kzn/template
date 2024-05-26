package models

import "time"

type (
	User struct {
		ID string
		CreatedAt time.Time
		UpdatedAt *time.Time
		Rating int
		PaymentStatus string
	}
)