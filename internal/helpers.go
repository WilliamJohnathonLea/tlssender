package internal

import (
	"crypto/tls"
	"encoding/binary"
	"errors"
	"strings"
)

func ConnectTCP(hostAndPort string) (*tls.Conn, error) {
	tlsConf := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}
	return tls.Dial("tcp", hostAndPort, tlsConf)
}

func bytesToTCPMessage(msg []byte) []byte {
	size := len(msg)
	sizeBuff := make([]byte, 4) // 32-bit unsigned int

	binary.LittleEndian.PutUint32(sizeBuff, uint32(size))

	out := append(sizeBuff, msg...)
	return out
}

func parseResponse(bytes []byte) string {
	dataLen := binary.LittleEndian.Uint32(bytes[:4])
	bytesToRead := 4 + dataLen
	// We need to exlcude the 1st four bytes
	// because they simply indicate the length
	// of the message
	return string(bytes[4:bytesToRead])
}

func handleAck(bytes []byte) error {
	res := parseResponse(bytes)
	if strings.Contains(res, "OK") {
		return nil
	}
	return errors.New("unexpected response " + res)
}
