package rmq

import (
	"github.com/thuongtruong1009/zoomer/pkg/queue/rmq/client"
	"testing"
)

const (
	host              = "localhost:8080"
	rmqURL            = "amqp://guest:guest@localhost:5672/"
	rpcServerExchange = "rpc_server"
	rpcClientExchange = "rpc_client"
	requests          = 10
)

func TestRMQClient(t *testing.T) {
	rmqClient, err := client.New(rmqURL, rpcServerExchange, rpcClientExchange)
	if err != nil {
		t.Fatal("Failed to create RMQ client")
	}

	defer func() {
		err = rmqClient.Shutdown()
		if err != nil {
			t.Fatal("Failed to shutdown RMQ client")
		}
	}()

	type argsTest struct {
		A int `json:"a"`
		B int `json:"b"`
	}

	type sumReplyTest struct {
		Sum []argsTest `json:"sum"`
	}

	// ref: https://github.dev/evrone/go-clean-template
	for i := 0; i < requests; i++ {
		var sum sumReplyTest

		err = rmqClient.RemoteCall("sum", nil, &sum)
		if err != nil {
			t.Fatal("Failed to call remote function")
		}

		if sum.Sum[0].A+sum.Sum[0].B != i+i {
			t.Fatal("Failed to get correct sum")
		}
	}
}
