package repository

import (
	"context"
	"database-example/model"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// NoSQL: MovieRepo struct encapsulating Neo4J api client
type FollowerRepository struct {
	// Thread-safe instance which maintains a database connection pool
	driver neo4j.DriverWithContext
	logger *log.Logger
}

// NoSQL: Constructor which reads db configuration from environment and creates a keyspace
func New(logger *log.Logger) (*FollowerRepository, error) {
	// Local instance
	uri := "bolt://neo4j:7687"
	user := "neo4j"
	pass := "ivanaanja"
	auth := neo4j.BasicAuth(user, pass, "")

	driver, err := neo4j.NewDriverWithContext(uri, auth)
	if err != nil {
		logger.Panic(err)
		return nil, err
	}

	// Return repository with logger and DB session
	return &FollowerRepository{
		driver: driver,
		logger: logger,
	}, nil
}

// Check if connection is established
func (mr *FollowerRepository) CheckConnection() {
	ctx := context.Background()
	err := mr.driver.VerifyConnectivity(ctx)
	if err != nil {
		mr.logger.Panic(err)
		return
	}
	// Print Neo4J server address
	mr.logger.Printf(`Neo4J server address: %s`, mr.driver.Target().Host)
}

// Disconnect from database
func (mr *FollowerRepository) CloseDriverConnection(ctx context.Context) {
	mr.driver.Close(ctx)
}

func (mr *FollowerRepository) SaveUser(user *model.User) (bool, error) {
	userInDatabase, err := mr.ReadUser(user.Id)
	if (userInDatabase == model.User{}) {
		err = mr.WriteUser(user)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	if err != nil {
		return false, err
	}
	return false, nil
}

func (mr *FollowerRepository) WriteUser(user *model.User) error {
	ctx := context.Background()
	session := mr.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)
	newUser, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			result, err := transaction.Run(ctx,
				"CREATE (u:User) SET u.Id = $id, u.Username = $username, u.Image = $image RETURN u.Username + ', from node ' + id(u)",
				map[string]any{"id": user.Id, "username": user.Username, "image": user.Image})
			if err != nil {
				return nil, err
			}

			if result.Next(ctx) {
				return result.Record().Values[0], nil
			}

			return nil, result.Err()
		})
	if err != nil {
		mr.logger.Println("Error inserting User:", err)
		return err
	}
	mr.logger.Println(newUser.(string))
	return nil
}

func (mr *FollowerRepository) ReadUser(userId string) (model.User, error) {
	ctx := context.Background()
	session := mr.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)
	user, err := session.ExecuteRead(ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			result, err := transaction.Run(ctx,
				"MATCH (u {Id: $id}) RETURN u.Id, u.Username, u.Image",
				map[string]any{"id": userId})
			if err != nil {
				return nil, err
			}

			if result.Next(ctx) {
				return result.Record().Values, nil
			}

			return nil, result.Err()
		})
	if err != nil {
		mr.logger.Println("Error reading user:", err)
		return model.User{}, err
	}
	if user == nil {
		return model.User{}, nil
	}
	var id, username, image string
	for _, value := range user.([]interface{}) {
		if val, ok := value.(string); ok {
			if id == "" {
				id = val
			} else if username == "" {
				username = val
			} else if image == "" {
				image = val
			}
		}
	}
	userFromDatabase := model.User{
		Id:       id,
		Username: username,
		Image:    image,
	}

	return userFromDatabase, nil
}
func (mr *FollowerRepository) SaveFollowing(user *model.User, userToFollow *model.User) error {
	ctx := context.Background()
	session := mr.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)
	mr.SaveUser(user)
	mr.SaveUser(userToFollow)
	_, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			result, err := transaction.Run(ctx,
				"MATCH (a:User), (b:User) WHERE a.Username = $username AND b.username = $followUsername CREATE (a)-[r: IS_FOLLOWING]->(b) RETURN type(r)",
				map[string]any{"username": user.Username, "followUsername": userToFollow.Username})
			if err != nil {
				return nil, err
			}
			if result.Next(ctx) {
				return result.Record().Values[0], nil
			}
			return nil, result.Err()
		})
	if err != nil {
		mr.logger.Println("Error inserting following:", err)
		return err
	}
	return nil
}
