package balance

import "gorm.io/gorm"

type Balance struct {
	gorm.Model
	Balance int64
	Owner   string
}
