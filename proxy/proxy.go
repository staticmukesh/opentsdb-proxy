package proxy

import (
	"fmt"
	"net"

	"github.com/staticmukesh/opentsdb-proxy/conf"
)

var cmdChans []chan string
var requestChan chan *string

func Init(conf *conf.Conf, cmdChan chan *string) {
	for index, server := range conf.Servers {
		cmdChans[index] = makeConnection(server)
	}

	requestChan = cmdChan

	go handleCommands(conf.Servers)
}

func makeConnection(server string) chan string {
	conn, err := net.Dial("tcp", server)

	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	cmdChan := make(chan string)
	go handleRequests(conn, cmdChan)

	return cmdChan
}

func handleRequests(conn net.Conn, cmdChan chan string) {
	for {
		cmd := <-cmdChan
		fmt.Fprintf(conn, cmd)
	}
}

func handleCommands(servers []string) {
	i := 0
	for {
		cmd := <-requestChan
		cmdChans[i] <- *cmd

		i++
		if i == len(servers) {
			i = 0
		}
	}
}
