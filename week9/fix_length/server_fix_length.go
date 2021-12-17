package fix_length

import (
	"fmt"
	"net"
)

func ServerTcpFixLength(server net.Conn) {
	fmt.Println("server fixed length")
	const BYTES = 1024
	for {
		var buf = make([]byte, BYTES)
		_, err := server.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("client data: ", string(buf))
	}
}




