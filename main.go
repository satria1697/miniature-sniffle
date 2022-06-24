package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	handler2 "six/auth/handler"
	"six/auth/repository"
	service2 "six/auth/service"
	"six/database"
	"six/user/handler"
	"six/user/service"
	util "six/utils"
	"time"
)

func main() {
	config := util.InitConfig()
	db, _ := database.InitDatabase(config)

	r := mux.NewRouter()
	authRepository := repository.NewAuthRepository(db)
	authService := service2.NewAuthService(authRepository)
	handler2.NewAuthHandler(r, authService)

	userService := service.NewUserService()
	handler.NewUserHandler(r, userService)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("run err: %v\n", err)
	}
}
