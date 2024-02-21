package postgres

import (
	"database/sql"
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

func (r *userRepo) UpdateUserByID(user *pb.User) (*pb.User, error) {

	defer r.db.Close()

	query := `UPDATE user_info SET first_name = $1, last_name = $2, email = $3,
                     password = $4, birthday = $5, image_url = $6,
                     card_num = $7, phone = $8, role = $9, 
                     refresh_token = $10 WHERE id = $11 RETURNING id,
                                         first_name, last_name, email,
                                         password, birthday, image_url, 
                                         card_num, phone, role, refresh_token`

	err := r.db.QueryRow(
		query,
		user.FirstName, user.LastName,
		user.Email, user.Password, user.Birthday,
		user.ImageUrl, user.CardNum, user.Phone,
		user.Role, user.RefreshToken, user.Id).Scan(
		&user.Id, &user.FirstName, &user.LastName,
		&user.Email, &user.Password, &user.Birthday,
		&user.ImageUrl, &user.CardNum, &user.Phone,
		&user.Role, &user.RefreshToken)
	if err != nil {
		logger.Error(err)
		return nil, err

	}
	return user, nil

}

func (r *userRepo) GetUserbyEmail(email string) (*pb.User, error) {

	defer r.db.Close()
	var user pb.User
	query := `SELECT id, 
	first_name, 
	last_name, 
	email, 
	password FROM user_info WHERE email = $1 AND deleted_at IS NULL`

	err := r.db.QueryRow(query, email).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password)

	if err != nil {
		logger.Error(err)
		return nil, err

	}
	return &user, nil
}

func (r *userRepo) PartCreate(user *pb.PartUser) (*pb.PartUser, error) {
	defer r.db.Close()

	var partUser pb.PartUser

	query := `INSERT INTO user_info (
                       id,first_name,last_name,
                       email,password, refresh_token) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id,
                                 first_name, 
                                 last_name,
                                 email, refresh_token`

	err := r.db.QueryRow(query,
		user.Id,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password, user.RefreshTkn).Scan(&partUser.Id,
		&partUser.FirstName,
		&partUser.LastName,
		&partUser.Email, &partUser.RefreshTkn)

	if err != nil {
		logger.Error(err)
		return nil, err

	}
	return &partUser, nil
}

func (r *userRepo) GetUserByRefreshTkn(tokin string) (*pb.User, error) {
	defer r.db.Close()

	var firstName, lastName, email, password, birthday, imageUrl, cardNum, phone, role, refreshToken sql.NullString
	var user pb.User
	query := `SELECT id, first_name, last_name, 
	email, password, birthday, image_url,
	card_num, phone, role, refresh_token FROM user_info WHERE refresh_token = $1 AND deleted_at IS NULL`

	err := r.db.QueryRow(query, tokin).Scan(&user.Id, &firstName, &lastName,
		&email, &password, &birthday, &imageUrl,
		&cardNum, &phone, &role, &refreshToken)
	if err != nil {
		logger.Error(err)
		return nil, err

	}

	switch {
	case firstName.Valid:
		user.FirstName = firstName.String
	case lastName.Valid:
		user.LastName = lastName.String
	case email.Valid:
		user.Email = email.String
	case password.Valid:
		user.Password = password.String
	case birthday.Valid:
		user.Birthday = birthday.String
	case imageUrl.Valid:
		user.ImageUrl = imageUrl.String
	case cardNum.Valid:
		user.CardNum = cardNum.String
	case phone.Valid:
		user.Phone = phone.String
	case role.Valid:
		user.Role = role.String
	case refreshToken.Valid:
		user.RefreshToken = refreshToken.String
	}
	return &user, nil
}

func (r *userRepo) Create(user *pb.User) (*pb.User, error) {
	defer r.db.Close()

	query := `INSERT INTO user_info (id, 
	first_name, last_name, email, 
	password, birthday, image_url, 
	card_num, phone, role,
	refresh_token,
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id,
	first_name, last_name, email, 
	password, birthday, image_url, 
	card_num, phone, role, 
	refresh_token`

	err := r.db.QueryRow(
		query,
		user.Id, user.FirstName, user.LastName,
		user.Email, user.Password, user.Birthday,
		user.ImageUrl, user.CardNum, user.Phone,
		user.Role, user.RefreshToken).Scan(
		&user.Id, &user.FirstName, &user.LastName,
		&user.Email, &user.Password, &user.Birthday,
		&user.ImageUrl, &user.CardNum, &user.Phone,
		&user.Role, &user.RefreshToken)
	if err != nil {
		logger.Error(err)
		return nil, err

	}
	return user, nil

}

func (r *userRepo) CheckUniqueEmail(req *pb.CheckUniqueRequest) (*pb.CheckUniqueRespons, error) {
	defer r.db.Close()
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
	defer r.db.Close()
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
