package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
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
	//defer r.db.Close()
	var userReq pb.User
	var firstName, lastName, email, password, birthday, imageUrl, cardNum, phone, role, refreshToken sql.NullString

	currentTime := time.Now()

	query := `UPDATE user_info SET first_name = $1, last_name = $2, email = $3,
                     password = $4, birthday = $5, image_url = $6,
                     card_num = $7, phone = $8, role = $9, 
                     refresh_token = $10, updated_at = $11 WHERE id = $12 RETURNING id,
                                         first_name, last_name, email,
                                         password, birthday, image_url, 
                                         card_num, phone, role, refresh_token`

	err := r.db.QueryRow(
		query,
		user.FirstName, user.LastName,
		user.Email, user.Password, user.Birthday,
		user.ImageUrl, user.CardNum, user.Phone,
		user.Role, user.RefreshToken, currentTime, user.Id).Scan(&userReq.Id, &firstName, &lastName,
		&email, &password, &birthday, &imageUrl,
		&cardNum, &phone, &role, &refreshToken)
	if err != nil {
		logger.Error(err)
		return nil, err

	}

	// Convert sql.NullString to regular strings if they are valid
	if firstName.Valid {
		userReq.FirstName = firstName.String
	}
	if lastName.Valid {
		userReq.LastName = lastName.String
	}
	if email.Valid {
		userReq.Email = email.String
	}
	if password.Valid {
		userReq.Password = password.String
	}
	if birthday.Valid {
		userReq.Birthday = birthday.String
	}
	if imageUrl.Valid {
		userReq.ImageUrl = imageUrl.String
	}
	if cardNum.Valid {
		userReq.CardNum = cardNum.String
	}
	if phone.Valid {
		userReq.Phone = phone.String
	}
	if role.Valid {
		userReq.Role = role.String
	}
	if refreshToken.Valid {
		userReq.RefreshToken = refreshToken.String
	}
	return &userReq, nil

}

func (r *userRepo) GetUserbyEmail(emailReq string) (*pb.User, error) {
	//defer func(db *sqlx.DB) {
	//	err := db.Close()
	//	if err != nil {
	//
	//	}
	//}(r.db)

	var user pb.User
	var firstName, lastName, email, password, birthday, imageUrl, cardNum, phone, role, refreshToken sql.NullString

	query := `SELECT id, 
	first_name, last_name, email, 
	password, birthday, image_url, card_num, 
	phone, role, refresh_token FROM user_info WHERE email = $1 AND deleted_at IS NULL`

	err := r.db.QueryRow(query, emailReq).Scan(&user.Id, &firstName, &lastName,
		&email, &password, &birthday, &imageUrl,
		&cardNum, &phone, &role, &refreshToken)

	if err != nil {
		logger.Error(err)
		return nil, err

	}
	// Convert sql.NullString to regular strings if they are valid
	if firstName.Valid {
		user.FirstName = firstName.String
	}
	if lastName.Valid {
		user.LastName = lastName.String
	}
	if email.Valid {
		user.Email = email.String
	}
	if password.Valid {
		user.Password = password.String
	}
	if birthday.Valid {
		user.Birthday = birthday.String
	}
	if imageUrl.Valid {
		user.ImageUrl = imageUrl.String
	}
	if cardNum.Valid {
		user.CardNum = cardNum.String
	}
	if phone.Valid {
		user.Phone = phone.String
	}
	if role.Valid {
		user.Role = role.String
	}
	if refreshToken.Valid {
		user.RefreshToken = refreshToken.String
	}
	return &user, nil
}

