package infra

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func init() {
	pgDriver = PostgresDriver{
		Host:     Host,
		Port:     Port,
		User:     User,
		Password: Password,
		DbName:   DbName,
	}

}

const (
	driverName = "postgres"
	User       = "postgres"
	Host       = "localhost"
	Port       = 5432
	Password   = "postgres"

	DbName = "allanfs"
)

var (
	pgDriver PostgresDriver
)

type PostgresDriver struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

// PostgrersDriver String method to return the connection string
func (p *PostgresDriver) String() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", p.Host, p.Port, p.User, p.Password, p.DbName)
}

func StartPostgres() (*sql.DB, error) {
	db, err := sql.Open(driverName, pgDriver.String())

	if err != nil {
		return nil, NewEssentialError(err, "Postgres", "Error open with Postgres")
	}

	err = db.Ping()
	if err != nil {
		return nil, NewEssentialError(err, "Postgres", "Error pinging Postgres")
	}

	return db, nil
}
