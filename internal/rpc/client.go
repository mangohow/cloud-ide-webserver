package rpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sync"
	"time"
)

var (
	clients = map[string]*grpc.ClientConn{}
	lock    sync.Mutex
)

func GrpcClient(name string) *grpc.ClientConn {
	lock.Lock()
	defer lock.Unlock()
	if c, ok := clients[name]; ok {
		return c
	}

	conn := newClient()
	clients[name] = conn

	return conn
}

func newClient() *grpc.ClientConn {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	conn, err := grpc.DialContext(ctx, "192.168.44.100:6387", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	return conn
}