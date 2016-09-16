// Go driver for MySQL X Protocol
//
// Copyright 2016 Simon J Mudd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package test

import (
	"io"
	"testing"

	"github.com/sjmudd/go-mysqlx-driver/debug"
)

/*
output seen on my laptop

root@localhost [(none)]> show global variables like '%char%';
+--------------------------+---------------------------------------------------------+
| Variable_name            | Value                                                   |
+--------------------------+---------------------------------------------------------+
| character_set_client     | latin1                                                  |
| character_set_connection | latin1                                                  |
| character_set_database   | latin1                                                  |
| character_set_filesystem | binary                                                  |
| character_set_results    | latin1                                                  |
| character_set_server     | latin1                                                  |
| character_set_system     | utf8                                                    |
| character_sets_dir       | /usr/local/mysql-5.7.14-osx10.11-x86_64/share/charsets/ |
+--------------------------+---------------------------------------------------------+
8 rows in set (0.01 sec)
*/

// change the dsn to use ... /<db>?charset=utf8mb4

// Test04CharsetConnection
func Test04CharsetConnection(t *testing.T) {
	const myName = "Test04CharsetConnection"
	var one int

	db, err := charsetOpen(myName)
	if err != nil {
		t.Errorf("%s: charsetOpen failed: %v", myName, err)
	}

	// Run simple SELECT query
	t.Logf("%s: test Query()", myName)
	rows, err := db.Query("SELECT 1")
	if err != nil {
		t.Fatal(err)
	}
	for rows.Next() {
		t.Logf("%s: read a row", myName)
		if err := rows.Scan(&one); err != nil {
			t.Fatal(err)
		}
		debug.Msg("processed row: one: %d", one)
		t.Logf("processed row: one: %d", one)
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
	if err = defaultClose(myName, db); err != nil {
		t.Errorf(err.Error())
	}
}

// Test05CharsetConnection
//
// Should do something here to store the data in one character
// set and then connect and get it back in another one.
// That's a bit harder as we don't use/process the result set charset
// data as this may need to be converted within Go to a string format
// which would be UTF8 (unicode) I guess.
