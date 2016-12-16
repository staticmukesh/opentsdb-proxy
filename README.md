# opentsdb-proxy
TCP proxy for opentsdb servers

### Usages

You can pass config file using `-conf` parameter to the proxy. Currently following fields are supported.

`Host = ":8080"`    
`Servers = ["localhost:4242","localhost:4241"] `

To run the project, use the following steps  
`
	go build
`  
`./opentsdb-proxy -conf=opentsdb.conf
`

### Features

- [x] Multiple Opentsdb support
- [ ] Limit number of connection
- [ ] Junk filter for rpcs
- [ ] Configurable number of tsdb connections
- [ ] Support for http requests
- [ ] Heart beat for Opentsdbs
