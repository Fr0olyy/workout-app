package main

import (
	"fmt"
	"log"
	"traning/internal/db"
	"traning/internal/handlers"
	middle "traning/internal/middleware"
	"traning/internal/repository"
	"traning/internal/services"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	//Загрузка конфига к бд из .env
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading .env file")
	}
	fmt.Println("Config will load")

	//Подключение к бд
	db, err := db.InitDB()
	if err != nil {
		log.Fatalf("Server not connect to database: %v", err)
	}
	fmt.Println("Server connect to db")

	//Инициализация Echo
	e := echo.New()
	repos := repository.NewRepository(db)
	fmt.Println("Repository working")

	services := services.NewService(repos)
	fmt.Println("Services working")

	handlers := handlers.NewHandler(services)
	fmt.Println("Handlers working")

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	// Регистрируем обработчики ошибок
	e.HTTPErrorHandler = middle.ErrorHandler
	e.GET("*", middle.NotFoundHandler)

	api := e.Group("/api", handlers.UserIdentity)
	api.GET("/profile", handlers.GetUserProfile)

	auth := e.Group("/auth")
	auth.POST("/sign-in", handlers.SignIn)
	auth.POST("/sign-up", handlers.SignUp)

	e.Logger.Fatal(e.Start("localhost:5173"))
}
