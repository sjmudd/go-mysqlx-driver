package test

import (
	"fmt"
	"testing"

	"github.com/sjmudd/go-mysqlx-driver/debug"
)

// genericRowsTest runs the query and checks we get back the expected number of rows
func genericRowsTest(name, query string, expectedRows int) error {
	const myName = "genericRowsTest"

	db, err := defaultOpen(myName)
	if err != nil {
		return fmt.Errorf("%s: failed: %+v", myName, err)
	}

	debug.Msg("%s: %s: %q", myName, name, query)
	var (
		collectedRows int
		i             int    // integer response from the query
		s             string // string response from the query
	)

	rows, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("%s failed: %+v", myName, err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&i, &s); err != nil {
			return fmt.Errorf("%s: scan failed: %+v", myName, err)
		}
		collectedRows++
		debug.Msg("genericRowsTest: %s: collectedRows: %d, values: %d, %q", name, collectedRows, i, s)
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("%s rows.Err() != nil: %+v", myName, err)
	}
	if collectedRows != expectedRows {
		return fmt.Errorf("%s: Received %d rows, exected: %d", myName, collectedRows, expectedRows)
	}

	if err = defaultClose(myName, db); err != nil {
		return fmt.Errorf("%s failed: %v", myName, err)
	}

	return nil // all ok
}

// Ensure we get no rows back
func TestZeroRows(t *testing.T) {
	const myName = "TestZeroRows"
	const myQuery = `SELECT '' AS result WHERE result IS NULL`

	if err := genericRowsTest(myName, myQuery, 0); err != nil {
		t.Fatalf("%s failed: %+v", myName, err)
	}
}

// Ensure we get one row back
func TestOneRow(t *testing.T) {
	const myName = "TestOneRow"
	const myQuery = `SELECT 1, 'something'`

	if err := genericRowsTest(myName, myQuery, 1); err != nil {
		t.Fatalf("%s failed: %+v", myName, err)
	}
}

// Ensure we get two rows back
func TestTwoRows(t *testing.T) {
	const myName = "TestTwoRows"
	const myQuery = `SELECT * FROM ( (SELECT 1, "ONE") UNION ALL ( SELECT 2, "TWO") ) a`

	if err := genericRowsTest(myName, myQuery, 2); err != nil {
		t.Fatalf("%s failed: %+v", myName, err)
	}
}
// Ensure we get two rows back
func TestThreeRows(t *testing.T) {
	const myName = "TestThreeRows"
	const myQuery = `SELECT * FROM ( (SELECT 1, "ONE") UNION ALL (SELECT 2, "TWO") UNION ALL (SELECT 3, "THREE") ) a`

	if err := genericRowsTest(myName, myQuery, 3); err != nil {
		t.Fatalf("%s failed: %+v", myName, err)
	}
}
