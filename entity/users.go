package entity

import (
	"time"

	"github.com/kdsama/rate-limiter/utils"
)

type User struct {
	ID             string `json:"uuid" bson:"uuid"`
	Email          string `json:"email" bson:"email"`
	Password       string `json:"password" bson:"password"`
	CreatedAt      int64  `bson:"createdAt"`
	UpdatedAt      int64  `bson:"updatedAt"`
	RemainingSlots int32  `bson:"remaining" json:"remaining"`
}

func NewUser(email, password string) *User {
	id := utils.GenerateUUID()
	t := time.Now().Unix()
	return &User{
		ID:             id,
		Email:          email,
		Password:       password,
		CreatedAt:      t,
		RemainingSlots: 5,
	}
}
