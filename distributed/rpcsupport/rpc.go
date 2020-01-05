package rpcsupport

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServeRpc(host string, rpcService interface{}, serverNotifier chan struct{}) error {
	rpc.Register(rpcService)
	listenr, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}
	serverNotifier <- struct{}{}
	for {
		conn, err := listenr.Accept()
		if err != nil {
			log.Printf("error: %v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}

func NewRpcClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	fmt.Println("")
	client := jsonrpc.NewClient(conn)
	return client, err
}