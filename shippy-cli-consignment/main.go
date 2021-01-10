// shippy/shippy-cli-consignment/main.go
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"context"

	"github.com/micro/micro/service"
	pb "github.com/xoronedge/shippy/shippy-service-consignment/proto/consignment"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	// Set up a connection to the server.
	///-------------------------------------------------------Without micro
	// conn, err := grpc.Dial(address, grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("Did not connect: %v", err)
	// }
	// defer conn.Close()
	// client := pb.NewShippingServiceClient(conn)

	///--------------------WITH MICROOO
	srv := service.New()
	srv.Init()

	client := pb.NewShippingService("shippy.consignment.service", srv.Client())
	// Contact the server and print out its response.
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	r, err = client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}

//----------------------------------------------------------
//----Every network have own mdsn so both be in same network to work
//service in container and container have own network
//hots cli is on other netwrok and
//micro work with mdsn which every network have its own
//----------------------------------------------------------
