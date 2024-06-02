package models

import "time"

type (
	User struct {
		ID            string
		CreatedAt     time.Time
		UpdatedAt     *time.Time
		Rating        int
		PaymentStatus string
	}

	UserEntity struct {
	ID string
	CreatedAt     time.Time
	UpdatedAt     *time.Time
	Rating        int
	PaymentStatus string
	}

	UserStat struct {
		ID string
		PaymentStatus string
	}
)

func (u *UserEntity) Convert() *User {
	if u != nil {
		return &User{
			ID: u.ID,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
			Rating: u.Rating,
			PaymentStatus: u.PaymentStatus,
		}
	}

	return nil
}