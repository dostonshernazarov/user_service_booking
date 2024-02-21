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

func (s *UserService) PartCreate(ctx context.Context, req *pb.PartUser) (*pb.PartUser, error) {
	res, err := s.storage.User().PartCreate(req)
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
