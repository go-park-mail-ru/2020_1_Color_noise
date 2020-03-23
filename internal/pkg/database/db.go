package database

import (
	"2020_1_Color_noise/internal/pkg/config"
	"database/sql"
	_ "github.com/lib/pq"
)

//THIS IS FOR POSTGRES ONLY

type PostgresSQL struct {
	DB *sql.DB
}

func (db *PostgresSQL) Open(config config.DataBaseConfig)  error{
	var (
		err error
		database *sql.DB
		connection string
	)

	connection = config.ConnString

	if database, err = sql.Open("postgres", connection); err != nil {
		return err
	}

	db.DB = database
	db.DB.SetMaxOpenConns(config.MaxConns)
	//проверяем подключение
	err = db.Ping()
	return err
}

func (db *PostgresSQL) Ping()  error{
	return db.DB.Ping()
}

func (db *PostgresSQL) Close()  error{
	return db.DB.Close()
}

func (db *PostgresSQL) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.DB.Exec(query, args...)
}

func (db *PostgresSQL) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.DB.Query(query, args...)
}

func (db *PostgresSQL) QueryRow(query string, args ...interface{}) *sql.Row {
	return db.DB.QueryRow(query, args...)
}