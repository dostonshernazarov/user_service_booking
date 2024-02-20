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

func (s *UserService) Create(ctx context.Context, req *pb.User) (*pb.User, error) {
	return nil, nil
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
