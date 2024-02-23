package repo

import (
	pb "user_service_booking/genproto/user_proto"
)

// UserStorageI ...
type UserStorageI interface {
	UpdateUserByID(*pb.User) (*pb.User, error)
	GetUserbyEmail(string) (*pb.User, error)
	GetUserByRefreshTkn(string) (*pb.User, error)
	Create(*pb.User) (*pb.User, error)
	CheckUniqueEmail(*pb.CheckUniqueRequest) (*pb.CheckUniqueRespons, error)
	CheckUniqueNum(*pb.CheckUniqueRequest) (*pb.CheckUniqueRespons, error)
	GetUserByID(string) (*pb.User, error)
	GetAllUsers(*pb.GetAllUsersRequest) ([]*pb.User, error)
	SoftDeleteUserByID(string) error
	GetWithColumnAndItem(*pb.GetWithColumnAndItemReq) ([]*pb.User, error)
}
