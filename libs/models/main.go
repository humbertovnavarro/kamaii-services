package models

import (
	"github.com/humbertovnavarro/kamaii-services/libs/models/pkg/balance"
	"github.com/humbertovnavarro/kamaii-services/libs/models/pkg/database"
	"github.com/humbertovnavarro/kamaii-services/libs/models/pkg/transaction"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	_db, err := database.New()
	if err != nil {
		panic(err)
	}
	DB = _db
	DB.AutoMigrate(&transaction.Transaction{})
	DB.AutoMigrate(&balance.Balance{})
}
