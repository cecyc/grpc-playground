package main

import (
	"log"
	"net"

	pb "cecy/proto-fun/doggos/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

//Empty server, this could have things needed for the server, like db connections
type server struct{}

//Implement all methods in defined in DoggosServer interface in the pb file
func (s *server) CreateNewDoggo(ctx context.Context, in *pb.NewDoggoRequest) (*pb.Doggo, error) {
	log.Printf("Creating new Doggo name %s and age %d", in.Name, in.Age)
	doggo := &pb.Doggo{Name: in.Name, Age: in.Age}
	return doggo, nil
}

//This should just return whatever it's given, but IRL this would get data from a DB
func (s *server) GetDoggo(ctx context.Context, in *pb.GetDoggoRequest) (*pb.Doggo, error) {
	log.Printf("Getting Doggo %s", in.Name)
	doggo := &pb.Doggo{Name: in.Name}
	return doggo, nil
}

func (s *server) DeleteDoggo(ctx context.Context, in *pb.DeleteDoggoRequest) (*pb.DeleteDoggoResponse, error) {
	log.Printf("Deleting Doggo %s", in.Name)
	return &pb.DeleteDoggoResponse{Success: true}, nil
}

func (s *server) ListDoggos(ctx context.Context, in *pb.ListDoggosRequest) (*pb.ListDoggosResponse, error) {
	log.Print("Listing Doggos")
	var d []*pb.Doggo
	d = append(d, &pb.Doggo{Name: "Rusty", Age: 10})
	d = append(d, &pb.Doggo{Name: "Baloo", Age: 7})
	return &pb.ListDoggosResponse{d}, nil
}

//Start the server and listen
func main() {
	log.Println("***Starting server***")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}

	s := grpc.NewServer()
	pb.RegisterDoggosServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
