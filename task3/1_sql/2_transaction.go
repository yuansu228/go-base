package main

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Balance float32
}

type Transaction struct {
	gorm.Model
	FromAccountId uint
	ToAccountId   uint
	Amount        float32
}

func transaction(db *gorm.DB) {
	// db.AutoMigrate(&Account{})
	// db.AutoMigrate(&Transaction{})
	// db.Create(&[]Account{{Balance: 500}, {Balance: 100}})

	err := db.Transaction(func(tx *gorm.DB) error {
		var accountA, accountB Account
		db.First(&accountA, "id = ?", 1)
		db.First(&accountB, "id = ?", 2)
		if accountA.Balance < 100 {
			return errors.New("余额不足")
		}

		accountA.Balance -= 100
		accountB.Balance += 100
		db.Save(&accountA)
		db.Save(&accountB)

		db.Create(&Transaction{FromAccountId: accountA.ID, ToAccountId: accountB.ID, Amount: 100})

		// 返回 nil 提交事务
		return nil
	})
	fmt.Println(err)
}
