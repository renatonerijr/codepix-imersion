package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type Bank struct {
	Id        string     `json:"id" gorm:"type:uuid;primary_key" valid:"required"`
	Code      string     `json:"code" gorm:"type:varchar(20);not null" valid:"required"`
	Name      string     `json:"name" gorm:"type:varchar(255);not null" valid:"required"`
	Accounts  []*Account `gorm:"ForeignKey:BankId" valid:"-"`
	CreatedAt time.Time  `json:"createdAt" valid:"required"`
	UpdatedAt time.Time  `json:"updatedAt" valid:"required"`
}

// this is a method in GO, and it's associate to Struct Bank
func (bank *Bank) isValid() error {
	_, err := govalidator.ValidateStruct(bank)
	if err != nil {
		return err
	}
	return nil
}

// NewBank this is only a function, and it isn't associate to Struct Bank
func NewBank(code string, name string) (*Bank, error) {
	bank := Bank{
		Id:        uuid.NewString(),
		Code:      code,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := bank.isValid()
	if err != nil {
		return nil, err
	}

	return &bank, nil
}
