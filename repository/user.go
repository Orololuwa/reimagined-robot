package repository

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/orololuwa/reimagined-robot/models"
)

type user struct {
    DB *sql.DB
}

func NewUserRepo(conn *sql.DB) UserRepo {
    return &user{
        DB: conn,
    }
}
func (m *user) CreateAUser(user models.User) (int, error){
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    var newId int

    query := `
            INSERT into users 
                (first_name, last_name, email, password, created_at, updated_at)
            values 
                ($1, $2, $3, $4, $5, $6)
            returning id`

    err := m.DB.QueryRowContext(ctx, query, 
        user.FirstName, 
        user.LastName, 
        user.Email, 
        user.Password,
        time.Now(),
        time.Now(),
    ).Scan(&newId)

    if err != nil {
        log.Println(err.Error())
        return 0, err
    }

    return newId, nil
}


func (m *user) GetAUser(id int) (models.User, error){
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    log.Println("open connections", m.DB.Stats().OpenConnections)
    log.Println("open connections", m.DB.Stats().WaitDuration.Microseconds())

    var user models.User

    query := `
        SELECT id, first_name, last_name, email, password, created_at, updated_at
        from users
        WHERE
        id=$1
    `

    err := m.DB.QueryRowContext(ctx, query, id).Scan(
        &user.ID,
        &user.FirstName,
        &user.LastName,
        &user.Email,
        &user.Password,
        &user.CreatedAt,
        &user.UpdatedAt,
    )

    if err != nil {
        log.Println(err.Error())
        return user, err
    }

    return user, nil
}