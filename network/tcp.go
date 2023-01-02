package network

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"net"
)

// Buffer size must be > 1
func TCPReader(bufferSize int, connection *net.Conn) []byte {
	var buffer []byte
	reader := bufio.NewReader(*connection)

	for {
		tempBuffer := make([]byte, bufferSize)
		_, err := reader.Read(tempBuffer)
		if err != nil {
			log.Fatal(err)
			if err == io.EOF {
				break
			}
		}

		// Append to global buffer
		buffer = append(buffer, tempBuffer...)

		// 0 Value indicates end of read for byte stream
		if tempBuffer[bufferSize-1] == 0 {
			break
		}

	}
	buffer = bytes.Trim(buffer, "\x00") // Remove extra zeroes from buffer
	return buffer
}
