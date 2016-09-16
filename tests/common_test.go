package test

import (
	"fmt"
	"os"

	"database/sql"
)

const (
	xprotocolDriver  = "mysql/xprotocol"
	dsnCharsetSuffix = "?xprotocol=1&charset=utf8mb4"
	defaultDsnUser   = "xprotocol_user"
	defaultDsnPass   = "xprotocol_pass"
	defaultDsnHost   = "127.0.0.1"
	defaultDsnPort   = "33060"
	dsnSuffix        = "?xprotocol=1"
	defaultDB        = "xprotocol_test"
)

var (
	defaultDSN string
	charsetDSN string
)

func init() {
	defaultDSN = dsnPrefix() + defaultDB + dsnSuffix
	charsetDSN = dsnPrefix() + defaultDB + dsnCharsetSuffix
}

// choose from the environment variable if it exists and is not empty
// - was hard-coded defaultDsnPrefix = "xprotocol_user:xprotocol_pass@tcp(127.0.0.1:33060)/"
func choose(envname, defaultValue string) string {
	if envValue := os.Getenv(envname); envValue != "" {
		return envValue
	}
	return defaultValue
}

// return the default prefix. We can now use environment variables to override the defaults.
func dsnPrefix() string {
	return defaultDsnUser +
		":" +
		defaultDsnPass +
		"@(" +
		choose("MYSQL_HOST", defaultDsnHost) + ":" +
		choose("MYSQLX_PORT", defaultDsnPort) + ")/"
}

type testDSN struct {
	dsn    string
	result bool
}
type testDSNs []testDSN

// open the connection to a standard port and provide the name of the test to simplify logging
func defaultOpen(name string) (*sql.DB, error) {
	db, err := sql.Open(xprotocolDriver, defaultDSN)
	if err != nil {
		err = fmt.Errorf("connect(%q) db.Open(%q,%q) failed: %v", name, xprotocolDriver, defaultDSN, err)
	}
	return db, err
}

// open the connection to a standard port and provide the name of the test to simplify logging
func charsetOpen(name string) (*sql.DB, error) {
	db, err := sql.Open(xprotocolDriver, charsetDSN)
	if err != nil {
		err = fmt.Errorf("connect(%q) db.Open(%q,%q) failed: %v", name, xprotocolDriver, charsetDSN, err)
	}
	return db, err
}

// closing the connection
func defaultClose(name string, db *sql.DB) error {
	err := db.Close()
	if err != nil {
		err = fmt.Errorf("defaultClose(%q,db) db.Close() failed: %v", name, err)
	}
	return err
}

// close rows
func defaultCloseRows(name string, rows *sql.Rows) error {
	err := rows.Close()
	if err != nil {
		err = fmt.Errorf("defaultCloseRows(%q,rows) rows.Close() failed: %v", name, err)
	}
	return err
}
