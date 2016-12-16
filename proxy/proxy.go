package proxy

import (
	"fmt"
	"net"

	"github.com/staticmukesh/opentsdb-proxy/conf"
)

func Init(conf *conf.Conf, cmds chan *string) {
	fmt.Println("Proxy Init")
	for _, server := range conf.Servers {
		go handleRequests(server, cmds)
	}
}

func handleRequests(server string, cmds chan *string) {
	fmt.Println("Making connection with ", server)
	conn, err := net.Dial("tcp", server)

	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	defer conn.Close()

	for {
		select {
		case cmd := <-cmds:
			fmt.Println("Command Received", *cmd)
			fmt.Fprintf(conn, *cmd)
		}
	}
}
