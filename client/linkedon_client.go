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
	ctx := context.Background()
	client := linkedon.NewLinkedonServiceClient(conn)
	person1 := &linkedon.Person{
		Name: "person 1",
		Age:  20,
		Id:   101,
	}
	res, err := client.CreatePerson(ctx, &linkedon.CreatePersonRequest{Person: person1})
	if err != nil {
		log.Fatalf("failed to create person: %v", err)
		panic(err)
	}
	log.Printf("created person: %v", res)
	person2 := &linkedon.Person{
		Name: "person 2",
		Age:  21,
		Id:   102,
	}
	res, err = client.CreatePerson(ctx, &linkedon.CreatePersonRequest{Person: person2})
	if err != nil {
		log.Fatalf("failed to create person: %v", err)
		panic(err)
	}
	log.Printf("created person: %v", res)
	followRes, err := client.Follow(ctx, &linkedon.FollowRequest{FollowerId: 101, FolloweeId: 102})
	if err != nil {
		log.Fatalf("failed to follow: %v", err)
		panic(err)
	}
	log.Printf("followed: %v", followRes)
}
