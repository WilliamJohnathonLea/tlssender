package internal

import (
	"crypto/tls"
	"encoding/binary"
	"errors"
	"log"
	"net"
	"strings"
)

func ConnectTCP(hostAndPort string) (net.Conn, error) {
	return net.Dial("tcp", hostAndPort)
}

func HandleFile(hostAndPort, filepath string, ack bool) {
	conn, err := ConnectTCP(hostAndPort)
	if err != nil {
		log.Fatalf("failed to connect: %v", err.Error())
	}
	defer conn.Close()
	err = SendFile(filepath, conn, ack)
	if err != nil {
		log.Fatalf("failed while sending file: %v", err)
	}
}

func HandleDir(hostAndPort, dirpath string, ack bool) {
	conn, err := ConnectTCP(hostAndPort)
	if err != nil {
		log.Fatalf("failed to connect: %v", err.Error())
	}
	defer conn.Close()
	err = SendDir(dirpath, conn, ack)
	if err != nil {
		log.Fatalf("failed while sending files: %v", err)
	}
}

func ConnectSecureTCP(hostAndPort string) (*tls.Conn, error) {
	tlsConf := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}
	return tls.Dial("tcp", hostAndPort, tlsConf)
}

func SecureHandleFile(hostAndPort, filepath string, ack bool) {
	conn, err := ConnectSecureTCP(hostAndPort)
	if err != nil {
		log.Fatalf("failed to connect: %v", err.Error())
	}
	defer conn.Close()
	err = SecureSendFile(filepath, conn, ack)
	if err != nil {
		log.Fatalf("failed while sending file: %v", err)
	}
}

func SecureHandleDir(hostAndPort, dirpath string, ack bool) {
	conn, err := ConnectSecureTCP(hostAndPort)
	if err != nil {
		log.Fatalf("failed to connect: %v", err.Error())
	}
	defer conn.Close()
	err = SecureSendDir(dirpath, conn, ack)
	if err != nil {
		log.Fatalf("failed while sending files: %v", err)
	}
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
