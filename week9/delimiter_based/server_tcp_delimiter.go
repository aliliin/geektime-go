package delimiter_based

import (
	"bufio"
	"fmt"
	"net"
)

func ServerTcpDelimiter(conn net.Conn) {
	fmt.Println("server, delimiter based")
	reader := bufio.NewReader(conn)
	for {
		slice, err := reader.ReadSlice('\n')
		Log(slice)
		Log(err)
		if err != nil {
			Log("delimiter based: ", err)
			continue
		}
		fmt.Printf("slice %s", slice)
	}
}
