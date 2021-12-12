package rest

import "github.com/gin-gonic/gin"

func RunApiWithHandler(address string, h HandlerInterface) error {
	r := gin.Default()

	r.GET("/products", h.GetProducts)

	r.GET("/promos", h.GetPromos)

	userGroup := r.Group("/user")
	{
		userGroup.POST("/:id/signout", h.SignOut)
		userGroup.GET("/:id/orders", h.GetOrders)
	}

	usersGroup := r.Group("/users")
	{
		usersGroup.POST("/signin", h.SignIn)

		usersGroup.POST("/charge", h.Charge)

		usersGroup.POST("", h.AddUser)
	}

	return r.Run(address)
}

func RunApi(address string) error {
	h, err := NewHandler()
	if err != nil {
		return err
	}

	return RunApiWithHandler(address, h)
}
