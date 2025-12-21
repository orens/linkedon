package main

import (
	"context"
	"log"

	"github.com/orens/linkedon/linkedon"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:8044", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
		panic(err)
	}
	defer conn.Close()
	client := linkedon.NewLinkedonServiceClient(conn)
	person := &linkedon.Person{
		Name: "oren s",
		Age:  48,
		Id:   1,
	}
	res, err := client.CreatePerson(context.Background(), &linkedon.CreatePersonRequest{Person: person})
	if err != nil {
		log.Fatalf("failed to create person: %v", err)
		panic(err)
	}
	log.Printf("created person: %v", res)
}
