package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
	"github.com/orens/linkedon/linkedon"
	"github.com/orens/linkedon/server/linkedonserver"
	"google.golang.org/grpc"
)

const (
	NEO4JURI      = "neo4j://localhost"
	NEO4JUSER     = "neo4j"
	NEO4JPASSWORD = "rootroot"
)

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
	linkedon.RegisterLinkedonServiceServer(grpcServer, linkedonserver.NewLinkedonServer(driver))
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
	session := driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: linkedonserver.NEO4JDATABASE, AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)
	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		return tx.Run(ctx, `
		CREATE CONSTRAINT person_id IF NOT EXISTS FOR (p:Person) REQUIRE p.id IS UNIQUE
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
