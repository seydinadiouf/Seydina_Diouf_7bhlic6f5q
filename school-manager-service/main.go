package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	pb "school-manager/proto"
	"school-manager/school-manager-service/config"
	"school-manager/school-manager-service/server"
)

func main() {
	var addr string = "0.0.0.0:9090"

	// Connect To Database
	config.DatabaseInit()
	gorm := config.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	err = dbGorm.Ping()
	if err != nil {
		return
	}

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Printf("Listening at %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterSchoolManagerServiceServer(s, &server.Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}
