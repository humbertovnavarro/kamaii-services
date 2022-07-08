package transaction

import "gorm.io/gorm"

const (
	Withdrawal = "withdrawal"
	Deposit    = "deposit"
	Credit     = "credit"
	Fee        = "fee"
)

type Transaction struct {
	gorm.Model
	Type   string
	Amount int64
	From   string
	To     string
}
