package collectors


import (
	"net"
	"strconv"
	"time"
)

type fluentCollector struct {
	conn net.Conn
}

func (fluentPtr *fluentCollector) connect(host string, port uint16) error {

	const (
		defaultHost = "127.0.0.1"
		defaultPort = 24224
		defaultNet  = "tcp"
		defaultTimeout = 3 * time.Second
	)

	target := func () {

		if host == "":
			host = defaultHost
		if port == 0:
			port = defaultPort

		return host + ":" + strconv.Itoa(port)
	}

	var err error = nil

	(*fluentPtr).conn, err = net.DialTimeout(
		defaultNet, target(), defaultTimeout)

	return err
}


func (fluentPtr *fluentCollector) disconnect() {

	if (*fluentPtr).conn != nil {
		(*fluentPtr).conn.Close()
		(*fluentPtr).conn = nil
	}
}
