package sqllite

import (
	"database/sql"
	"user-service/models"
	"time"
)

type SqlLiteDatasource struct {
	db        *sql.DB
}

func NewSqlLiteDatasource(db *sql.DB) *SqlLiteDatasource {
    return &SqlLiteDatasource{db}
}

func (r *SqlLiteDatasource) InsertUser(user models.User) *models.User {
	statement, _ := r.db.Prepare("INSERT INTO users (user_address, created) VALUES (?, ?)")
    statement.Exec(user.UserAddress, time.Now())
    rows, _ := r.db.Query("SELECT user_id FROM users")
    var id int
    for rows.Next() {
        rows.Scan(&id)
    }

	defer statement.Close()
	defer rows.Close()

	return &models.User{UserId : id}
}