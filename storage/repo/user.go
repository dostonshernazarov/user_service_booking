package repo

import (
	pb "user_service_booking/genproto/user_proto"
)

// UserStorageI ...
type UserStorageI interface {
	GetUserbyEmail(string) (*pb.User, error)
	PartCreate(*pb.PartUser) (*pb.PartUser, error)
	Create(*pb.User) (*pb.User, error)
	CheckUniqueEmail(req *pb.CheckUniqueRequest) (*pb.CheckUniqueRespons, error)
	CheckUniqueNum(req *pb.CheckUniqueRequest) (*pb.CheckUniqueRespons, error)
	GetUserByRefreshTkn(string) (*pb.User, error)
	UpdateUserByID(*pb.User) (*pb.User, error)
}
