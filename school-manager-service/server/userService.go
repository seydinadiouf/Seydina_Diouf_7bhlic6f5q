package server

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"school-manager/auth"
	pb "school-manager/proto"
	"school-manager/school-manager-service/config"
	"school-manager/school-manager-service/dto/payload"
	"school-manager/school-manager-service/model"
)

func (*Server) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInResponse, error) {
	log.Printf("SignIn was invoked with")
	db := config.DB()

	var user *model.User

	signIn := payload.SignIn{
		Username: in.Username,
		Password: in.Password,
	}

	message := "Username or password incorrect"

	if res := db.Where("username = ?", signIn.Username).First(&user); res.Error != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("User Not found error: %v", message),
		)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signIn.Password)); err != nil {
		// If the two passwords don't match, return a 401 status

		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Not authorized: %v", message),
		)
	}

	// Generate encoded token and send it as response.
	t, err := auth.GenerateJWT(signIn.Username)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Error when generating token"),
		)
	}

	return &pb.SignInResponse{Username: in.Username, Token: t}, nil

}
