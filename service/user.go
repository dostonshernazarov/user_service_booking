package service

import (
	"context"
	"github.com/jmoiron/sqlx"
	pb "user_service_booking/genproto/user_proto"
	l "user_service_booking/pkg/logger"
	"user_service_booking/storage"
)

// UserService ...
type UserService struct {
	storage storage.IStorage
	logger  l.Logger
}

// NewUserService ...
func NewUserService(db *sqlx.DB, log l.Logger) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *UserService) GetWithColumnAndItem(ctx context.Context, req *pb.GetWithColumnAndItemReq) (*pb.GetAllUsersRespons, error) {
	users, err := s.storage.User().GetWithColumnAndItem(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return &pb.GetAllUsersRespons{
		User: users,
	}, nil
}

func (s *UserService) DeleteUserByID(ctx context.Context, req *pb.IdRequest) (*pb.DeleteUserByIDRespons, error) {

	err := s.storage.User().SoftDeleteUserByID(req.Id)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return &pb.DeleteUserByIDRespons{
		Result: "User deleted",
	}, nil
}

func (s *UserService) GetAllUsers(ctx context.Context, req *pb.GetAllUsersRequest) (*pb.GetAllUsersRespons, error) {
	users, err := s.storage.User().GetAllUsers(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return &pb.GetAllUsersRespons{
		User: users,
	}, nil
}

func (s *UserService) GetUserByID(ctx context.Context, req *pb.IdRequest) (*pb.User, error) {

	res, err := s.storage.User().GetUserByID(req.Id)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return res, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.User) (*pb.User, error) {

	res, err := s.storage.User().UpdateUserByID(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return res, nil

}

func (s *UserService) GetUserByRfshToken(ctx context.Context, req *pb.GetUserByRfshTokenRequest) (*pb.User, error) {
	res, err := s.storage.User().GetUserByRefreshTkn(req.Token)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return res, nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, request *pb.GetUserByEmailRequest) (*pb.User, error) {
	res, err := s.storage.User().GetUserbyEmail(request.Email)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return res, nil
}

func (s *UserService) Create(ctx context.Context, req *pb.User) (*pb.User, error) {
	res, err := s.storage.User().Create(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return res, nil
}

func (s *UserService) CheckUniqueEmail(ctx context.Context, req *pb.CheckUniqueRequest) (*pb.CheckUniqueRespons, error) {
	resp, err := s.storage.User().CheckUniqueEmail(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return resp, nil
}

func (s *UserService) CheckUniqueNum(ctx context.Context, req *pb.CheckUniqueRequest) (*pb.CheckUniqueRespons, error) {
	resp, err := s.storage.User().CheckUniqueNum(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return resp, nil
}
