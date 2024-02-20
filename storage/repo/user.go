package repo

import (
	pb "user_service_booking/genproto/user_proto"
)

// UserStorageI ...
type UserStorageI interface {
	Create(*pb.User) (*pb.User, error)
	CheckUniqueEmail(req *pb.CheckUniqueRequest) (*pb.CheckUniqueRespons, error)
	CheckUniqueNum(req *pb.CheckUniqueRequest) (*pb.CheckUniqueRespons, error)
}
