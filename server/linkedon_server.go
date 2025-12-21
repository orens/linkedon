package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
	"github.com/orens/linkedon/linkedon"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	NEO4JURI      = "neo4j://localhost"
	NEO4JUSER     = "neo4j"
	NEO4JPASSWORD = "rootroot"
	NEO4JDATABASE = "neo4j"
)

type LinkedonServer struct {
	linkedon.UnimplementedLinkedonServiceServer
	driver neo4j.Driver
}

func NewLinkedonServer(driver neo4j.Driver) *LinkedonServer {
	return &LinkedonServer{
		driver: driver,
	}
}

func (s *LinkedonServer) CreatePerson(ctx context.Context, req *linkedon.CreatePersonRequest) (*linkedon.CreatePersonResponse, error) {
	person := req.GetPerson()
	if person == nil {
		return nil, status.Error(codes.InvalidArgument, "person is required")
	}
	session := s.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: NEO4JDATABASE, AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)
	result, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		return tx.Run(ctx, `
		CREATE (p:Person {name: $name, age: $age, id: $id})
	`, map[string]any{
			"id":   person.Id,
			"name": person.Name,
			"age":  person.Age,
		})
	})
	if err != nil {
		log.Printf("failed to create person: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "failed to create person: %v", err)
	}
	// summary := result.Summary()
	fmt.Printf("Created: %v\n", result) // TODO

	return &linkedon.CreatePersonResponse{Success: true}, nil
}

func (s *LinkedonServer) Follow(ctx context.Context, req *linkedon.FollowRequest) (*linkedon.FollowResponse, error) {
	followerId := req.GetFollowerId()
	followeeId := req.GetFolloweeId()
	if followerId == 0 || followeeId == 0 {
		return nil, status.Error(codes.InvalidArgument, "followerId and followeeId are required")
	}
	session := s.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: NEO4JDATABASE, AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)
	result, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		return tx.Run(ctx, `
	MATCH (follower:Person {id: $followerId})
	MATCH (followee:Person {id: $followeeId})
	CREATE (follower)-[:FOLLOWS]->(followee)
	`, map[string]any{
			"followerId": followerId,
			"followeeId": followeeId,
		})
	})
	if err != nil {
		log.Printf("failed to follow: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "failed to follow: %v", err)
	}
	// summary := result.Summary()
	fmt.Printf("Followed: %v\n", result) // TODO
	return &linkedon.FollowResponse{Success: true}, nil
}

func serve(driver neo4j.Driver) error {
	port := 8044
	log.Printf("Starting linkedon server on port %d", port)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}
	log.Printf("Listening on %v", lis.Addr())
	grpcServer := grpc.NewServer()
	linkedon.RegisterLinkedonServiceServer(grpcServer, NewLinkedonServer(driver))
	log.Println("Registered services")
	log.Println("Serving...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return err
	}
	log.Println("Server stopped")
	return nil
}

func newNeo4jDriver(ctx context.Context) (neo4j.Driver, error) {
	driver, err := neo4j.NewDriver(
		NEO4JURI,
		neo4j.BasicAuth(NEO4JUSER, NEO4JPASSWORD, ""))
	if err != nil {
		log.Fatalf("failed to connect to neo4j: %v", err)
		return nil, err
	}
	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		log.Fatalf("failed to verify connectivity to neo4j: %v", err)
		return nil, err
	}
	fmt.Println("Connection established.")
	return driver, nil
}

func setupNeo4j(ctx context.Context, driver neo4j.Driver) error {
	session := driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: NEO4JDATABASE, AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)
	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		return tx.Run(ctx, `
		CREATE CONSTRAINT person_id IF NOT EXISTS FOR (p:Person) REQUIRE p.id IS UNIQUE
	`, map[string]any{})
	})
	if err != nil {
		return fmt.Errorf("failed to setup neo4j: %v", err)
	}
	_, err = session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		return tx.Run(ctx, `
		MATCH (n) DETACH DELETE n
	`, map[string]any{})
	})
	if err != nil {
		return fmt.Errorf("failed to setup neo4j: %v", err)
	}
	fmt.Printf("Setup neo4j: done\n")
	return nil
}

func main() {
	ctx := context.Background()
	driver, err := newNeo4jDriver(ctx)
	if err != nil {
		log.Fatalf("failed to create neo4j driver: %v", err)
		panic(err)
	}
	defer driver.Close(ctx)
	if err := setupNeo4j(ctx, driver); err != nil {
		log.Fatalf("failed to setup neo4j: %v", err)
		panic(err)
	}
	if err := serve(driver); err != nil {
		log.Fatalf("failed to serve: %v", err)
		panic(err)
	}
}
