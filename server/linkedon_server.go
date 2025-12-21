package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/orens/linkedon/linkedon"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type linkedonServer struct {
	linkedon.UnimplementedLinkedonServiceServer
}

func (s *linkedonServer) CreatePerson(ctx context.Context, req *linkedon.CreatePersonRequest) (*linkedon.CreatePersonResponse, error) {
	person := req.GetPerson()
	if person == nil {
		return nil, status.Error(codes.InvalidArgument, "person is required")
	}
	// TODO add to DB
	log.Printf("Creating person: %v with id %v and age %v\n", person.Name, person.Id, person.Age)

	return &linkedon.CreatePersonResponse{Success: true}, nil
}

func (s *linkedonServer) Connect(ctx context.Context, req *linkedon.ConnectRequest) (*linkedon.ConnectResponse, error) {
	person1Id := req.GetPerson1Id()
	person2Id := req.GetPerson2Id()
	if person1Id == 0 || person2Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "person1Id and person2Id are required")
	}
	// TODO connect people in DB
	log.Printf("Connecting people with ids %v and %v\n", person1Id, person2Id)
	return &linkedon.ConnectResponse{Success: true}, nil
}

func serve() error {
	port := 8044
	log.Printf("Starting linkedon server on port %d", port)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}
	log.Printf("Listening on %v", lis.Addr())
	grpcServer := grpc.NewServer()
	linkedon.RegisterLinkedonServiceServer(grpcServer, &linkedonServer{})
	log.Println("Registered services")
	log.Println("Serving...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return err
	}
	log.Println("Server stopped")
	return nil
}

func main() {
	if err := serve(); err != nil {
		log.Fatalf("failed to serve: %v", err)
		panic(err)
	}
}

// func connectToNeo4j() {
// 	ctx := context.Background()
// 	// URI examples: "neo4j://localhost", "neo4j+s://xxx.databases.neo4j.io"
// 	dbUri := "neo4j://localhost"
// 	dbUser := "neo4j"
// 	dbPassword := "rootroot"
// 	driver, err := neo4j.NewDriver(
// 		dbUri,
// 		neo4j.BasicAuth(dbUser, dbPassword, ""))
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer driver.Close(ctx)

// 	err = driver.VerifyConnectivity(ctx)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Connection established.")
// 	result, err := neo4j.ExecuteQuery(ctx, driver, `
//     CREATE (a:Person {name: $name})
//     CREATE (b:Person {name: $friendName})
//     CREATE (a)-[:KNOWS]->(b)
//     `,
// 		map[string]any{
// 			"name":       "Alice",
// 			"friendName": "David",
// 		}, neo4j.EagerResultTransformer,
// 		neo4j.ExecuteQueryWithDatabase("neo4j"))
// 	if err != nil {
// 		panic(err)
// 	}

// 	summary := result.Summary
// 	fmt.Printf("Created %v nodes in %+v.\n",
// 		summary.Counters().NodesCreated(),
// 		summary.ResultAvailableAfter())
// }
