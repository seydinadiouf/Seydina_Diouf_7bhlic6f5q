package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"log"
	pb "school-manager/proto"
	"school-manager/school-manager-gateway/middlewares"
)

func main() {

	// Create a new Echo instance
	e := echo.New()

	// Add middleware to handle CORS
	e.Use(middleware.CORS())

	// Add token validation
	e.Use(middlewares.Auth())

	// Create a new gRPC-Gateway runtime
	gwmux := runtime.NewServeMux()

	// Dial gRPC server endpoint
	grpcEndpoint := "localhost:9090"
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := pb.RegisterSchoolManagerServiceHandlerFromEndpoint(context.Background(), gwmux, grpcEndpoint, opts); err != nil {
		log.Fatalf("failed to register endpoint: %v", err)
	}

	// Register the gRPC-Gateway runtime with Echo
	e.Any("*", func(c echo.Context) error {
		req := c.Request()
		resp := c.Response()
		gwmux.ServeHTTP(resp, req)
		return nil
	})

	// Start the Echo server
	httpAddr := ":8080"
	log.Printf("starting HTTP server on %s", httpAddr)
	if err := e.Start(httpAddr); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
