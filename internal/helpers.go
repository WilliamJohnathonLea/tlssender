package internal

import (
	"crypto/tls"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"strings"
)

func connectTCP(hostAndPort string, useSecureConnection bool) (io.ReadWriteCloser, error) {
	if useSecureConnection {
		tlsConf := &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
		return tls.Dial("tcp", hostAndPort, tlsConf)
	}

	return net.Dial("tcp", hostAndPort)
}

func bytesToTCPMessage(msg []byte, useBigEndian bool) []byte {
	size := len(msg)
	sizeBuff := make([]byte, 4) // 32-bit unsigned int

	if useBigEndian {
		binary.BigEndian.PutUint32(sizeBuff, uint32(size))
	} else {
		binary.LittleEndian.PutUint32(sizeBuff, uint32(size))
	}

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
