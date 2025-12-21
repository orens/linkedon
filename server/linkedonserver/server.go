package linkedonserver

import (
	"context"
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
	"github.com/orens/linkedon/linkedon"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const NEO4JDATABASE = "neo4j"

type LinkedonServer struct {
	linkedon.UnimplementedLinkedonServiceServer
	driver neo4j.Driver
}

func NewLinkedonServer(driver neo4j.Driver) *LinkedonServer {
	return &LinkedonServer{
		driver: driver,
	}
}

func (s *LinkedonServer) CreatePerson(ctx context.Context, req *linkedon.CreatePersonRequest) (*linkedon.Response, error) {
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
	fmt.Printf("Created: %v\n", result) // TODO

	return &linkedon.Response{Success: true, Extra: "Created"}, nil
}

func (s *LinkedonServer) FollowPerson(ctx context.Context, req *linkedon.FollowPersonRequest) (*linkedon.Response, error) {
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
	MERGE (follower)-[r:FOLLOWS]->(followee)
	RETURN r
	`, map[string]any{
			"followerId": followerId,
			"followeeId": followeeId,
		})
	})
	if err != nil {
		log.Printf("failed to follow: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "failed to follow: %v", err)
	}
	fmt.Printf("Followed: %v\n", result) // TODO
	return &linkedon.Response{Success: true, Extra: "Followed"}, nil
}

func (s *LinkedonServer) Reset(ctx context.Context, req *linkedon.ResetRequest) (*linkedon.Response, error) {
	session := s.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: NEO4JDATABASE, AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)
	result, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		return tx.Run(ctx, `
		MATCH (n) DETACH DELETE n
	`, map[string]any{})
	})
	if err != nil {
		log.Printf("failed to reset: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to reset: %v", err)
	}
	fmt.Printf("Followed: %v\n", result) // TODO
	return &linkedon.Response{Success: true, Extra: "Followed"}, nil
}

