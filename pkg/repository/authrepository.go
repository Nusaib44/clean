package repository

import (
	"clean/pkg/domain"
	inter "clean/pkg/repository/interfaces"
	"database/sql"
)

type authDatabase struct {
	DB *sql.DB
}

func NewAuthRepository(DB *sql.DB) inter.AuthRepository {
	return &authDatabase{
		DB: DB,
	}
}

// ------------------register------------------------------------------------------------------
func (c *authDatabase) Register(user domain.Users) (int, error) {
	var User_id int
	query := `INSERT INTO users(first_name,last_name,email,gender,phone,password)VALUES($1,$2,$3,$4,$5,$6)RETURNING user_id;`
	err := c.DB.QueryRow(query,
		user.First_Name,
		user.Last_Name,
		user.Email,
		user.Gender,
		user.Phone,
		user.Password).Scan(&User_id)
	return User_id, err

}

// --------------------------------find user-----------------------------------------------------
func (c *authDatabase) FindUser(email string) (domain.UserResponse, error) {
	var user domain.UserResponse
	query := `SELECT user_id,first_name,last_name,email,gender,password,phone FROM users WHERE email=$1;`

	err := c.DB.QueryRow(query, email).Scan(&user.ID,
		&user.First_Name,
		&user.Last_Name,
		&user.Email,
		&user.Gender,
		&user.Password,
		&user.Phone,
	)

	return user, err
}
