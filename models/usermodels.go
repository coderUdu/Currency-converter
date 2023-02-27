package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	User_ID      *string            `json: "user_id"`
	userType     *string            `json: "usertype" validate:"required, eq=ADMIN|eq=USER"`
	Balance      float64            `json:"balance"`
	BalanceRecord []float64         `json:"balancerecord"` 
	Password     *string            `json:"password" validate:"required,min = 8"`
	HashPassword *string            `json:"Hashpassword"`
	Token        *string            `json:"token"`
	RefreshToken *string            `json: "refreshtoken", validate:"Required"`
	Email        *string            `json: "email" validate:"required"`
	CreatedAt    time.Time          `json: "createdat"`
	UpdatedAt    time.Time          `json: "updatedat"`
}

type Currency struct {
	BaseCurrency string `json: "basecurrency"`
	Currencies   `json: "currencies"`
}
type Currencies struct {
	USD float64 `json:"USD"`
	NGN float64 `json:"NGN"`
}
