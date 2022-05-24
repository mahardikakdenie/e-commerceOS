package routes

import (
	"api/controller"
	"api/middleware"
	"api/modules/auth"
	"api/modules/cart"
	cart_customer_repo "api/modules/cart/customer/repository"
	cart_customer_service "api/modules/cart/customer/service"
	category_customer "api/modules/category/customer"

	cart_customer_controller "api/modules/cart/customer/controller"
	category_store "api/modules/category/store"
	"api/modules/customer"
	"api/modules/customer_auth"
	orders "api/modules/order"
	product_customer "api/modules/product/customer"
	product_store "api/modules/product/store"
	"api/modules/role"
	"api/modules/store"
	"api/modules/tshirt"
	"api/routes/module"
	"time"

	"api/modules/user"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"gorm.io/gorm"
)

func Router(db *gorm.DB, router gin.IRouter) {

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE,PATCH,PUT",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userController := controller.NewController(userService)

	//

	authRepository := auth.NewRepository(db)
	authService := auth.NewService(authRepository)
	authController := controller.NewAuthController(authService)

	//

	roleRepository := role.NewRepository(db)
	roleService := role.NewService(roleRepository)
	roleController := controller.NewRoleController(roleService)

	//

	tshirtRepository := tshirt.NewRepository(db)
	tshirtService := tshirt.NewService(tshirtRepository)
	tshirtController := controller.NewTShirtController(tshirtService)

	//
	orderRepository := orders.NewRepository(db)
	orderService := orders.NewService(orderRepository)
	orderController := controller.NewOrderController(orderService)

	cartRepository := cart.NewRepository(db)
	cartService := cart.NewService(cartRepository)
	cartController := controller.NewCartController(cartService)

	//
	storeRepository := store.NewRepository(db)
	storeService := store.NewService(storeRepository)
	storeController := controller.NewStoreController(storeService)

	//
	customerRepository := customer.NewRepository(db)
	customerService := customer.NewService(customerRepository)
	customerController := controller.NewCustomerController(customerService)

	//
	customerAuthRepository := customer_auth.NewRepository(db)
	customerAuthService := customer_auth.NewService(customerAuthRepository)
	customerAuthController := controller.NewCustomerAuthController(customerAuthService)

	//
	categoryStoreRepository := category_store.NewRepository(db)
	categoryStoreService := category_store.NewService(categoryStoreRepository)
	categoryStoreController := controller.NewCategoryStoreController(categoryStoreService)

	//
	productStoreRepository := product_store.NewRepository(db)
	productStoreService := product_store.NewService(productStoreRepository)
	productStoreController := controller.NewProductStoreController(productStoreService)

	//
	categoryCustomerRepository := category_customer.NewRepository(db)
	categoryCustomerService := category_customer.NewService(categoryCustomerRepository)
	categoryCustomerController := controller.NewCategoryCustomerController(categoryCustomerService)

	//
	productCustomerRepository := product_customer.NewRepository(db)
	productCustomerService := product_customer.NewService(productCustomerRepository)
	productCustomerController := controller.NewProductCustomerController(productCustomerService)

	//

	cartCustomerRepo := cart_customer_repo.NewRepository(db)
	cartCustomerService := cart_customer_service.NewService(&cartCustomerRepo, customerRepository, productStoreRepository)
	cartCustomerController := cart_customer_controller.NewCartsController(cartCustomerService)

	MyMiddleware := middleware.MyMiddleware(authService)
	CustomerMiddleware := middleware.CustomerMiddleware(customerAuthService)
	cors := middleware.SetupCorsMiddleware()

	v1 := router.Group("v1")
	v1.Use(cors)

	// routes
	module.AuthRoute(v1, authController)
	module.UserRoute(v1, userController, MyMiddleware)
	module.RoleRoute(v1, roleController, MyMiddleware)
	module.TshirtRoute(v1, tshirtController, MyMiddleware)
	module.OrderRoute(v1, orderController, MyMiddleware)
	module.CartRoute(v1, cartController, MyMiddleware)
	module.StoreRoute(v1, storeController, MyMiddleware, cors)
	module.CustomerRoute(v1, customerController, CustomerMiddleware, cors)
	module.CustomerAuthRoute(v1, customerAuthController, CustomerMiddleware, cors)
	module.CategoryStoreRoute(v1, categoryStoreController, MyMiddleware, cors)
	module.ProductStoreRoute(v1, productStoreController, MyMiddleware)
	module.CategoryCustomerRoute(v1, categoryCustomerController, cors)
	module.ProductCustomerRoute(v1, productCustomerController, cors)
	module.CartCustomerRoute(v1, cartCustomerController, cors, CustomerMiddleware)
}
