package main

import (
	"context"
	"database-example/model"
	"database-example/proto/follower"
	repository "database-example/repo"
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	/*
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

		getRecommendationsRouter := router.Methods(http.MethodGet).Subrouter()
		getRecommendationsRouter.HandleFunc("/recommendations/{userId}", FollowerHandler.GetRecommendations)

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
		logger.Println("Server stopped")*/

	lis, err := net.Listen("tcp", "localhost:8090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	follower.RegisterFollowerServer(grpcServer, Server{FollowerRepo: store}) //da li store?
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)

}

type Server struct {
	follower.UnimplementedFollowerServer
	FollowerRepo *repository.FollowerRepository
}

// sta je ova metoda, ne valjaju joj ni parametri
func (s Server) CreateUser(ctx context.Context, request *follower.FollowingResponseDto) {
	user := model.User{
		Id:       request.Id, //provjeriti
		Username: request.Username,
		Image:    request.Image,
	}
	userSaved, err := s.FollowerRepo.SaveUser(&user)
	if err != nil {
		println("Error while creating a new user")
		return
	}
	if userSaved {
		println("New user saved to database")

	}
}

func (s Server) CreateNewFollowing(ctx context.Context, request *follower.UserFollowingDto) (*follower.FollowerResponseDto, error) {
	newFollowing := model.UserFollowing{
		UserId:            request.UserId,
		Username:          request.Username,
		Image:             request.Image,
		FollowingUserId:   request.FollowingUserId,
		FollowingUsername: request.FollowingUsername,
		FollowingImage:    request.FollowingImage,
	}
	user := model.User{}
	userToFollow := model.User{}
	user.Id = newFollowing.UserId
	user.Username = newFollowing.Username
	user.Image = newFollowing.Image
	userToFollow.Id = newFollowing.FollowingUserId
	userToFollow.Username = newFollowing.FollowingUsername
	userToFollow.Image = newFollowing.FollowingImage
	println("djnjsdnskndksnd" + userToFollow.Username)
	err := s.FollowerRepo.SaveFollowing(&user, &userToFollow)
	if err != nil {
		println("Database exception: ", err)
	}
	return &follower.FollowerResponseDto{
		Id:           1, //sta ovo treba biti, u preth verziji se salje prazan User
		UserId:       1,
		FollowedById: 1,
	}, nil
}

func (s Server) GetFollowerRecommendations(ctx context.Context, request *follower.Id) (*follower.ListFollowingResponseDto, error) {
	id := request.Id
	users, err := s.FollowerRepo.GetRecommendations(id)
	if err != nil || users == nil {
		println("Database exception: ", err)
		return &follower.ListFollowingResponseDto{
			ResponseList: make([]*follower.FollowingResponseDto, 0), //da se vrati prazna
		}, nil
	}

	responseList := make([]*follower.FollowingResponseDto, len(users))
	for i, user := range users {
		responseList[i] = &follower.FollowingResponseDto{
			Id:       user.Id,
			Username: user.Username,
			Image:    user.Image,
		}
	}
	return &follower.ListFollowingResponseDto{
		ResponseList: responseList,
	}, nil
}

func (s Server) GetFollowings(ctx context.Context, request *follower.Id) (*follower.ListFollowingResponseDto, error) {
	id := request.Id
	users, err := s.FollowerRepo.GetFollowings(id)
	if err != nil || users == nil {
		println("Database exception: ", err)
		return &follower.ListFollowingResponseDto{
			ResponseList: make([]*follower.FollowingResponseDto, 0), //da se vrati prazna
		}, nil
	}
	responseList := make([]*follower.FollowingResponseDto, len(users))
	for i, user := range users {
		responseList[i] = &follower.FollowingResponseDto{
			Id:       user.Id,
			Username: user.Username,
			Image:    user.Image,
		}
	}
	return &follower.ListFollowingResponseDto{
		ResponseList: responseList,
	}, nil

}

func (s Server) GetFollowers(ctx context.Context, request *follower.Id) (*follower.ListFollowingResponseDto, error) {
	id := request.Id
	users, err := s.FollowerRepo.GetFollowers(id)
	if err != nil || users == nil {
		println("Database exception: ", err)
		return &follower.ListFollowingResponseDto{
			ResponseList: make([]*follower.FollowingResponseDto, 0), //da se vrati prazna
		}, nil
	}
	responseList := make([]*follower.FollowingResponseDto, len(users))
	for i, user := range users {
		responseList[i] = &follower.FollowingResponseDto{
			Id:       user.Id,
			Username: user.Username,
			Image:    user.Image,
		}
	}
	return &follower.ListFollowingResponseDto{
		ResponseList: responseList,
	}, nil
}

func (s Server) GetAllFromFollowingUsers(ctx context.Context, request *follower.Id) (*follower.BlogListResponse, error) {
	id := request.Id
	users, err := s.FollowerRepo.GetFollowings(id)
	if err != nil || users == nil {
		println("Database exception: ", err)
		return &follower.BlogListResponse{
			Following: make([]*follower.FollowingResponseDto, 0), //da se vrati prazna
		}, nil
	}
	responseList := make([]*follower.FollowingResponseDto, len(users))
	for i, user := range users {
		responseList[i] = &follower.FollowingResponseDto{
			Id:       user.Id,
			Username: user.Username,
			Image:    user.Image,
		}
	}
	return &follower.BlogListResponse{
		Following: responseList,
	}, nil
}
