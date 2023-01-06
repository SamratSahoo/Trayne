package network

import (
	"bufio"
	"bytes"
	"io"
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
			if err == io.EOF {
				break
			}
			return []byte{}
		}

		tempBuffer = bytes.Trim(tempBuffer, "\x00") // Remove extra zeroes from buffer
		// Append to global buffer
		buffer = append(buffer, tempBuffer...)
	}
	return buffer
}
