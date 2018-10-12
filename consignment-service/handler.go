package main

import (
	vesselProto "github.com/jneo8/go-micro/vessel-service/proto/vessel"
	pb "go-micro/consignment-service/proto/consignment"
	"golang.org/x/net/context"
	"log"
)

type handler struct {
	vesselClient vesselProto.VesselService
}

func (s *handler) getRepo() Repository {
	return &ConsignmentRepository{s.session.Clone()}
}

// createConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as
// an argument, these are handled by the gRPC server.

func (s *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	repo := s.getRepo()
	defer repo.Close()
	// Here we call a client instance of our vessel service with consignment weight,
	// and the amount of containers as the capacity value
	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	if err != nil {
		return err
	}

	// We set the VesselId as the vessel we got back from our vessel service
	req.VesslId = vesselResponse.Vessel.Id

	// save our consignment
	err = repo.Create(req)
	if err != nil {
		return err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	res.Created = true
	res.Consignment = req
	return nil
}

func (s *handler) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()
	consignments, err := repo.getAll()
	if err != nil {
		return err
	}

	res.Consignments = consignments
	return nil
}
