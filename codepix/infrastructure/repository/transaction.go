package repository

import (
	"codepix/domain/model"
	"fmt"

	"gorm.io/gorm"
)

type TransactionRepositoryDb struct {
	Db *gorm.DB
}

func (t *TransactionRepositoryDb) Register(transaction *model.Transaction) error {
	err := t.Db.Create(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionRepositoryDb) Save(transaction *model.Transaction) error {
	err := t.Db.Save(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionRepositoryDb) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction

	t.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.Id == "" {
		return nil, fmt.Errorf("no key was found")
	}

	return &transaction, nil
}
