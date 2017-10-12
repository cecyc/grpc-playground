package main

import (
	"log"

	pb "cecy/proto-fun/doggos/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	//Connect to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Printf("Error dialing: %s", err)
	}
	defer conn.Close()

	//Create a new client
	c := pb.NewDoggosClient(conn)

	newDoggo, err := c.CreateNewDoggo(context.Background(), &pb.NewDoggoRequest{Name: "Rusty", Age: 10})
	if err != nil {
		log.Printf("Error creating new Doggo: %s", err)
	}
	log.Print(newDoggo)

	getDoggo, err := c.GetDoggo(context.Background(), &pb.GetDoggoRequest{Name: "Rusty"})
	if err != nil {
		log.Printf("Error getting Doggo: %s", err)
	}
	log.Print(getDoggo)

	deleteDoggo, err := c.DeleteDoggo(context.Background(), &pb.DeleteDoggoRequest{Name: "Rusty"})
	log.Print(deleteDoggo)

	listDoggos, err := c.ListDoggos(context.Background(), &pb.ListDoggosRequest{})
	if err != nil {
		log.Printf("Error listing Doggos: %s", err)
	}
	log.Print(listDoggos)
}
