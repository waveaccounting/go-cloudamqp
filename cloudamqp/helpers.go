package cloudamqp

import (
	"errors"
	"net"
	"net/url"
	"time"
)

// checkInstanceUntilAvailable Blocks execution until the instance
// is fully operational
func checkInstanceUntilAvailable(instance *Instance, timeoutInSeconds float64) error {
	instanceURL, _ := url.Parse(instance.URL)
	initialTime := time.Now()

	var network string

	if instanceURL.Scheme != "amqp" {
		network = "tcp"
	}

	for {
		conn, err := net.DialTimeout(
			network,
			instanceURL.Host,
			time.Second,
		)

		if time.Now().Sub(initialTime).Seconds() > timeoutInSeconds {
			return errors.New("Max timeout reached")
		}

		if err == nil {
			conn.Close()
			return nil
		}
	}
}
