// client.go
package main

import (
	"bytes"
	"fmt"
	"log"
	"net"

	"encoding/gob"
)

type Color struct {
	Red   int
	Green int
	Blue  int
}

func main() {

	// error handling still truncated
	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		log.Fatal(err)
	}

	// loop through the connection to read incoming connections. If you're doing by
	// directional, you might want to make this into a seperate go routine
	for {
		// create a temp buffer
		tmp := make([]byte, 500)

		_, err = conn.Read(tmp)
		if err != nil {
			log.Fatal(err)
		}

		// convert bytes into Buffer (which implements io.Reader/io.Writer)
		tmpbuff := bytes.NewBuffer(tmp)

		tmpstruct := new(Color)

		// creates a decoder object
		gobobj := gob.NewDecoder(tmpbuff)

		// decodes buffer and unmarshals it into a Message struct
		gobobj.Decode(tmpstruct)

		// lets print out!
		fmt.Println(tmpstruct) // reflects.TypeOf(tmpstruct) == Message{}

	}
}
