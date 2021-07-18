package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username, Email, Password string
}

type Account struct {
	gorm.Model
	Type, Name      string
	Balance, UserId uint
}

type AccountsResponse struct {
	ID      int
	Name    string
	Balance uint
}

type UserResponse struct {
	ID              int
	UserName, Email string
	Accounts        []AccountsResponse
}

type LoginRequest struct {
	Username, Password string
}

type Validation struct {
	Value, Valid string
}

type RegisterRequest struct {
	Username, Email, Password string
}

type Transaction struct {
	gorm.Model
	From, To uint
	Amount   int
}

type TransactionRequest struct {
	UserId, From, To uint
	Amount           int
}
