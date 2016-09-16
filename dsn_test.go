// Go driver for MySQL X Protocol
//
// Copyright 2016 Simon J Mudd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.
//
// MySQL X protocol authentication using MYSQL41 method

package mysql

import (
	"testing"
)

type testDSN struct {
	dsn    string
	result bool
}
type testDSNs []testDSN

// test that the use of xpprotocol=1 works
func TestDSN(t *testing.T) {

	d := testDSNs{
		{dsn: "user:pass@/", result: false},
		{dsn: "user:pass@tcp(127.0.0.1:33060)/test?xprotocol=1", result: true},
		{dsn: "user:pass@tcp(127.0.0.1)/test?xprotocol=1", result: true},
		{dsn: "user:pass@tcp(127.0.0.1:33060)/test", result: false},
		{dsn: "user:pass@tcp(127.0.0.1)/test", result: false},
		{dsn: "user:pass@/test", result: false},
		{dsn: "user:pass@/test?xprotocol=1", result: true},
	}

	for i := range d {
		cfg, err := parseDSN(d[i].dsn)
		if err != nil {
			t.Errorf("TestDSN: dsn: %s, parseDSN gives error: %v", d[i].dsn, err)
		}
		if cfg.useXProtocol != d[i].result {
			t.Errorf("TestDSN: dsn: %s, cfg.useXprotocol: %v, expected: %v", d[i].dsn, cfg.useXProtocol, d[i].result)
		}
	}
}
