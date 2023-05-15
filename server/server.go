package server

import pb "school-manager/proto"

type Server struct {
	pb.UnimplementedSchoolManagerServiceServer
}
