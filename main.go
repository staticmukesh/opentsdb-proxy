package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/staticmukesh/opentsdb-proxy/conf"
	proxy "github.com/staticmukesh/opentsdb-proxy/proxy"
)

var flagConf = flag.String("conf", "opentsdb.conf", "Location of configuration file. Defaults to opentsdb.toml in directory of the proxy executable.")
var cmds chan *string

func main() {
	flag.Parse()
	conf := conf.ReadConf(flagConf)

	initialize(conf)
	proxy.Init(conf, cmds)

	startServer(conf.Host)
}

func initialize(conf *conf.Conf) {
	cmds = make(chan *string)
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

		go handleConnection(conn, cmds)
	}
}

func handleConnection(conn net.Conn, cmds chan *string) {
	fmt.Println("Handling request...")
	for {
		cmd, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error while reading data: ", err.Error())
			conn.Close()
			break
		}
		cmds <- &cmd
		fmt.Print("Command received: ", string(cmd))
	}
}
