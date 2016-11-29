package proxy

import (
	"fmt"
	"net"

	"github.com/staticmukesh/opentsdb-proxy/conf"
)

var cmdChans [1]chan *string //TODO Hardcoded 1
var requestChan chan *string

func Init(conf *conf.Conf, cmdChan chan *string) {
	fmt.Println("Proxy Init")

	for index, server := range conf.Servers {
		cmdChans[index] = makeConnection(server)
	}

	requestChan = cmdChan

	go handleCommands(conf.Servers)
}

func makeConnection(server string) chan *string {
	fmt.Println("Making connection with ", server)
	conn, err := net.Dial("tcp", server)

	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	cmdChan := make(chan *string)
	go handleRequests(conn, cmdChan)

	return cmdChan
}

func handleRequests(conn net.Conn, cmdChan chan *string) {
	fmt.Println("Handling requests")
	for {
		cmd := <-cmdChan
		fmt.Println("Command Received", *cmd)
		fmt.Fprintf(conn, *cmd)
	}
}

func handleCommands(servers []string) {
	fmt.Println("Handling Commands")
	i := 0
	for {
		cmd := <-requestChan
		cmdChans[i] <- cmd

		fmt.Println("Sending ", cmd, " to ", i)

		i++
		if i == len(servers) {
			i = 0
		}
	}
}
