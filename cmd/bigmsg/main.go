package main

import (
	"log"
	"net"
)

const (
	destination = "127.0.0.1:33060"
	waitSize    = 4                  // 4 message length bytes we expect to get back
)

var (
	fullMsg  = []byte{255, 255, 255, 255, 255} // to send: 2^32-1 bytes of message type 255 (does not exist)
)

func main() {
	conn, err := net.Dial("tcp", destination)
	if err != nil {
		log.Fatalf("net.Dial failed: %v", err)
	}
	log.Printf("connected to: %v", destination)

	var bytes int
	bytes, err = conn.Write(fullMsg)
	if bytes != 5 || err != nil {
		log.Fatalf("net.Write failed: only wrote %d of %d bytes: %v", bytes, len(fullMsg), err)
	}
	log.Printf("Wrote %d bytes indicating full length message size", bytes)

	var buf = make([]byte, waitSize) // bytes to wait for

	log.Printf("Waiting for %d bytes of response", waitSize)
	bytes, err = conn.Read(buf)

	if err != nil {
		log.Fatalf("net.Read failed: only read %d of %d bytes: %v", bytes, len(buf), err)
	}
	log.Printf("Read %d bytes of buf Message", bytes)

	if err = conn.Close(); err != nil {
		log.Fatalf("conn.Close failed: %v", err)
	}
	log.Printf("closed connected to: %v", destination)

}
