package dblayer

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/lafin716/gomusic/models"
)

type DBLayer interface {
	GetAllProducts() ([]models.Product, error)
	GetPromos() ([]models.Product, error)
	GetCustomerByName(string, string) (models.Customer, error)
	GetCustomerByID(int) (models.Customer, error)
	GetProduct(uint) (models.Product, error)
	AddUser(customer models.Customer) (models.Customer, error)
	SignInUser(username, password string) (models.Customer, error)
	SignOutUserById(int) error
	GetCustomerOrdersByID(int) ([]models.Order, error)
}

func (db *DBORM) GetAllProducts() (products []models.Product, err error) {
	return products, db.Find(&products).Error
}

func (db *DBORM) GetPromos() (products []models.Product, err error) {
	return products, db.Where("promotion is not null").Find(&products).Error
}

func (db *DBORM) GetCustomerByName(firstname string, lastname string) (customer models.Customer, err error) {
	return customer, db.Where(&models.Customer{FirstName: firstname, LastName: lastname}).Find(&customer).Error
}

func (db *DBORM) AddUser(customer models.Customer) (models.Customer, error) {
	hashPassword(&customer.Pass)
	customer.LoggedIn = true
	return customer, db.Create(&customer).Error
}

func (db *DBORM) SignInUser(email, pass string) (customer models.Customer, err error) {
	if !checkPassword(pass) {
		return customer, errors.New("Invalid Password")
	}

	result := db.Table("customers").Where(&models.Customer{Email: email})

	err = result.Update("loggedin", 1).Error
	if err != nil {
		return customer, err
	}

	return customer, result.Find(&customer).Error
}

func (db *DBORM) SignInUser(id int) error {
	customer := models.Customer{Model: gorm.Model{
		ID: uint(id),
	}}

	return db.Table("customers").Where(&customer).Update("loggedin", 0).Error
}

func (db *DBORM) GetCustomerOrdersByID(id int) (orders []models.Order, err error) {
	return orders, db.Table("orders").Select("*").Joins("join customers on customers.id = customer_id").Joins("join products on products.id = product_id").Where("customer_id = ?", id).Scan(&orders).Error
}