package core

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type DBAccess struct {
	Name string `env:"DB_NAME,notEmpty"`
	Host string `env:"DB_HOST,notEmpty"`
	Port uint   `env:"DB_PORT,notEmpty"`
	User string `env:"DB_USER,notEmpty"`
	Pass string `env:"DB_PASSWORD,notEmpty"`
}

func NewDBAccess(
	name string,
	host string,
	port uint,
	user string,
	pass string,
) DBAccess {
	return DBAccess{
		Name: name,
		Host: host,
		Port: port,
		User: user,
		Pass: pass,
	}
}

func (dbAccess DBAccess) CreateConnection() (*sql.DB, error) {
	return sql.Open("pgx", fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbAccess.Host,
		dbAccess.Port,
		dbAccess.User,
		dbAccess.Pass,
		dbAccess.Name,
	))
}
