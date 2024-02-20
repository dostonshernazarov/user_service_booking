package postgres

import (
	"github.com/jmoiron/sqlx"
	pb "user_service_booking/genproto/user_proto"
	"user_service_booking/pkg/logger"
)

type userRepo struct {
	db *sqlx.DB
}

// NewUserRepo ...
func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *pb.User) (*pb.User, error) {
	return nil, nil
}

func (r *userRepo) CheckUniqueEmail(req *pb.CheckUniqueRequest) (*pb.CheckUniqueRespons, error) {
	query := `SELECT count(1) FROM user_info WHERE email=$1 AND deleted_at IS NULL`

	var result int

	err := r.db.QueryRow(query, req.Value).Scan(&result)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	if result == 0 {
		return &pb.CheckUniqueRespons{
			IsExist: false,
		}, nil
	}

	return &pb.CheckUniqueRespons{IsExist: true}, nil
}

func (r *userRepo) CheckUniqueNum(req *pb.CheckUniqueRequest) (*pb.CheckUniqueRespons, error) {
	query := `SELECT count(1) FROM user_info WHERE phone=$1 AND deleted_at IS NULL`

	var result int

	err := r.db.QueryRow(query, req.Value).Scan(&result)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	if result == 0 {
		return &pb.CheckUniqueRespons{
			IsExist: false,
		}, nil
	}

	return &pb.CheckUniqueRespons{IsExist: true}, nil
}
