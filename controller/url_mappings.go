package controller

import (
	"github.com/gin-gonic/gin"
	"marketplace-mvc/db"
	"marketplace-mvc/service"
)

func ConfigureLayers(router *gin.Engine) {
	DB := db.Instance

	addressService := &service.AddressService{DB: DB}
	addressController := &AddressController{service: addressService}
	router.POST("/address", addressController.Create)
	router.GET("/address/:id", addressController.GetById)
	router.GET("/addresses", addressController.GetAll)
	router.PUT("/address", addressController.Update)
	router.DELETE("/address/:id", addressController.Delete)

	paymentService := &service.PaymentService{DB: DB}
	paymentController := &PaymentController{service: paymentService}
	router.POST("/payment", paymentController.Create)
	router.GET("/payment/:id", paymentController.GetById)
	router.GET("/payments", paymentController.GetAll)
	router.PUT("/payment", paymentController.Update)
	router.DELETE("/payment/:id", paymentController.Delete)

	productService := &service.ProductService{DB: DB}
	productController := &ProductController{service: productService}
	router.POST("/product", productController.Create)
	router.GET("/product/:id", productController.GetById)
	router.GET("/products", productController.GetAll)
	router.PUT("/product", productController.Update)
	router.DELETE("/product/:id", productController.Delete)

	purchaseService := &service.PurchaseService{DB: DB}
	purchaseController := &PurchaseController{service: purchaseService}
	router.POST("/purchase", purchaseController.Create)
	router.GET("/purchase/:id", purchaseController.GetById)
	router.GET("/purchases", purchaseController.GetAll)
	router.PUT("/purchase", purchaseController.Update)
	router.DELETE("/purchase/:id", purchaseController.Delete)

	questionService := &service.QuestionService{DB: DB}
	questionController := &QuestionController{service: questionService}
	router.POST("/question", questionController.Create)
	router.GET("/question/:id", questionController.GetById)
	router.GET("/questions", questionController.GetAll)
	router.PUT("/question", questionController.Update)
	router.DELETE("/question/:id", questionController.Delete)

	ratingService := &service.RatingService{DB: DB}
	ratingController := &RatingController{service: ratingService}
	router.POST("/rating", ratingController.Create)
	router.GET("/rating/:id", ratingController.GetById)
	router.GET("/ratings", ratingController.GetAll)
	router.PUT("/rating", ratingController.Update)
	router.DELETE("/rating/:id", ratingController.Delete)

	shippingService := &service.ShippingService{DB: DB}
	shippingController := &ShippingController{service: shippingService}
	router.POST("/shipping", shippingController.Create)
	router.GET("/shipping/:id", shippingController.GetById)
	router.GET("/shippings", shippingController.GetAll)
	router.PUT("/shipping", shippingController.Update)
	router.DELETE("/shipping/:id", shippingController.Delete)

	shoppingCartService := &service.ShoppingCartService{DB: DB}
	shoppingCartController := &ShoppingCartController{service: shoppingCartService}
	router.POST("/shoppingCart", shoppingCartController.Create)
	router.POST("/shoppingCart/buy", shoppingCartController.BuyCart)
	router.GET("/shoppingCart/:id", shoppingCartController.GetById)
	router.GET("/shoppingCarts", shoppingCartController.GetAll)
	router.PUT("/shoppingCart", shoppingCartController.Update)
	router.DELETE("/shoppingCart/:id", shoppingCartController.Delete)

	userService := &service.UserService{DB: DB}
	userController := &UserController{service: userService}
	router.POST("/user", userController.Create)
	router.POST("/user/login", userController.Login)
	router.GET("/user/:id", userController.GetById)
	router.GET("/users", userController.GetAll)
	router.PUT("/user", userController.Update)
	router.DELETE("/user/:id", userController.Delete)
}
