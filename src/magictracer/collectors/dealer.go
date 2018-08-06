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

func (dco Dealer) Connect(adc Collector) {
	err = adc.connect()
	return err
}

func (dco Dealer) Disconnect(adc Collector) {

	dco.muconn.Lock()

	adc.disconnect()

	dco.muconn.Unlock()
}

func (dco Dealer) Send(adc Collector, buffPtr *[]byte) error {

	attemptSend := func(wasteTimeEvent func(int)) error {

		for i := 0; i < dco.maxRetry; i++ {

			dco.muconn.Lock()

			err := adc.assureConn()

			if err != nil {
				dco.muconn.Unlock()
				wasteTimeEvent(i)
				continue
			}

			dco.muconn.Unlock()

			err = adc.send(buffPtr)

			if err != nil {
				adc.disconnect()
			} else {
				return err
			}
		}

		return fmt.Errof("Failed to reconnect, max retry: %v",
			dco.maxRetry)
	}

	return attemptSend(func(i int) {

		const (
			defaultReconnWaitInc = 1.5
		)

		rate := int(math.Pow(defaultReconnWaitInc, float64(i-1)))

		wTime := dco.RetryWait * rate

		if wTime > dco.MaxRetryWait {
			wTime = dco.MaxRetryWait
		}

		time.Sleep(time.Duration(wTime) * time.Millisecond)
	})
}
