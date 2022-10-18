package main

import (
	"flag"
	"log"
	"os"

	"github.com/WilliamJohnathonLea/tlssender/internal"
)

const (
	fileCmdStr = "file"
	dirCmdStr  = "dir"
)

func main() {
	var useSecureConnection bool = false
	var useBigEndian bool = false
	var ack bool = false

	fileCmd := flag.NewFlagSet(fileCmdStr, flag.ExitOnError)
	fileCmd.BoolVar(&ack, "ack", false, "Get acknowlegment of sent packet")
	fileCmd.BoolVar(&useSecureConnection, "secure", false, "Use TLS secure connection")
	fileCmd.BoolVar(&useBigEndian, "big", false, "Use Big Endian byte order")
	dirCmd := flag.NewFlagSet(dirCmdStr, flag.ExitOnError)
	dirCmd.BoolVar(&ack, "ack", false, "Get acknowlegment for each sent packet")
	dirCmd.BoolVar(&useSecureConnection, "secure", false, "Use TLS secure connection")
	dirCmd.BoolVar(&useBigEndian, "big", false, "Use Big Endian byte order")

	switch os.Args[1] {
	case fileCmdStr:
		fileCmd.Parse(os.Args[2:])
		hostAndPort := fileCmd.Arg(0)
		filepath := fileCmd.Arg(1)

		internal.HandleFile(hostAndPort, filepath, ack, useSecureConnection, useBigEndian)
	case dirCmdStr:
		dirCmd.Parse(os.Args[2:])
		hostAndPort := dirCmd.Arg(0)
		dir := dirCmd.Arg(1)

		internal.HandleDir(hostAndPort, dir, ack, useSecureConnection, useBigEndian)
	default:
		log.Fatalf("expected subcommand '%s' or '%s'", fileCmdStr, dirCmdStr)
	}
}
