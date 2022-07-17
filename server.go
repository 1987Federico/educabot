package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/golang_api/config"
	"github.com/ydhnwb/golang_api/controller"
	"github.com/ydhnwb/golang_api/middleware"
	"github.com/ydhnwb/golang_api/repository"
	"github.com/ydhnwb/golang_api/service"
	"gorm.io/gorm"
)

var (
	db               *gorm.DB                    = config.SetupDatabaseConnection()
	userRepository   repository.UserRepository   = repository.NewUserRepository(db)
	roleRepository   repository.RoleRepository   = repository.NewRoleRepository(db)
	driverRepository repository.DriverRepository = repository.NewDriverRepository(db)
	jwtService       service.JWTService          = service.NewJWTService()
	userService      service.UserService         = service.NewUserService(userRepository)
	bookService      service.BookService         = service.NewBookService(driverRepository)
	authService      service.AuthService         = service.NewAuthService(userRepository, roleRepository, driverRepository)
	authController   controller.AuthController   = controller.NewAuthController(authService, jwtService)
	userController   controller.UserController   = controller.NewUserController(userService, jwtService)
	bookController   controller.BookController   = controller.NewBookController(bookService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
	}

	userRoutes := r.Group("api/user").Use(middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.POST("/register/driver", authController.Register)
		userRoutes.POST("/role", userController.Profile)
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	bookRoutes := r.Group("api/books").Use(middleware.AuthorizeJWT(jwtService))
	{
		bookRoutes.GET("/", bookController.All)
		bookRoutes.POST("/", bookController.Insert)
		bookRoutes.GET("/:id", bookController.FindByID)
		bookRoutes.PUT("/:id", bookController.Update)
		bookRoutes.DELETE("/:id", bookController.Delete)
	}

	r.Run()
}
