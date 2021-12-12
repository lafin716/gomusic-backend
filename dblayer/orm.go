package dblayer

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/lafin716/gomusic/models"
	"golang.org/x/crypto/bcrypt"
)

type DBORM struct {
	*gorm.DB
}

func (db *DBORM) GetCustomerByID(i int) (models.Customer, error) {
	panic("implement me")
}

func (db *DBORM) GetProduct(u uint) (models.Product, error) {
	panic("implement me")
}

func (db *DBORM) SignOutUserById(i int) error {
	panic("implement me")
}

func NewORM(dbname, con string) (*DBORM, error) {
	db, err := gorm.Open(dbname, con)
	return &DBORM{
		DB: db,
	}, err
}

func hashPassword(s *string) error {
	if s == nil {
		return errors.New("Reference provided for hashing password is nil")
	}

	sBytes := []byte(*s)
	hashedBytes, err := bcrypt.GenerateFromPassword(sBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	*s = string(hashedBytes[:])
	return nil
}

func checkPassword(existHash, incomeHash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(existHash), []byte(incomeHash)) == nil
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
	err := db.Create(&customer).Error
	customer.Pass = ""
	return customer, err
}

func (db *DBORM) SignInUser(email, pass string) (customer models.Customer, err error) {
	result := db.Table("customers").Where(&models.Customer{Email: email})
	err = result.First(&customer).Error
	if err != nil {
		return customer, err
	}

	if !checkPassword(customer.Pass, pass) {
		return customer, ErrINVALIDPASSWORD
	}

	customer.Pass = ""

	err = result.Update("loggedin", 1).Error
	if err != nil {
		return customer, err
	}

	return customer, result.Find(&customer).Error
}

func (db *DBORM) GetCustomerOrdersByID(id int) (orders []models.Order, err error) {
	return orders, db.Table("orders").Select("*").Joins("join customers on customers.id = customer_id").Joins("join products on products.id = product_id").Where("customer_id = ?", id).Scan(&orders).Error
}
