package main

import (
	"codelab/thrift-demo-go/gen-go/thriftrpc"
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

const NetworkAddr = "127.0.0.1:19090"

func main() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket(NetworkAddr)
	check(err)
	fmt.Printf("%T\n", transport)

	transport, err = transportFactory.GetTransport(transport)
	check(err)
	defer transport.Close()
	check(transport.Open())

	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	client := thriftrpc.NewCalculatorClient(thrift.NewTStandardClient(iprot, oprot))

	ctx := context.Background()
	result, err := client.Add(ctx, 3, 4)
	check(err)
	fmt.Println(result)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
