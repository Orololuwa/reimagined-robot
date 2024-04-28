package repository

import (
	"database/sql"
	"errors"

	"github.com/orololuwa/reimagined-robot/models"
)

type testUserDBRepo struct {
	DB *sql.DB
}

func NewUserTestingDBRepo() UserRepo {
	return &testUserDBRepo{
	}
}

func (m *testUserDBRepo) CreateAUser(user models.User) (int, error){
	var newId int

	if user.Password == "invalid"{
		return newId, errors.New("CreateAUser: DB repo fail")
	}

	return newId, nil
}

func (m *testUserDBRepo) GetAUser(id int)(models.User, error){
	var user models.User

	if id == 0{
		return user, errors.New("GetAUser: DB repo fail")
	}

	return user, nil
}