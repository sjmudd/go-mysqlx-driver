package main

import (
	"log"
	"net"
)

const (
	inputBufSize = 1024
	destination  = "127.0.0.1:33060"
	waitSize     = 4                  // 4 message length bytes we expect to get back
)

var (
	nullMsg  = []byte{0, 0, 0, 0} // to send
)

func main() {
	conn, err := net.Dial("tcp", destination)
	if err != nil {
		log.Fatalf("net.Dial failed: %v", err)
	}
	log.Printf("connected to: %v", destination)

	var bytes int
	bytes, err = conn.Write(nullMsg)
	if bytes != len(nullMsg) || err != nil {
		log.Fatalf("net.Write failed: only wrote %d of %d bytes: %v", bytes, len(nullMsg), err)
	}
	log.Printf("wrote %d bytes of null message length", bytes)

	var buf = make([]byte, waitSize, inputBufSize)

	log.Printf("waiting for %d bytes of response", waitSize)
	bytes, err = conn.Read(buf)

	if err != nil {
		log.Fatalf("net.Read failed: only read %d of %d bytes: %v", bytes, len(buf), err)
	}
	log.Printf("read %d bytes of buf Message", bytes)

	if err = conn.Close(); err != nil {
		log.Fatalf("conn.Close failed: %v", err)
	}
	log.Printf("closed connected to: %v", destination)

}
