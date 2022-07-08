package account

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	ID             string
	HashedPassword string
	OAuth          string
}
