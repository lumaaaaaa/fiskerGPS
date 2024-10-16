package main

import (
	"fmt"
	"net"
	"os"

	// UBX decoder
	"github.com/daedaleanai/ublox"

	// pretty print (for messages with unknow fields)
	"github.com/k0kubun/pp"
)

const (
	// your vehicle's IP
	FiskerIP = "192.168.4.25"
)

func main() {
	// connect to the Fisker Ocean over port 8888
	conn, err := net.Dial("tcp", FiskerIP+":8888")
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("[i] Connected to Fisker Ocean at 192.168.4.25:8888")

	decoder := ublox.NewDecoder(conn)
	for {
		// create a new UBX decoder
		decode, err := decoder.Decode()
		if err != nil {
			panic(err)
		}

		// pretty print the UBX message
		pp.Println(decode)
	}
}
