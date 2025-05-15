package models

import "time"

type Product struct {
	ID       string `gorm:"primaryKey"`
	Name     string
	Category string
}

type Customer struct {
	ID      string `gorm:"primaryKey"`
	Name    string
	Email   string
	Address string
}

type Order struct {
	ID            string `gorm:"primaryKey"`
	ProductID     string
	CustomerID    string
	Region        string
	DateOfSale    time.Time
	QuantitySold  int
	UnitPrice     float64
	Discount      float64
	ShippingCost  float64
	PaymentMethod string
	Product       Product  `gorm:"foreignKey:ProductID"`
	Customer      Customer `gorm:"foreignKey:CustomerID"`
}
