package fix_length

import (
	"fmt"
	"net"
)

//ClientTcpFixLength 客户端 fix length
func ClientTcpFixLength(server net.Conn) {
	fmt.Println("client fixed length")
	sendByte := make([]byte, 1024)
	sendMessage := "{\"gyl01\":1,\"gyl02\",2}"
	for i := 0; i < 1000; i++ {
		temp := []byte(sendMessage)
		for j := 0; j < len(temp) && j < 1024; j++ {
			sendByte[j] = temp[j]
		}
		_, err := server.Write(sendByte)
		if err != nil {
			panic(err)
		}
		fmt.Println("send msg over")
	}
}
