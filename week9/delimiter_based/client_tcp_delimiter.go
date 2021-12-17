package delimiter_based

import (
	"fmt"
	"log"
	"net"
)

func ClientTcpDelimiter(conn net.Conn) {
	fmt.Println("client, delimiter based")
	var sends string
	sendMsg := "{\"gyl01\":1,\"gyl02\",2}\n"
	for i := 0; i < 10; i++ {
		sends += sendMsg
		s, err := conn.Write([]byte(sends))
		Log(s)
		if err != nil {
			fmt.Println(err, ",err index=", i)
			return
		}
		Log("send over once")
	}
}

// Log 日志处理
func Log(v ...interface{}) {
	log.Println(v...)
}
