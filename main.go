package main

import (
	"flag"
	"log"
	"os"

	"github.com/WilliamJohnathonLea/tlssender/internal"
)

const (
	fileCmdStr = "file"
	dirCmdStr = "dir"
)

func main() {
	fileCmd := flag.NewFlagSet(fileCmdStr, flag.ExitOnError)
	fileAck := fileCmd.Bool("ack", false, "Get acknowlegment of sent packet")
	dirCmd := flag.NewFlagSet(dirCmdStr, flag.ExitOnError)
	dirAck := dirCmd.Bool("ack", false, "Get acknowlegment for each sent packet")

	switch os.Args[1] {
	case fileCmdStr:
		fileCmd.Parse(os.Args[2:])
		hostAndPort := fileCmd.Arg(0)
		filepath := fileCmd.Arg(1)
		conn, err := internal.ConnectTCP(hostAndPort)
		if err != nil {
			log.Fatalf("failed to connect: %v", err.Error())
		}
		defer conn.Close()
		err = internal.SendFile(filepath, conn, *fileAck)
		if err != nil {
			log.Fatalf("failed while sending file: %v", err)
		}
	case dirCmdStr:
		dirCmd.Parse(os.Args[2:])
		hostAndPort := fileCmd.Arg(0)
		dir := fileCmd.Arg(1)
		conn, err := internal.ConnectTCP(hostAndPort)
		if err != nil {
			log.Fatalf("failed to connect: %v", err.Error())
		}
		defer conn.Close()
		err = internal.SendDir(dir, conn, *dirAck)
		if err != nil {
			log.Fatalf("failed while sending files: %v", err)
		}
	default:
		log.Fatalf("expected subcommand '%s' or '%s'", fileCmdStr, dirCmdStr)
	}
}