func (r *userRepo) GetUserByRefreshTkn(token string) (*pb.User, error) {
	//defer r.db.Close()

	var firstName, lastName, email, password, birthday, imageUrl, cardNum, phone, role, refreshToken sql.NullString
	var user pb.User
	query := `SELECT id, first_name, last_name, 
	email, password, birthday, image_url,
	card_num, phone, role, refresh_token FROM user_info WHERE refresh_token = $1 AND deleted_at IS NULL`

	err := r.db.QueryRow(query, token).Scan(&user.Id, &firstName, &lastName,
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
	//defer r.db.Close()
	var firstName, lastName, email, password, birthday, imageUrl, cardNum, phone, role, refreshToken sql.NullString

	var userRes pb.User

	query := `INSERT INTO user_info (id,
	first_name, last_name, email,
	password, birthday, image_url,
	card_num, phone, role,
	refresh_token) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id, first_name, last_name, email,
	password, birthday, image_url,
	card_num, phone, role,
	refresh_token`

	err := r.db.QueryRow(query, user.Id, sql.NullString{String: user.FirstName, Valid: user.FirstName != ""},
		sql.NullString{String: user.LastName, Valid: user.LastName != ""},
		sql.NullString{String: user.Email, Valid: user.Email != ""},
		sql.NullString{String: user.Password, Valid: user.Password != ""},
		sql.NullString{String: user.Birthday, Valid: user.Birthday != ""},
		sql.NullString{String: user.ImageUrl, Valid: user.ImageUrl != ""},
		sql.NullString{String: user.CardNum, Valid: user.CardNum != ""},
		sql.NullString{String: user.Phone, Valid: user.Phone != ""},
		sql.NullString{String: user.Role, Valid: user.Role != ""},
		sql.NullString{String: user.RefreshToken, Valid: user.RefreshToken != ""}).Scan(&userRes.Id,
		&firstName, &lastName, &email,
		&password, &birthday, &imageUrl,
		&cardNum, &phone, &role, &refreshToken)

	if err != nil {
		logger.Error(err)
		return nil, err

	}
	if firstName.Valid {
		userRes.FirstName = firstName.String
	}
	if lastName.Valid {
		userRes.LastName = lastName.String
	}
	if email.Valid {
		userRes.Email = email.String
	}
	if password.Valid {
		userRes.Password = password.String
	}
	if birthday.Valid {
		userRes.Birthday = birthday.String
	}
	if imageUrl.Valid {
		userRes.ImageUrl = imageUrl.String
	}
	if cardNum.Valid {
		userRes.CardNum = cardNum.String
	}
	if phone.Valid {
		userRes.Phone = phone.String
	}
	if role.Valid {
		userRes.Role = role.String
	}
	if refreshToken.Valid {
		userRes.RefreshToken = refreshToken.String
	}

	return &userRes, nil

}

func (r *userRepo) CheckUniqueEmail(req *pb.CheckUniqueRequest) (*pb.CheckUniqueRespons, error) {
	//defer r.db.Close()
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
	//defer r.db.Close()
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

func (r *userRepo) GetUserByID(id string) (*pb.User, error) {
	var user pb.User
	var firstName, lastName, email, password, birthday, imageUrl, cardNum, phone, role, refreshToken sql.NullString

	query := `SELECT id, first_name, last_name, 
	email, password, birthday, image_url,
	card_num, phone, role, refresh_token FROM user_info WHERE id = $1 AND deleted_at IS NULL`

	err := r.db.QueryRow(query, id).Scan(&user.Id, &firstName, &lastName,
		&email, &password, &birthday, &imageUrl,
		&cardNum, &phone, &role, &refreshToken)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	// Convert sql.NullString to regular strings if they are valid
	if firstName.Valid {
		user.FirstName = firstName.String
	}
	if lastName.Valid {
		user.LastName = lastName.String
	}
	if email.Valid {
		user.Email = email.String
	}
	if password.Valid {
		user.Password = password.String
	}
	if birthday.Valid {
		user.Birthday = birthday.String
	}
	if imageUrl.Valid {
		user.ImageUrl = imageUrl.String
	}
	if cardNum.Valid {
		user.CardNum = cardNum.String
	}
	if phone.Valid {
		user.Phone = phone.String
	}
	if role.Valid {
		user.Role = role.String
	}
	if refreshToken.Valid {
		user.RefreshToken = refreshToken.String
	}

	return &user, nil
}

func (r *userRepo) GetAllUsers(req *pb.GetAllUsersRequest) ([]*pb.User, error) {
	var users []*pb.User

	query := `
		SELECT id, first_name, last_name, email, password, birthday, image_url,
		card_num, phone, role, refresh_token
		FROM user_info
		WHERE deleted_at IS NULL LIMIT $1 OFFSET $2
	`

	offset := req.Limit * (req.Page - 1)
	rows, err := r.db.Query(query, req.Limit, offset)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user pb.User
		var firstName, lastName, email, password, birthday, imageUrl, cardNum, phone, role, refreshToken sql.NullString

		err := rows.Scan(&user.Id, &firstName, &lastName,
			&email, &password, &birthday, &imageUrl,
			&cardNum, &phone, &role, &refreshToken)
		if err != nil {
			logger.Error(err)
			return nil, err
		}

		user.FirstName = stringValue(firstName)
		user.LastName = stringValue(lastName)
		user.Email = stringValue(email)
		user.Password = stringValue(password)
		user.Birthday = stringValue(birthday)
		user.ImageUrl = stringValue(imageUrl)
		user.CardNum = stringValue(cardNum)
		user.Phone = stringValue(phone)
		user.Role = stringValue(role)
		user.RefreshToken = stringValue(refreshToken)

		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		logger.Error(err)
		return nil, err
	}

	return users, nil
}

func (r *userRepo) SoftDeleteUserByID(id string) error {
	query := "UPDATE user_info SET deleted_at = $1 WHERE id = $2"

	currentTime := time.Now()

	result, err := r.db.Exec(query, currentTime, id)
	if err != nil {
		logger.Error(err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error(err)
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no user found with the provided ID")
	}

	return nil
}

func (r *userRepo) GetWithColumnAndItem(req *pb.GetWithColumnAndItemReq) ([]*pb.User, error) {
	var users []*pb.User

	iteamQ := "%" + req.Item + "%"

	query := fmt.Sprintf("SELECT id, first_name, last_name, email, password, birthday, image_url,card_num, phone, role, refresh_token FROM user_info WHERE %s LIKE '%s' AND deleted_at IS NULL LIMIT $1 OFFSET $2", req.Column, iteamQ)

	offset := req.Limit * (req.Page - 1)
	rows, err := r.db.Query(query, req.Limit, offset)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user pb.User
		var firstName, lastName, email, password, birthday, imageUrl, cardNum, phone, role, refreshToken sql.NullString

		err := rows.Scan(&user.Id, &firstName, &lastName,
			&email, &password, &birthday, &imageUrl,
			&cardNum, &phone, &role, &refreshToken)
		if err != nil {
			logger.Error(err)
			return nil, err
		}

		user.FirstName = stringValue(firstName)
		user.LastName = stringValue(lastName)
		user.Email = stringValue(email)
		user.Password = stringValue(password)
		user.Birthday = stringValue(birthday)
		user.ImageUrl = stringValue(imageUrl)
		user.CardNum = stringValue(cardNum)
		user.Phone = stringValue(phone)
		user.Role = stringValue(role)
		user.RefreshToken = stringValue(refreshToken)

		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		logger.Error(err)
		return nil, err
	}

	return users, nil
}

// stringValue returns the string value of a sql.NullString, handling null values.
func stringValue(s sql.NullString) string {
	if s.Valid {
		return s.String
	}
	return ""
}
