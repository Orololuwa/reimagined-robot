package repository

import "github.com/orololuwa/reimagined-robot/models"

type UserRepo interface {
    CreateAUser(user models.User) (int, error)
    GetAUser(id int) (models.User, error)
}