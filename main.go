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
var maxConnChan chan int
var commandChan chan *string

func main() {
	flag.Parse()
	conf := conf.ReadConf(flagConf)

	initialize(conf)
	proxy.Init(conf, commandChan)

	startServer(conf.Host)
}

func initialize(conf *conf.Conf) {
	maxConnChan = make(chan int, conf.LimitConnection)

	for i := 0; i < conf.LimitConnection; i++ {
		maxConnChan <- 0
	}

	cmdChan := make(chan *string)
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
		<-maxConnChan
		fmt.Println("Accepting connection...")
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error: ", err.Error())
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println("Handling request...")
	for {
		cmd, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error while reading data: ", err.Error())
			conn.Close()
			maxConnChan <- 0
			break
		}

		cmdChan <- &cmd
		fmt.Print("Command received: ", string(cmd))
	}
}
