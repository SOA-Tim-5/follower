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

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func initTracer() (func(context.Context) error, error) {
	jaegerExporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://localhost:14268/api/traces")))
	if err != nil {
		return nil, err
	}

	res, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("service.name", "followers"),
		),
	)
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(jaegerExporter),
		sdktrace.WithResource(res),
	)

	otel.SetTracerProvider(tp)
	return tp.Shutdown, nil
}

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

	shutdown, err := initTracer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	lis, err := net.Listen("tcp", "localhost:8090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//logger.Println("Server stopped")

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
	tr := otel.Tracer("follower")
	ctx, span := tr.Start(ctx, "CreateNewFollowing")
	defer span.End()
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
		span.RecordError(err)
		println("Database exception: ", err)
	}
	return &follower.FollowerResponseDto{
		Id:           1, //sta ovo treba biti, u preth verziji se salje prazan User
		UserId:       1,
		FollowedById: 1,
	}, nil
}

func (s Server) GetFollowerRecommendations(ctx context.Context, request *follower.Id) (*follower.ListFollowingResponseDto, error) {
	tr := otel.Tracer("follower")
	ctx, span := tr.Start(ctx, "GetFollowerRecommendations")
	defer span.End()
	id := request.Id
	users, err := s.FollowerRepo.GetRecommendations(id)
	if err != nil || users == nil {
		span.RecordError(err)
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
	tr := otel.Tracer("follower")
	ctx, span := tr.Start(ctx, "GetFollowings")
	defer span.End()
	id := request.Id
	users, err := s.FollowerRepo.GetFollowings(id)
	if err != nil || users == nil {
		span.RecordError(err)
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
	tr := otel.Tracer("follower")
	ctx, span := tr.Start(ctx, "GetFollowers")
	defer span.End()
	id := request.Id
	users, err := s.FollowerRepo.GetFollowers(id)
	if err != nil || users == nil {
		span.RecordError(err)
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
	tr := otel.Tracer("follower")
	ctx, span := tr.Start(ctx, "GetAllFromFollowingUsers")
	defer span.End()
	id := request.Id
	users, err := s.FollowerRepo.GetFollowings(id)
	if err != nil || users == nil {
		span.RecordError(err)
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
