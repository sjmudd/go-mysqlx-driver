package test

import (
	"database/sql"
	"testing"

	"github.com/sjmudd/go-mysqlx-driver/debug"
)

// test getting back different types
func Test04SystemVariables(t *testing.T) {
	var (
		hostname         string
		maxAllowedPacket int
		version          string
	)

	dsn := dsnPrefix() + defaultDB + dsnSuffix
	db, err := sql.Open("mysql/xprotocol", dsn)
	if err != nil {
		t.Errorf("Test04SystemVariables: could not connect to dsn:%s: %v", dsn, err)
	}

	// test a system variable (we need it later)
	query := `SELECT @@max_allowed_packet, @@version, @@hostname`
	t.Logf("Test04SystemVariables: testing: %q", query)
	rows, err := db.Query(query)
	if err != nil {
		t.Fatal(err)
	}
	var counter int
	defer rows.Close()
	for rows.Next() {
		counter++
		if err := rows.Scan(&maxAllowedPacket, &version, &hostname); err != nil {
			t.Fatal(err)
		}
		t.Logf("processed row[%d]: max_allowed_packet: %d, version: %q, hostname: %q", counter, maxAllowedPacket, version, hostname)
		debug.Msg("processed row[%d]: max_allowed_packet: %d, version: %q, hostname: %q", counter, maxAllowedPacket, version, hostname)
	}
	if err := rows.Err(); err != nil {
		t.Fatal(err)
	}

	err = db.Close()
	if err != nil {
		t.Errorf("Test04SystemVariables: could not connect to dsn:%s: %v", dsn, err)
	}
}
