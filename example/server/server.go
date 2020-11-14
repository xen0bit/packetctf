// server.go

package main

import (
	"bytes"
	"log"
	"net"
	"time"

	"encoding/gob"
)

// Create your custom data struct
type Color struct {
	Red   int
	Green int
	Blue  int
}

func main() {
	// for purpose of verbosity, I will be removing error handling from this
	// sample code

	server, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := server.Accept()

	// loop through the connection to read incoming connections. If you're doing by
	// directional, you might want to make this into a seperate go routine
	for {
		msg := Color{Red: 255, Green: 0, Blue: 0}
		bin_buf := new(bytes.Buffer)

		// create a encoder object
		gobobj := gob.NewEncoder(bin_buf)

		// encode buffer and marshal it into a gob object
		gobobj.Encode(msg)

		conn.Write(bin_buf.Bytes())

		time.Sleep(1 * time.Second)

	}

}
