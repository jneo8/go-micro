// consignment-service/main.go
package main

import (
	"fmt"
	"log"
	"os"

	// Import the generated protobuf code
	vesselProto "github.com/jneo8/go-micro/vessel-service/proto/vessel"
	micro "github.com/micro/go-micro"
	pb "go-micro/consignment-service/proto/consignment"
)

const (
	defaultHost = "localhost:27017"
)

func main() {
	// Database host from the environment variables
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)

	// Mgo create a 'Master' session, we need to end that session
	// before the main function closes.
	defer session.Close()

	if err != nil {
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselService("go.micro.srv.vessel", srv.Client())

	// Init will parse the command line flags.
	srv.Init()

	// Register handler
	pb.RegisterShippingServiceHandler(srv.Server(), &service{session, vesselClient})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

// type Repository interface {
// 	Create(*pb.Consignment) (*pb.Consignment, error)
// 	GetAll() []*pb.Consignment
// }
//
// // Repository - Dummy repository, this simulates the use of a datastore
// // of some kind. We'll replace this with a real implementation later on.
// type ConsignmentRepository struct {
// 	consignments []*pb.Consignment
// }
//
// func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
// 	updated := append(repo.consignments, consignment)
// 	repo.consignments = updated
// 	return consignment, nil
// }
//
// func (repo *ConsignmentRepository) GetAll() []*pb.Consignment {
// 	return repo.consignments
// }
//
// // Service should implement all of the methods to satisfy the service
// // we defined in our protobuf definition. You can check the interface
// // in the generated code itself for the exact method signatures etc
// // to give you a better idea.
// type service struct {
// 	repo         Repository
// 	vesselClient vesselProto.VesselService
// }
//
// // CreateConsignment - we created just one method on our service,
// // which is a create method, which takes a context and a request as an
// // argument, these are handled by the gRPC server.
// func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
//
// 	// Here we call a client instance of our vessel service with our consignment weight,
// 	// and the amount of containers as teh capacity value
// 	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
// 		MaxWeight: req.Weight,
// 		Capacity:  int32(len(req.Containers)),
// 	})
//
// 	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
//
// 	if err != nil {
// 		return err
// 	}
//
// 	// We set VesselId as the vessel Weight got back from our vessel service.
// 	req.VesselId = vesselResponse.Vessel.Id
//
// 	// Save our consignment
// 	consignment, err := s.repo.Create(req)
// 	if err != nil {
// 		return err
// 	}
//
// 	// Return matching the `Response` message we created in our
// 	// protobuf definition.
// 	res.Created = true
// 	res.Consignment = consignment
// 	return nil
// }
//
// func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
// 	consignments := s.repo.GetAll()
// 	res.Consignments = consignments
// 	return nil
// }
//
// func main() {
//
// 	repo := &ConsignmentRepository{}
//
// 	// Create a new service. Optionally include some options here.
// 	srv := micro.NewService(
// 		// This name must match the package name given in your protobuf definition
// 		micro.Name("go.micro.srv.consignment"),
// 		micro.Version("latest"),
// 	)
//
// 	vesselClient := vesselProto.NewVesselService("go.micro.srv.vessel", srv.Client())
//
// 	srv.Init()
//
// 	// Register handler
// 	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo, vesselClient})
//
// 	// Run the server
// 	if err := srv.Run(); err != nil {
// 		fmt.Println(err)
// 	}
// }
