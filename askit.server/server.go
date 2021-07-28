package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mirzaceligija/askit/config"
	"github.com/mirzaceligija/askit/controller"
	"github.com/mirzaceligija/askit/middleware"
	"github.com/mirzaceligija/askit/repository"
	"github.com/mirzaceligija/askit/service"
	"gorm.io/gorm"
)

var (
	db                 *gorm.DB                      = config.SetupDatabaseConnection()
	userRepository     repository.UserRepository     = repository.NewUserRepository(db)
	questionRepository repository.QuestionRepository = repository.NewQuestionRepository(db)
	answerRepository   repository.AnswerRepository   = repository.NewAnswerRepository(db)
	jwtService         service.JWTService            = service.NewJWTService()
	authService        service.AuthService           = service.NewAuthService(userRepository)
	questionService    service.QuestionService       = service.NewQuestionService(questionRepository)
	answerService      service.AnswerService         = service.NewAnswerService(answerRepository)
	userService        service.UserService           = service.NewUserService(userRepository)
	authController     controller.AuthController     = controller.NewAuthController(authService, jwtService)
	userController     controller.UserController     = controller.NewUserController(userService, jwtService)
	answerController   controller.AnswerController   = controller.NewAnswerController(answerService, jwtService)
	questionController controller.QuestionController = controller.NewQuestionController(questionService, jwtService)
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		c.Next()
	}
}

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	r.Use(CORS())

	authRoutes := r.Group("api/v1/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/v1/users")
	{
		//authRoutes.GET("/", userController.All)
		userRoutes.GET("/profile", middleware.AuthorizeJWT(jwtService), userController.Profile)
		userRoutes.PUT("/profile", middleware.AuthorizeJWT(jwtService), userController.Update)
		userRoutes.GET("/questions", middleware.AuthorizeJWT(jwtService), questionController.FindByUserID)
	}

	questionRoutes := r.Group("api/v1/questions")
	{
		questionRoutes.GET("/", questionController.All)
		questionRoutes.POST("/", middleware.AuthorizeJWT(jwtService), questionController.Insert)
		questionRoutes.GET("/:id", questionController.FindByID)
		questionRoutes.PUT("/:id", middleware.AuthorizeJWT(jwtService), questionController.Update)
		questionRoutes.DELETE("/:id", middleware.AuthorizeJWT(jwtService), questionController.Delete)
	}

	answerRoutes := r.Group("api/v1/answers")
	{
		answerRoutes.POST("/:questionId", middleware.AuthorizeJWT(jwtService), answerController.Insert)
		answerRoutes.PUT("/:id", middleware.AuthorizeJWT(jwtService), answerController.Update)
		answerRoutes.DELETE("/:id", middleware.AuthorizeJWT(jwtService), answerController.Delete)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
