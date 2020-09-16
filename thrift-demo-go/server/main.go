package main

import (
	"codelab/thrift-demo-go/gen-go/thriftrpc"
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"os"
)

const NetworkAddr = "127.0.0.1:19090"

var _ thriftrpc.Calculator = (*CalculatorService)(nil)

type CalculatorService struct{}

func (c *CalculatorService) Add(ctx context.Context, a int32, b int32) (r int32, err error) {
	return a + b, nil
}

func main() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTServerSocket(NetworkAddr)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}
	fmt.Printf("%T\n", transport)

	handler := &CalculatorService{}
	processor := thriftrpc.NewCalculatorProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	fmt.Println("Starting the simple server... on ", NetworkAddr)
	server.Serve()
}
