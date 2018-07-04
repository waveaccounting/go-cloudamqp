package cloudamqp

import (
	"fmt"
	"log"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func rabbitMQTestServer() net.Listener {
	startPort := 5672

	var listener net.Listener
	var errListener error

	for {
		listener, errListener = net.Listen(
			"tcp",
			fmt.Sprintf("127.0.0.1:%d", startPort),
		)

		if errListener != nil {
			startPort++
			continue
		}

		return listener
	}
}

func checkUnusedPort() int {
	port := 10000

	for {
		conn, err := net.DialTimeout(
			"tcp",
			fmt.Sprintf("127.0.0.1:%d", port),
			time.Millisecond*10,
		)

		if err == nil {
			conn.Close()
			port++
			continue
		}

		return port
	}
}

func TestCheckInstanceUntilAvailableFailture(t *testing.T) {
	unusedPort := checkUnusedPort()

	instance := Instance{
		URL: fmt.Sprintf(
			"tcp://127.0.0.1:%d",
			unusedPort,
		),
	}

	resultErr := checkInstanceUntilAvailable(
		&instance,
		float64(1),
	)

	assert.Error(t, resultErr)
}

func TestCheckInstanceUntilAvailableSuccess(t *testing.T) {
	listener := rabbitMQTestServer()

	go func() {
		for {
			conn, errAccept := listener.Accept()

			if errAccept != nil {
				log.Panicln(errAccept)
			}

			conn.Write([]byte("OK"))
		}
	}()

	instance := Instance{
		URL: fmt.Sprintf(
			"%s://%s",
			listener.Addr().Network(),
			listener.Addr().String(),
		),
	}

	resultErr := checkInstanceUntilAvailable(
		&instance,
		float64(1),
	)

	assert.NoError(t, resultErr)
}
