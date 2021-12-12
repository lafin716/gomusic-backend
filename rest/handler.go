package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/lafin716/gomusic/dblayer"
	"github.com/lafin716/gomusic/models"
	"net/http"
	"strconv"
)

type HandlerInterface interface {
	GetProducts(c *gin.Context)
	GetPromos(c *gin.Context)
	AddUser(c *gin.Context)
	SignIn(c *gin.Context)
	SignOut(c *gin.Context)
	GetOrders(c *gin.Context)
	Charge(c *gin.Context)
}

type Handler struct {
	db dblayer.DBLayer
}

func NewHandler(db, constring string) (HandlerInterface, error) {

	database, err := dblayer.NewORM(db, constring)
	if err != nil {
		return nil, err
	}

	return &Handler{
		db: database,
	}, nil
}

func (h *Handler) GetProducts(c *gin.Context) {
	if h.db == nil {
		return
	}

	products, err := h.db.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *Handler) GetPromos(c *gin.Context) {
	if h.db == nil {
		return
	}

	products, err := h.db.GetPromos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *Handler) SignIn(c *gin.Context) {
	if h.db == nil {
		return
	}

	var customer models.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer, err = h.db.SignInUser(customer.Email, customer.Pass)
	if err != nil {
		if err == dblayer.ErrINVALIDPASSWORD {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func (h *Handler) AddUser(c *gin.Context) {
	if h.db == nil {
		return
	}

	var customer models.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer, err = h.db.AddUser(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (h *Handler) SignOut(c *gin.Context) {
	if h.db == nil {
		return
	}

	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.db.SignOutUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func (h *Handler) GetOrders(c *gin.Context) {
	if h.db == nil {
		return
	}

	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orders, err := h.db.GetCustomerOrdersByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (h *Handler) Charge(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}
}
