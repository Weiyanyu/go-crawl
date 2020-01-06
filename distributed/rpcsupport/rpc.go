package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServeRpc(host string, rpcService interface{}) error {
	rpc.Register(rpcService)
	listenr, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}
	log.Printf("server %s already listen", host)
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
	client := jsonrpc.NewClient(conn)
	return client, err
}
