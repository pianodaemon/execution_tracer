package collectors

import (
	"net"
	"strconv"
	"time"
)

type fluentCollector struct {
	host         string
	port         int
	conn         net.Conn
	writeTimeout time.Time
}

func (fluentPtr *fluentCollector) connect() error {

	const (
		defaultHost    = "127.0.0.1"
		defaultPort    = 24224
		defaultNet     = "tcp"
		defaultTimeout = 3 * time.Second
	)

	target := func() {

		if (*fluentPtr).host == "" {
			(*fluentPtr).host = defaultHost
		}

		if (*fluentPtr).port == 0 {
			(*fluentPtr).port = defaultPort
		}

		portStr = strconv.Itoa((*fluentPtr).port)

		return (*fluentPtr).host + ":" + portStr
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

func (fluentPtr *fluentCollector) send(buffPtr *[]byte) error {

	var err error = nil

	{
		const defaultWriteTimeout = time.Duration(0)
		t := time.Time{}

		if defaultWriteTimeout < (*fluentPtr).writeTimeout {
			t = (*fluentPtr).writeTimeout
		}

		(*fluentPtr).conn.SetWriteDeadline(t)
	}

	_, err = (*fluentPtr).conn.Write(*buffPtr)

	return err
}
