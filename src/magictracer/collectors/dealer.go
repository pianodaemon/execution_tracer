package collectors

import (
	"fmt"
	"sync"
	"time"
)

type Collector interface {
	connect() error
	disconnect()
	send() error
	assureConn() error
}

type Dealer struct {
	muconn sync.Mutex
}

func (dcol Dealer) Connect(adc Collector) {
	err = adc.connect()
	return err
}

func (dcol Dealer) Disconnect(adc Collector) {

	dcol.muconn.Lock()

	adc.disconnect()

	dcol.muconn.Unlock()
}

func (dcol Dealer) Send(adc Collector, buffPtr *[]byte) error {

	attemptSend := func(wasteTimeEvent func(int)) error {

		for i := 0; i < dcol.maxRetry; i++ {

			dcol.muconn.Lock()

			err := adc.assureConn()

			if err != nil {
				dcol.muconn.Unlock()
				wasteTimeEvent(i)
				continue
			}

			dcol.muconn.Unlock()

			err = adc.send(buffPtr)

			if err != nil {
				adc.disconnect()
			} else {
				return err
			}
		}

		return fmt.Errof("Failed to reconnect, max retry: %v",
			dcol.maxRetry)
	}

	return attemptSend(func(i int) {

		const (
			defaultReconnWaitInc = 1.5
		)

		rate := int(math.Pow(defaultReconnWaitInc, float64(i-1)))

		wTime := dcol.RetryWait * rate

		if wTime > dcol.MaxRetryWait {
			wTime = dcol.MaxRetryWait
		}

		time.Sleep(time.Duration(wTime) * time.Millisecond)
	})
}
