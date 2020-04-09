package database

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx"
	"time"
)

type PgxDB struct {
	dbPool *pgx.ConnPool
}

func NewPgxDB() *PgxDB {
	return &PgxDB{}
}

func (db *PgxDB) Open() (err error) {
	//TODO: брать из конфига
	port := 5432

	connConfig := pgx.ConnConfig{
		Host:     "127.0.0.1",
		Port:     uint16(port),
		Database: "pinterest",
		User:     "postgres",
		Password: "password",
	}

	poolConfig := pgx.ConnPoolConfig{
		ConnConfig:     connConfig,
		MaxConnections: 50,
		AcquireTimeout: 10 * time.Second,
		AfterConnect:   nil,
	}

	if db.dbPool != nil {
		return errors.New("pool was created already")
	}

	db.dbPool, err = pgx.NewConnPool(poolConfig)
	if err != nil {
		return fmt.Errorf("connection is not established")
	}

	return nil
}

func (db *PgxDB) Close() error {
	db.dbPool.Close()
	return nil
}

func (db *PgxDB) Ping() error {
	if db.dbPool != nil {
		return nil
	}
	return fmt.Errorf("pool is not created")
}
