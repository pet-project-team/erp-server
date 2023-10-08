package models

import uuid "github.com/satori/go.uuid"

type Transaction struct {
	BaseModel
	OrderId uuid.UUID `gorm:"column:order_id;type:varchar(255);not null"`
	Amount  float64   `gorm:"column:amount;type:float;not null"`
	Status  string    `gorm:"column:status;type:varchar(255);not null"` // in, out
	Note    string    `gorm:"column:note;type:varchar(255);"`
}
