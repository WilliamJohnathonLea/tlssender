package internal

import (
	"log"
)

func HandleFile(hostAndPort, filepath string, ack, useSecureConnection, useBigEndian bool) {
	conn, err := connectTCP(hostAndPort, useSecureConnection)
	if err != nil {
		log.Fatalf("failed to connect: %v", err.Error())
	}
	defer conn.Close()
	err = sendFile(filepath, conn, ack, useBigEndian)
	if err != nil {
		log.Fatalf("failed while sending file: %v", err)
	}
}

func HandleDir(hostAndPort, dirpath string, ack, useSecureConnection, useBigEndian bool) {
	conn, err := connectTCP(hostAndPort, useSecureConnection)
	if err != nil {
		log.Fatalf("failed to connect: %v", err.Error())
	}
	defer conn.Close()
	err = sendDir(dirpath, conn, ack, useBigEndian)
	if err != nil {
		log.Fatalf("failed while sending files: %v", err)
	}
}
