package database

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/config"
	"database/sql"
	_ "github.com/lib/pq"
	"time"
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

//TODO:создание пина
func (db *PostgresSQL) CreatePin(pin models.DataBasePin) *sql.Row {
	return db.QueryRow("INSERT INTO pins(user_id, name, description, image, board_id, created_at) " +
		"VALUES($1, $2, $3, $4, $5, $6) RETURNING id",
		pin.UserId, pin.Name, pin.Description, pin.Image, pin.BoardId, time.Now())
}

//TODO:обновление пина
func (db *PostgresSQL) UpdatePin(pin models.DataBasePin) (sql.Result, error) {
	return db.Exec("UPDATE pins SET" +
	"name = $1, description = $2, board_id = $3" +
	"WHERE id = &4", pin.Name, pin.Description, pin.BoardId, pin.Id)
}

func (db *PostgresSQL) DeletePin(pin models.DataBasePin) (sql.Result, error) {
	return db.Exec("DELETE from pins WHERE id = $1", pin.Id)
}

func (db *PostgresSQL) GetPinById(id int) *sql.Row {
	return db.QueryRow("SELECT * FROM pins WHERE id = $1", id)
}

func (db *PostgresSQL) GetPinsByUserId(id int) (*sql.Rows, error) {
	return db.Query("SELECT * FROM pins WHERE user_id = $1", id)
}

//TODO: заменить на полнотекстовый поиск
func (db *PostgresSQL) GetPinByName(name string) (*sql.Rows, error) {
	return db.Query("SELECT * FROM pins WHERE name = $1", name)
}