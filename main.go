package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/staticmukesh/opentsdb-proxy/conf"
)

var flagConf = flag.String("conf", "opentsdb.conf", "Location of configuration file. Defaults to opentsdb.toml in directory of the proxy executable.")

func main() {
	flag.Parse()

	conf := conf.ReadConf(flagConf)

	startServer(conf.Host)
}

func startServer(host string) {
	ln, err := net.Listen("tcp", host)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}

	defer ln.Close()

	fmt.Println("Server listening on " + host)

	for {
		fmt.Println("Accepting connection...")
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error: ", err.Error())
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	fmt.Println("Handling request...")
	for {
		cmd, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error while reading data: ", err.Error())
			conn.Close()
			break
		}

		fmt.Print("Command received: ", string(cmd))
	}
}
