package repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofrs/uuid"
	"github.com/rafaelgfirmino/authion/exceptions"
	"github.com/rafaelgfirmino/authion/infra/store"
	"github.com/rafaelgfirmino/authion/user/domain"
	"strings"
)

type mysqlUserRepository struct {
	DB *sql.DB
}

func NewMysqlUserRepository() UserRepository {
	return &mysqlUserRepository{store.Mysql}
}

const (
	queryFindUserByID                = "select ID, Email from User where id = ?"
	queryFindUserByEmail             = "select Email from User WHERE Email=?"
	queryUserInsertNewUser           = "insert into User (Email, Enabled, Password, ConfirmationToken) values (?, ?, ?, ?)"
	queryUserSelectConfirmationToken = "select ID, ConfirmationToken FROM User WHERE ConfirmationToken=?"
	queryUserSelectPasswordByEmail   = "select ID, Password FROM User WHERE Email = ?;"
	queryUserUpdateConfirmationToken = "update User u set u.ConfirmationToken = '' where u.id =?"
	queryUserUpdateEnabled           = "update User u set u.Enabled = true where u.id =?"
)

func (mysql mysqlUserRepository) FindByID(id int64) (*domain.User, error) {
	row := mysql.DB.QueryRow(queryFindUserByID, id)
	var user domain.User

	switch err := row.Scan(&user.ID, &user.Email); err {
	case sql.ErrNoRows:
		return &user, exceptions.ErrorUserNotFound
	case nil:
		return &user, nil
	default:
		panic(err)
	}
}

func (mysql mysqlUserRepository) FindByEmail(email string) (*domain.User, error) {
	row := mysql.DB.QueryRow(queryFindUserByEmail, email)
	var user domain.User

	switch err := row.Scan(&user.Email); err {
	case sql.ErrNoRows:
		return &user, exceptions.ErrorUserNotFound
	case nil:
		return &user, nil
	default:
		panic(err)
	}
}

func (mysql mysqlUserRepository) ConfirmationToken(confirmationToken string) error {
	sqlStatement := queryUserSelectConfirmationToken
	row := mysql.DB.QueryRow(sqlStatement, confirmationToken)

	var user domain.User

	switch err := row.Scan(&user.ID, &user.ConfirmationToken); err {
	case sql.ErrNoRows:
		return exceptions.ErrorTokenNotFound
	case nil:
		tx, _ := mysql.DB.Begin()
		var err error
		deleteConfirmationToken, _ := tx.Prepare(queryUserUpdateConfirmationToken)
		enableUser, _ := tx.Prepare(queryUserUpdateEnabled)
		_, err = deleteConfirmationToken.Exec(user.ID)
		if err != nil {
			tx.Rollback()
			return exceptions.ErrorTryingDeleteToken
		}
		_, err = enableUser.Exec(user.ID)
		if err != nil {
			tx.Rollback()
			return exceptions.ErrorTryingEnableUser
		}
		tx.Commit()
		return nil
	default:
		panic(err)
	}
}

func (mysql mysqlUserRepository) RegisterNewUser(user *domain.User) (*domain.User, error) {
	password, err := user.GetPassword()
	if err != nil {
		return &domain.User{}, err
	}
	user.ConfirmationToken = uuid.Must(uuid.NewV4()).String()
	stmt, err := mysql.DB.Prepare(queryUserInsertNewUser)
	result, err := stmt.Exec(
		strings.ToLower(user.Email),
		user.Enabled,
		password,
		user.ConfirmationToken,
	)

	if err != nil {
		return &domain.User{
			ID:    user.ID,
			Email: user.Email,
		}, err
	}

	user.ID, err = result.LastInsertId()
	if err != nil {
		return user, err
	}

	return user, nil
}

func (mysql mysqlUserRepository) Authenticate(currentUser *domain.User) error {
	sqlStatement := queryUserSelectPasswordByEmail
	row := mysql.DB.QueryRow(sqlStatement, currentUser.Email)
	var user domain.User
	switch err := row.Scan(&user.ID, &user.Password); err {
	case sql.ErrNoRows:
		return exceptions.ErrorEmailNotFound
	case nil:
		if !currentUser.CheckPasswordHash(user.Password) {
			return exceptions.ErrorPasswordNotFound
		}
		return nil
	default:
		panic(err)
	}
}
