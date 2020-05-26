package database

import (
	"2020_1_Color_noise/internal/pkg/config"
	"errors"
	"fmt"
	"github.com/jackc/pgx"
	"log"
	"time"
)

type PgxDB struct {
	dbPool *pgx.ConnPool
}

func NewPgxDB() *PgxDB {
	return &PgxDB{}
}

func (db *PgxDB) Open(con config.DataBaseConfig) (err error) {

	connConfig := pgx.ConnConfig{
		Host:     con.Host,
		Port:     uint16(con.Port),
		Database: con.Database,
		User:     con.User,
		Password: con.Password,
	}

	poolConfig := pgx.ConnPoolConfig{
		ConnConfig:     connConfig,
		MaxConnections: 1,
		AcquireTimeout: 10 * time.Second,
		AfterConnect:   nil,
	}

	if db.dbPool != nil {
		return errors.New("pool was created already")
	}

	db.dbPool, err = pgx.NewConnPool(poolConfig)
	if err != nil {
		log.Print(err)
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
