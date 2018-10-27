package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/micro/go-micro"
	pb "go-micro/vessel-service/proto/vessel"
)

func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}
