package main

import (
	"flag"
	"fmt"

	"github.com/staticmukesh/opentsdb-proxy/conf"
)

var flagConf = flag.String("conf", "opentsdb.conf", "Location of configuration file. Defaults to opentsdb.toml in directory of the proxy executable.")

func main() {
	flag.Parse()

	conf := conf.ReadConf(flagConf)

	fmt.Println(conf)
}
