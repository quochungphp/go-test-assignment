package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/quochungphp/go-test-assignment/src/domain/auth"
	"github.com/quochungphp/go-test-assignment/src/infrastructure/pg_driver"
	"github.com/quochungphp/go-test-assignment/src/pkgs/settings"
	"github.com/tylerb/graceful"
)

func main() {
	// Init postgresql

	pgSession, err := pg_driver.Setup(pg_driver.DBConfiguration{
		Driver:   os.Getenv(settings.DbDriver),
		Host:     os.Getenv(settings.PgHost),
		Port:     os.Getenv(settings.PgPort),
		Database: os.Getenv(settings.PgDB),
		User:     os.Getenv(settings.PgUser),
		Password: os.Getenv(settings.PgPass),
	})
	if err != nil {
		panic(err)
	}
	// Init Gorilla Router
	router := mux.NewRouter()
	authLoginAction := auth.AuthLoginAction{pgSession}
	authCtrl := auth.AuthController{
		authLoginAction,
	}
	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/login", authCtrl.Login).Methods("POST").Name("AuthLoginAction")

	srv := &graceful.Server{
		Timeout: 5 * time.Second,
		BeforeShutdown: func() bool {
			pgSession.Close()
			fmt.Printf("shutting down database connection")
			return true
		},
		Server: &http.Server{
			Addr:    ":" + os.Getenv(settings.Port),
			Handler: router,
		},
	}
	fmt.Printf("Server is runing port: %v\n", os.Getenv(settings.Port))
	if err := srv.ListenAndServe(); err != nil {
		fmt.Errorf("while start server")
	}
}
