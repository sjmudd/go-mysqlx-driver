package test

import (
	"fmt"
	"testing"

	"github.com/sjmudd/go-mysqlx-driver/debug"
)

// test doing a query and returning the result
func testAnyNonNullType(caller string, column string, table string, myVariable interface{}) error {
	db, err := defaultOpen(caller)
	if err != nil {
		return err
	}

	// test all types by collecting data
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s IS NOT NULL", column, table, column)
	debug.Msg("%s: query: %q", caller, query)
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	for rows.Next() {
		if err := rows.Scan(myVariable); err != nil {
			return err
		}
		debug.Msg("processed row: %s: %v", column, myVariable)
	}
	if err := rows.Err(); err != nil {
		return err
	}
	if err := rows.Close(); err != nil {
		return fmt.Errorf("%s: Unable to close rows:", caller, err)
	}

	err = defaultClose(caller, db)
	if err != nil {
		return err
	}

	return nil
}

// TestTinyIntU tests MySQL tinyint unsigned, go uint8
func TestTinyIntU(t *testing.T) {
	var myTinyintU uint8

	const (
		caller = "TestTinyIntU"
		column = "my_tinyint_u"
		table  = "xprotocol_test_tinyint_u"
	)

	if err := testAnyNonNullType(caller, column, table, &myTinyintU); err != nil {
		t.Error("%s: ", caller, err)
	}
}

// TestTinyInt tests MySQL tinyint, go int8
func TestTinyInt(t *testing.T) {
	var myTinyint int8
	const (
		caller = "TestTinyInt"
		column = "my_tinyint"
		table  = "xprotocol_test_tinyint"
	)

	if err := testAnyNonNullType(caller, column, table, &myTinyint); err != nil {
		t.Error("%s: ", caller, err)
	}
}

func TestSmallintU(t *testing.T) {
	var mySmallintU uint16
	const (
		caller = "TestSmallintU"
		column = "my_smallint_u"
		table  = "xprotocol_test_smallint_u"
	)

	if err := testAnyNonNullType(caller, column, table, &mySmallintU); err != nil {
		t.Error("%s: ", caller, err)
	}
}

// TestSmallint tests MySQL smallint, go int16
func TestSmallint(t *testing.T) {
	var mySmallint int16
	const (
		caller = "TestSmallint"
		column = "my_smallint"
		table  = "xprotocol_test_smallint"
	)

	if err := testAnyNonNullType(caller, column, table, &mySmallint); err != nil {
		t.Error("%s: ", caller, err)
	}
}

// TestMediumint tests MySQL mediumint unsigned, has no direct go type, use go uint32
func TestMediumintU(t *testing.T) {
	var myMediumintU uint32
	const (
		caller = "TestMediumintU"
		column = "my_mediumint_u"
		table  = "xprotocol_test_mediumint_u"
	)

	if err := testAnyNonNullType(caller, column, table, &myMediumintU); err != nil {
		t.Error("%s: ", caller, err)
	}
}

// TestMediumint tests MySQL mediumint, has no direct go type, use go int32
func TestMediumint(t *testing.T) {
	var myMediumint int32
	const (
		caller = "TestMediumint"
		column = "my_mediumint"
		table  = "xprotocol_test_mediumint"
	)

	if err := testAnyNonNullType(caller, column, table, &myMediumint); err != nil {
		t.Error("%s: ", caller, err)
	}
}

// TestIntU type (unsigned 32bits)
func TestIntU(t *testing.T) {
	var myIntU uint32
	const (
		caller = "TestInt"
		column = "my_int_u"
		table  = "xprotocol_test_int_u"
	)

	if err := testAnyNonNullType(caller, column, table, &myIntU); err != nil {
		t.Error("%s: ", caller, err)
	}
}

// TestInt tests MySQL int, go int32
func TestInt(t *testing.T) {
	var myInt int32
	const (
		caller = "TestInt"
		column = "my_int"
		table  = "xprotocol_test_int"
	)

	if err := testAnyNonNullType(caller, column, table, &myInt); err != nil {
		t.Error("%s: ", caller, err)
	}
}

// TestFloat tests MySQL float, go float32
func TestFloat(t *testing.T) {
	var myFloat float32
	const (
		caller = "TestFloat"
		column = "my_float"
		table  = "xprotocol_test_float"
	)

	if err := testAnyNonNullType(caller, column, table, &myFloat); err != nil {
		t.Error("%s: ", caller, err)
	}
}

// TestDouble tests MySQL double, go float64
func TestDouble(t *testing.T) {
	var myDouble float64
	const (
		caller = "TestDouble"
		column = "my_double"
		table  = "xprotocol_test_double"
	)

	if err := testAnyNonNullType(caller, column, table, &myDouble); err != nil {
		t.Error("%s: ", caller, err)
	}
}
