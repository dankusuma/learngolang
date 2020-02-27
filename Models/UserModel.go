package Models

import (
	"time"
)

type User struct {
	ID         string    `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Birth_Date time.Time `json:"birth_date"`
	Password   string    `json:"password"`
}
