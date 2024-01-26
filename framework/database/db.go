package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Db         *sql.DB
	Dsn        string
	DsnTest    string
	DbType     string
	DbTypeTest string
	Env        string
}

func NewDb() *Database {
	return &Database{}
}

func NewDbTest() *sql.DB {
	dbInstace := NewDb()
	dbInstace.Env = "test"
	dbInstace.DbTypeTest = "sqlite3"
	dbInstace.DsnTest = "./test.db"

	conn, err := dbInstace.Connect()

	if err != nil {
		panic(err)
	}

	return conn
}

func (d *Database) Connect() (*sql.DB, error) {

	var err error

	if d.Env != "test" {
		d.Db, err = sql.Open(d.DbType, d.Dsn)
	} else {
		d.Db, err = sql.Open(d.DbTypeTest, d.DsnTest)
	}

	if err != nil {
		return nil, err
	}
	return d.Db, nil
}
