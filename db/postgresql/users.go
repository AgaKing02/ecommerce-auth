package postgresql

import (
	"context"
	"database/sql"
	"ecommerce-auth/models"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	insertSql                    = "INSERT INTO users (username,password,name,age) VALUES ($1,$2,$3,$4) RETURNING id"
	getUserById                  = "SELECT id,username,password,name,age FROM users where id=$1 LIMIT 1"
	getUserByUsernameAndPassword = "SELECT id,username,password,name,age FROM users where username=$1 and password=$2 LIMIT 1"
)

type UserModel struct {
	Pool *pgxpool.Pool
}

func (m *UserModel) Insert(username, password, name string, age int) (int, error) {
	var id uint64
	row := m.Pool.QueryRow(context.Background(), insertSql, username, password, name, age)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
func (m *UserModel) Get(id int) (*models.User, error) {
	s := &models.User{}
	err := m.Pool.QueryRow(context.Background(), getUserById, id).Scan(&s.ID, &s.Username, &s.Password, &s.Name, &s.Age)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("No such user")
		} else {
			return nil, err
		}
	}
	return s, nil
}
func (m *UserModel) GetByAuthToken(token *models.AuthToken) (*models.User, error) {
	s := &models.User{}
	err := m.Pool.QueryRow(context.Background(), getUserByUsernameAndPassword, token.Username, token.Password).Scan(&s.ID, &s.Username, &s.Password, &s.Name, &s.Age)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("No such user")
		} else {
			return nil, err
		}
	}
	return s, nil

}
