package test

import (
	"testing"

	driver "github.com/sjmudd/go-mysqlx-driver"
)

func TestMySQL41(t *testing.T) {
	dbname := "world_x"
	username := "root"
	dbPassword := "fff"
	nineteenBytes := make([]byte, 19)
	twentyBytes := make([]byte, 20)
	twentytwoBytes := make([]byte, 22)

	// check we have a username
	if auth := driver.NewMySQL41(dbname, "", dbPassword); auth != nil {
		t.Errorf("MySQL41() with empty user should return nil, returned %+v", auth)
	}

	// throw an error if we have < 20 bytes
	auth := driver.NewMySQL41(dbname, username, dbPassword)
	if _, err := auth.GetNextAuthData(nineteenBytes); err == nil {
		t.Errorf("MySQL41().GetNextAuthData(<19 bytes>) did not return an error")
	}

	// throw an error if we have > 20 bytes
	if _, err := auth.GetNextAuthData(twentytwoBytes); err == nil {
		t.Errorf("MySQL41().GetNextAuthData(<22 bytes>) did not return an error")
	}

	// check length of result, should be length(usernname) + 2 x \x00 + 1 x '*' + 20-byte hash
	result, _ := auth.GetNextAuthData(twentyBytes)
	expected := len(dbname) + 2 + len(username) + 1 + 40
	if len(result) != expected {
		t.Errorf("MySQL41().GetNextAuthData(<20 bytes>) length: %d, expected: %d",
			len(result),
			expected)
	}

	// The first characters are the dbname so check \x00 after string of dbname
	if len(dbname) < len(result) && result[len(dbname)] != 0 && len(dbname) == 0 {
		t.Errorf("MySQL41().GetNextAuthData(<20 bytes>) result[0]: %d, expected: %d",
			result[0], 0)
	}

	// check that username is after the <dbname>\x00
	if string(result[1+len(dbname):len(dbname)+len(username)+1]) != username {
		t.Errorf("TestMySQL41() username not found where expected, got %s, expected: %s", string(result[1+len(dbname):len(dbname)+len(username)+1]), dbPassword)
	}

	// check for \x00 after username
	if result[1+len(dbname)+len(username)] != 0 {
		t.Errorf("TestMySQL41() no null byte after username in result")
	}

	// check expectd hash of password (1)
	{
		buf := []byte{0xa, 0x35, 0x42, 0x1a, 0x43, 0x47, 0x6d, 0x65, 0x1, 0x4a, 0xf, 0x4c, 0x9, 0x5c, 0x32, 0x61, 0x64, 0x3c, 0x13, 0x6}
		result, _ := auth.GetNextAuthData(buf)

		exp := "*34439ed3004cf0e6030a9ec458338151bfb4e22d"
		if string(result[len(dbname)+len(dbPassword)+3:]) != exp {
			t.Errorf("hash of password %q: got %s, expected: %s", dbPassword, result, exp)
		}
	}
	// check expectd hash of password (2)
	{
		buf := []byte{0x41, 0x43, 0x56, 0x6e, 0x78, 0x19, 0x2c, 0x2c, 0x19, 0x6f, 0x18, 0x29, 0x05, 0x52, 0x3c, 0x62, 0x39, 0x3d, 0x5c, 0x77}
		result, _ := auth.GetNextAuthData(buf)

		exp := "*af1ef523d254181abb1155c1fbc933b80c2ec853"
		if string(result[len(dbname)+len(dbPassword)+3:]) != exp {
			t.Errorf("hash of password %q: got %s, expected: %s", dbPassword, result, exp)
		}
	}

	// should properly encode a result containing 0x10 and write "10" into the response buffer
	{
		buf := []byte{0x7a, 0x59, 0x6b, 0x6e, 0x19, 0x7f, 0x44, 0x1, 0x6f, 0x4a, 0xf, 0xf, 0x3e, 0x19, 0x50, 0x4c, 0x4f, 0x47, 0x53, 0x5b}
		result, _ := auth.GetNextAuthData(buf)

		exp := "*950d944626109ab5bce8bc56a4e78a296e34271d"
		if string(result[len(dbname)+len(dbPassword)+3:]) != exp {
			t.Errorf("hash of password %q: got %s, expected: %s", dbPassword, result, exp)
		}
	}

	// FIXME: I don't distinguish between no password and empty password. Need to fix.
	{
		auth := driver.NewMySQL41(dbname, username, "")
		result, _ := auth.GetNextAuthData(twentyBytes)
		if len(result) != len(dbname)+len(username)+2 {
			t.Errorf("GetNextAuthData() called with dbname %q, username %q and no password should have length %d, got %d", dbname, username, len(dbname)+len(username)+2, len(result))
		}
	}
}
