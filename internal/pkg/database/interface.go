package database

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/config"
	"database/sql"
)

type DBInterface interface {
	//соединиться
	Open(config config.DataBaseConfig) error
	//проверить
	Ping() error
	//закрыть соединение
	Close() error
	//выполнить что-то
	Exec(query string, args ...interface{}) (sql.Result, error)
	//получить выборку значений
	Query(query string, args ...interface{}) (*sql.Rows, error)
	//получить одно значение
	QueryRow(query string, args ...interface{}) *sql.Row
	//методы для пинов
	CreatePin(pin models.DataBasePin) *sql.Row
	UpdatePin(pin models.DataBasePin) (sql.Result, error)
	DeletePin(pin models.DataBasePin) (sql.Result, error)

	GetPinById(id int) *sql.Row
	GetPinsByUserId(id int) (*sql.Rows, error)
	GetPinByName(name string) (*sql.Rows, error)
}

