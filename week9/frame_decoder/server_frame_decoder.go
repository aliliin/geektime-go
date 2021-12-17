package frame_decoder

import (
	"fmt"
	"log"
	"net"
	"week9/protocol"
)

// ServerTcpFrameDecoder length field based frame decoder
func ServerTcpFrameDecoder(conn net.Conn) {
	fmt.Println("server, length field based frame decoder")
	// 缓冲区，存储被截断的数据
	tmpBuffer := make([]byte, 0)
	//接收解包
	readerChannel := make(chan []byte, 10000)
	go reader(readerChannel)

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			Log(conn.RemoteAddr().String(), "connection error: ", err)
			return
		}

		tmpBuffer = protocol.Depack(append(tmpBuffer, buffer[:n]...))
		readerChannel <- tmpBuffer //接收的信息写入通道
	}

}

// 获取通道数据
func reader(readerChannel chan []byte) {
	for {
		select {
		case data := <-readerChannel:
			Log("channel:", string(data)) // 打印通道内的信息
		}
	}
}

// Log 日志处理
func Log(v ...interface{}) {
	log.Println(v...)
}
