package collectors

import (
	"sync"
)

type Collector interface {
	connect() error
	disconnect()
	send() error
}

type Dealer struct {
	muconn sync.Mutex
}

func (dcol Dealer) Connect(adc Collector) {
	err = dcol.connect()
	return err
}

func (dcol Dealer) Disconnect(adc Collector) {
	muconn.Lock()
	dcol.disconnect()
	muconn.Unlock()
}
