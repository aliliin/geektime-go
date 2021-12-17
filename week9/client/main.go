package main

import (
	"fmt"
	"net"
	"os"
	"week9/delimiter_based"
)

func main() {
	server := "localhost:7373"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("connect success")

	//fix_length.ClientTcpFixLength(conn)
	// delimiter based
	delimiter_based.ClientTcpDelimiter(conn)
	// length field based frame decoder
	//frame_decoder.ClientTcpFrameDecoder(conn)
	defer conn.Close()
}
