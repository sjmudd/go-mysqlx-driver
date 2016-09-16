package test

import (
	"database/sql"
	"io"
	"testing"

	"github.com/sjmudd/go-mysqlx-driver/debug"
)

// Test01ConnectToDB ensures that if we give a database name we end up in that db.
func Test01ConnectToDB(t *testing.T) {
	var database []byte

	dbList := []string{defaultDB, ""}

	for i := range dbList {
		dbName := dbList[i]

		dsn := dsnPrefix() + dbName + dsnSuffix
		db, err := sql.Open(xprotocolDriver, dsn)
		if err != nil {
			t.Errorf("Test01ConnectToDB: could not connect to dsn:%s: %v", dsn, err)
		} else {
			t.Logf("OK: Open(%q,%q) succeeded", xprotocolDriver, dsn)
		}

		sql := `SELECT DATABASE()`
		rows, err := db.Query(sql)
		if err != nil {
			t.Fatal(err)
		}
		for rows.Next() {
			if err := rows.Scan(&database); err != nil {
				t.Fatal(err)
			}
			if string(database) != dbName {
				t.Errorf("Test01ConnectToDB: %q returned: %q, expected: %q", sql, database, dbName)
			} else {
				t.Logf("OK: connection to database %q matches result from SELECT DATABASE()", dbName)
			}
		}
		if err := rows.Err(); err != nil {
			t.Fatal(err.Error())
		}

		if err = defaultClose("Test01ConnectToDB", db); err != nil {
			t.Errorf(err.Error())
		}
	}
}

// Test02Ping checks we can connect and ping a database
func Test02Ping(t *testing.T) {
	db, err := defaultOpen("Test02Ping")
	if err != nil {
		t.Errorf(err.Error())
	}

	// force a connection to the server with Ping
	t.Logf("Test02Ping: test sql.Ping()")
	if err = db.Ping(); err != nil {
		t.Errorf("db.Ping() failed: %v", err)
	} else {
		t.Logf("OK: db.Ping() succeeded")
	}

	if err = defaultClose("Test02Ping", db); err != nil {
		t.Errorf(err.Error())
	}
}

// Test03SimpleQuery checks we can do a simple query
func Test03SimpleQuery(t *testing.T) {
	var (
		one   int
		two   string
		three string
	)
	db, err := defaultOpen("Test03SimpleQuery")
	if err != nil {
		t.Errorf("Test03SimpleQuery: connect() failed: %v", err)
	}

	// test simplest SELECT possible
	t.Logf("Test03SimpleQuery: test Query()")
	rows, err := db.Query("SELECT 1, DATABASE(), 'xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx'")
	if err != nil {
		t.Fatal(err)
	}
	for rows.Next() {
		t.Logf("Test03SimpleQuery: read a row")
		if err := rows.Scan(&one, &two, &three); err != nil {
			t.Fatal(err)
		}
		debug.Msg("processed row: one: %d, two: %q, three: %q", one, two, three)
		t.Logf("processed row: one: %d, two: %q, three: %q", one, two, three)
	}

	t.Logf("Checking for errors after reading all rows")
	if err := rows.Err(); err != nil && err != io.EOF {
		t.Fatal(err)
	} else {
		t.Logf("OK: no errors")
	}

	// close rows
	if err := rows.Close(); err != nil {
		t.Errorf(err.Error())
	}

	// closing the connection
	if err = defaultClose("Test03SimpleQuery", db); err != nil {
		t.Errorf(err.Error())
	}
}
