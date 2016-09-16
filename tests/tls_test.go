package test

/*

// test we can connect to TLS using certs in this directory
func TestTLS(t *testing.T) {

	.....RegisterTLSConfig("testConfig", :x

        dsn := dsnPrefix + defaultDB + dsnSuffix
	const path = "./" // could be somewhere else but this will do for testing.

	rootCertPool := x509.NewCertPool()
	pem, err := ioutil.ReadFile(path + "ca-cert.pem")
	if err != nil {
		log.Fatal(err)
	}
	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		log.Fatal("Failed to append PEM.")
	}
	clientCert := make([]tls.Certificate, 0, 1)
	certs, err := tls.LoadX509KeyPair(path + "client-cert.pem", path + "client-key.pem")
	if err != nil {
		log.Fatal(err)
	}
	clientCert = append(clientCert, certs)
	mysql.RegisterTLSConfig("custom", &tls.Config{
		RootCAs: rootCertPool,
		Certificates: clientCert,
	})
	db, err := sql.Open("mysql", "user@tcp(localhost:3306)/test?tls=custom")
	if err != nil {
	}
	db.Close()
}


+// testing setting capabilities
+func (mc *mysqlXConn) capabilityTestTLS() error {
+       // - Trigger a no-op change
+       //   - check for 'tls' and set to the same value it has already (expected tobe false)
+       if mc.capabilities.Exists("tls") {
+               debug.Msg("Checking tls capability")
+               tls := mc.capabilities.Values("tls")
+               debug.Msg("- got back %d value(s) for tls", len(tls))
+               if len(tls) == 1 {
+                       tlsType := tls[0].Type()
+                       debug.Msg("- tls has one value of type: %s", tlsType)
+                       if tlsType == "bool" {
+                               tlsValue := tls[0].Bool()
+                               debug.Msg("- tls value: %v", tlsValue)
+
+                               // reset the value to the value it has at the moment...
+                               debug.Msg("Setting tls to the same value it has already")
+                               if err := mc.setScalarBoolCapability("tls", tlsValue); err != nil {
+                                       //                              debug.Msg("Enabling tls though the code won't handle it yet...")
+                                       //                              if err := mc.setScalarBoolCapability("tls", true); err != nil {
+                                       debug.Msg("capabilityTestTLS fails: %v", err)
+                                       return err
+                               }
+                       }
+               }
+       }
+       return nil
+}

Test05TLSTest(t *testing.T) {
        var (
                hostname           string
                max_allowed_packet int
                version            string
        )

        dsn := dsnPrefix + defaultDB + dsnSuffix
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
        defer rows.Close()
        for rows.Next() {
                if err := rows.Scan(&max_allowed_packet,&version,&hostname); err != nil {
                        t.Fatal(err)
                }
                t.Logf("processed row: max_allowed_packet: %d, version: %q, hostname: %q", max_allowed_packet,version,hostname)
                debug.Msg("processed row: max_allowed_packet: %d, version: %q, hostname: %q", max_allowed_packet,version,hostname)
        }
        if err := rows.Err(); err != nil {
                t.Fatal(err)
        }

        err = db.Close()
        if err != nil {
                t.Errorf("Test04SystemVariables: could not connect to dsn:%s: %v", dsn, err)
        }
}
*/
