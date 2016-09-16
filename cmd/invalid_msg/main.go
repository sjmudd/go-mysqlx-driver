package main

import (
	"log"
	"net"
)

const (
	bufSize     = 1024
	destination = "127.0.0.1:33060"
	msgType     = 255 // invalid message, currently unknown
	waitSize    = 4   // 4 message length bytes we expect to get back
	msgText     = "hello"
)

func prepare(msgType uint8, msgText []byte) []byte {

	b := make([]byte, len(msgText)+5)
	len := uint8(len(msgText) + 1) // horrible hack

	b[0] = len // little endian hack (ignoring top 3 bytes)
	b[4] = msgType
	copy(b[5:], msgText)

	return b
}

func main() {
	conn, err := net.Dial("tcp", destination)
	if err != nil {
		log.Fatalf("net.Dial failed: %v", err)
	}
	log.Printf("connected to: %v", destination)

	// create msg
	b := prepare(255, []byte(msgText))

	var bytes int
	bytes, err = conn.Write(b)
	if bytes != len(b) || err != nil {
		log.Fatalf("net.Write failed: only wrote %d of %d bytes: %v", bytes, len(b), err)
	}
	log.Printf("Wrote %d bytes indicating full length message: %q", bytes, msgText)

	var buf = make([]byte, waitSize, bufSize) // bytes to wait for

	log.Printf("Waiting for %d bytes of response", waitSize)
	bytes, err = conn.Read(buf)

	if err != nil {
		log.Fatalf("net.Read failed: only read %d of %d bytes: %v", bytes, len(buf), err)
	}
	log.Printf("Read %d bytes of message length: %+v", bytes, buf)

	// horrible hack, only look at length first byte
	msgLen := buf[0]
	buf = make([]byte, msgLen)
	log.Printf("Reading %d bytes as message content", msgLen)
	bytes, err = conn.Read(buf)
	if err != nil {
		log.Fatalf("net.Read failed: only read %d of %d bytes (trying to pull in msg): %v", bytes, msgLen, err)
	}
	log.Printf("msg type: %d", buf[0])
	log.Printf("msg text (raw): %+v", string(buf[1:]))

	// try to read some more data (connection expected to be dropped so expect to get EOF)
	log.Printf("Trying to read more data if its there")
	buf = make([]byte, 1)
	bytes, err = conn.Read(buf)
	if err != nil {
		if err.Error() == "EOF" {
			log.Printf("We're at EOF (as expected)")
		} else {
			log.Fatalf("net.Read failed: only read %d of %d bytes (trying to pull in extra data): %v", bytes, len(buf), err)
		}
	}

	if err = conn.Close(); err != nil {
		log.Fatalf("conn.Close failed: %v", err)
	}
	log.Printf("closed connected to: %v", destination)

}
