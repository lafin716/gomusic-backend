package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Product struct {
	gorm.Model
	Image string `json:"img"`
	ImageAlt string `json:"imgalt"`
	Price float64 `json:"price"`
	Promotion float64 `json:"promotion"`
	ProductName string `json:"productname"`
	Description string `json:"desc"`
}

func (Product) TableName() string {
	return "products"
}

type Customer struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Email string `json:"email"`
	Pass string `json:"password"`
	LoggedIn bool `json:"loggedin"`
}

func (Customer) TableName() string {
	return "customers"
}

type Order struct {
	gorm.Model
	Product
	Customer
	CustomerID int `json:"customer_id"`
	ProductID int `json:"product_id"`
	Price float64 `json:"sell_price"`
	PurchaseDate time.Time `json:"purchase_date"`
}

func (Order) TableName() string {
	return "orders"
}