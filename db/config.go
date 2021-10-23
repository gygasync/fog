package db

import (
	"database/sql"
	"fmt"
	"fog/common"

	_ "github.com/mattn/go-sqlite3"
)

type DbConn struct {
	conn   *sql.DB
	logger *common.StdLogger
}

type Config interface {
	Up()
}

func NewDbConn(connection string, lgr *common.StdLogger) *DbConn {
	if connection == "" {
		return nil
	}
	dbc, err := open(connection, lgr)
	if err != nil {
		return nil
	}
	return &DbConn{conn: dbc, logger: lgr}
}

func open(connection string, logger *common.StdLogger) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", connection)
	l := *logger

	if err != nil {
		l.Error(fmt.Sprintf("%v", err))
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		l.Error(fmt.Sprintf("Error pinging DB: %v \n", err))
		return nil, err
	}

	l.Info("Connected to SQLite db!")

	return db, nil
}

func (conn *DbConn) Up() {
	up(conn.conn, conn.logger)
}
