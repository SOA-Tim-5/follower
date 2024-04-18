package main

import (
	"context"
	handlers "database-example/handler"
	repository "database-example/repo"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	logger := log.New(os.Stdout, "[followers-api] ", log.LstdFlags)
	storeLogger := log.New(os.Stdout, "[followers-store] ", log.LstdFlags)

	store, err := repository.New(storeLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer store.CloseDriverConnection(timeoutContext)
	store.CheckConnection()
	FollowerHandler := handlers.NewFollowersHandler(logger, store)

	//Initialize the router and add a middleware for all the requests
	router := mux.NewRouter()
	router.Use(FollowerHandler.MiddlewareContentTypeSet)

	postUserRouter := router.Methods(http.MethodPost).Subrouter()
	postUserRouter.HandleFunc("/user", FollowerHandler.CreateUser)
	postUserRouter.Use(FollowerHandler.MiddlewarePersonDeserialization)

	getUserRouter := router.Methods(http.MethodGet).Subrouter()
	getUserRouter.HandleFunc("/user/{userId}", FollowerHandler.GetUser)

	postFollowingRouter := router.Methods(http.MethodPost).Subrouter()
	postFollowingRouter.HandleFunc("/follower/create", FollowerHandler.CreateFollowing)
	postFollowingRouter.Use(FollowerHandler.MiddlewareFollowingDeserialization)

	getFollowingsRouter := router.Methods(http.MethodGet).Subrouter()
	getFollowingsRouter.HandleFunc("/followings/{userId}", FollowerHandler.GetFollowings)

	getFollowersRouter := router.Methods(http.MethodGet).Subrouter()
	getFollowersRouter.HandleFunc("/followers/{userId}", FollowerHandler.GetFollowers)

	cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}))

	server := http.Server{
		Addr:         ":8090",
		Handler:      cors(router),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	logger.Println("Server listening on port 8090")
	//Distribute all the connections to goroutines
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, os.Kill)

	sig := <-sigCh
	logger.Println("Received terminate, graceful shutdown", sig)

	//Try to shutdown gracefully
	if server.Shutdown(timeoutContext) != nil {
		logger.Fatal("Cannot gracefully shutdown...")
	}
	logger.Println("Server stopped")
}
