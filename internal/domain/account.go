package domain

import "time"

type Account struct {
	ID        string
	Name      string
	Email     string
	APIKey    string
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccount(name, email string) *Account {
	account := &Account{
		ID:        "1",
		Name:      name,
		Email:     email,
		APIKey:    "some-generated-api-key",
		Balance:   0.0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return account
}
