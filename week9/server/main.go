package main

import (
	"fmt"
	"log"
	"net"
	"week9/delimiter_based"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:7373")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer lis.Close()
	Log("Waiting for client ...") // 启动后，等待客户端访问。

	for {
		conn, err := lis.Accept() // 监听客户端
		if err != nil {
			fmt.Println("conn error:", err)
			return
		}
		Log(conn.RemoteAddr().String(), "tcp connection success")
		// fix length
		//go fix_length.ServerTcpFixLength(conn)
		// delimiter based
		go delimiter_based.ServerTcpDelimiter(conn)
		// length field based frame decoder
		//go frame_decoder.ServerTcpFrameDecoder(conn)
	}

}

//日志处理
func Log(v ...interface{}) {
	log.Println(v...)
}