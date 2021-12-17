package protocol

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

//使用了length field based frame decoder
const (
	PackageLengthBytes   = 4
	HeaderLengthBytes = 2
	VersionBytes      = 2
	OperationBytes    = 4
	SequenceIDBytes      = 4

	HeaderLength = PackageLengthBytes + HeaderLengthBytes + VersionBytes + OperationBytes + SequenceIDBytes

	// Body
)

// Depack 解码器
func Depack(buffer []byte) []byte {
	length := len(buffer)

	var i int
	data := make([]byte, 32)

	for i = 0; i < length; i++ {
		if length < i+HeaderLength {
			break
		}

		messageLength := ByteToInt(buffer[i : i+PackageLengthBytes])
		if length < i+HeaderLength+messageLength {
			break
		}

		site := i + PackageLengthBytes
		headerLength := ByteToInt(buffer[site : site+HeaderLengthBytes])
		site += HeaderLengthBytes

		protocolVersion := ByteToInt16(buffer[site : site+VersionBytes])
		site += VersionBytes

		operation := ByteToInt(buffer[site : site+OperationBytes])
		site += OperationBytes

		sequenceID := ByteToInt(buffer[site : site+SequenceIDBytes])
		site += SequenceIDBytes

		fmt.Printf("packageLength: %d, headerLength: %d , protocolVersion: %d, operation: %d, sequenceID: %d \n", messageLength, headerLength, protocolVersion, operation, sequenceID)

		data = buffer[i+HeaderLength : i+HeaderLength+messageLength]
		break
	}

	if i == length {
		return make([]byte, 0)
	}

	return data
}

// ByteToInt 字节转换成整形
func ByteToInt(n []byte) int {
	bytesbuffer := bytes.NewBuffer(n)
	var x int32
	binary.Read(bytesbuffer, binary.BigEndian, &x)

	return int(x)
}

func ByteToInt16(n []byte) int {
	bytebuffer := bytes.NewBuffer(n)
	var x int16
	binary.Read(bytebuffer, binary.BigEndian, &x)

	return int(x)
}

// Packet 封包
func Packet(message []byte) []byte {
	b := append(Int32ToBytes(len(message)), Int16ToBytes(0)...)
	b = append(b, Int16ToBytes(8)...)
	b = append(b, Int32ToBytes(99)...)
	b = append(b, Int32ToBytes(10)...)
	b = append(b, message...)

	return b
}

// Int32ToBytes 整数转换成字节
func Int32ToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

// Int16ToBytes 整数转换成字节
func Int16ToBytes(n int) []byte {
	x := int16(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}
