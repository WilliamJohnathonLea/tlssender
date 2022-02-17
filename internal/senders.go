package internal

import (
	"crypto/tls"
	"io/ioutil"
	"os"
)

// The maximum size of a TCP frame in bytes
const maxFrameSize = 256

// SendFile sends the contents of a file to a TCP connection.
// filepath is the path to the data to be sent.
// conn is the connection where the data will be sent.
// if ack is == true then we will make sure the data was received with an 'OK'
func SendFile(filepath string, conn *tls.Conn, ack bool) error {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	_, err = conn.Write(bytesToTCPMessage(bytes))
	if ack {
		respBytes := make([]byte, maxFrameSize)
		nRead, err := conn.Read(respBytes)
		if err != nil {
			return err
		}
		return handleAck(respBytes[:nRead])
	}
	return err
}

// SendDir sends the contents of files in a directory to a TCP connection.
// dirpath is the path to the data directory.
// conn is the connection where the data will be sent.
// If ack is == true then we will make sure the data was received with an 'OK'
func SendDir(dirpath string, conn *tls.Conn, ack bool) error {
	files, err := ioutil.ReadDir(dirpath)
	if err != nil {
		return err
	}
	for _, f := range files {
		filepath := dirpath + "/" + f.Name()
		err = SendFile(filepath, conn, ack)
		if err != nil {
			return err
		}
	}
	return nil
}
