package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kollekcioner47/finance-app/internal/config"
	"github.com/kollekcioner47/finance-app/internal/handlers"
	"github.com/kollekcioner47/finance-app/internal/middleware"
	"github.com/kollekcioner47/finance-app/internal/repository"
	"github.com/kollekcioner47/finance-app/internal/service"
	"github.com/kollekcioner47/finance-app/internal/session"
)

func main() {
	cfg := config.Load()
	session.InitStore(cfg.SessionKey)
	// Подключение к БД
	db, err := repository.NewPostgresDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Миграции
	if err := repository.RunMigrations(cfg.DatabaseURL); err != nil {
		log.Fatal(err)
	}

	// Инициализация репозиториев и сервисов
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	transactionRepo := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepo)

	categoryHandler := handlers.NewCategoryHandler(categoryService)
	transactionHandler := handlers.NewTransactionHandler(transactionService, categoryService)

	//_ = userService // говорит компилятору: "я знаю, что переменная не используется, не ругайся"
	userHandler := handlers.NewUserHandler(userService)
	// Инициализация обработчиков (пока без userHandler, добавим позже)
	// userHandler := handlers.NewUserHandler(userService)

	r := mux.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Публичные маршруты
	r.HandleFunc("/register", userHandler.Register).Methods("GET", "POST")
	r.HandleFunc("/login", userHandler.Login).Methods("GET", "POST")
	r.HandleFunc("/logout", userHandler.Logout).Methods("POST")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))
	// Защищённые маршруты
	protected := r.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthRequired)
	protected.HandleFunc("/", handlers.Index).Methods("GET")
	protected.HandleFunc("/categories", categoryHandler.List).Methods("GET")
	protected.HandleFunc("/categories/create", categoryHandler.CreateForm).Methods("GET")
	protected.HandleFunc("/categories", categoryHandler.Create).Methods("POST")
	protected.HandleFunc("/transactions", transactionHandler.List).Methods("GET")
	protected.HandleFunc("/transactions/create", transactionHandler.CreateForm).Methods("GET")
	protected.HandleFunc("/transactions", transactionHandler.Create).Methods("POST")

	// Заглушка для главной
	r.HandleFunc("/", handlers.Index).Methods("GET")

	log.Printf("Server starting on :%s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
