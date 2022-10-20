package rpctest

import (
	"context"
	"net"
	"storj.io/drpc"
	"storj.io/drpc/drpcconn"
	"storj.io/drpc/drpcmux"
	"storj.io/drpc/drpcserver"
)

func NewTestServer() *TesServer {
	ts := &TesServer{
		Mux: drpcmux.New(),
	}
	ts.Server = drpcserver.New(ts.Mux)
	return ts
}

type TesServer struct {
	*drpcmux.Mux
	*drpcserver.Server
}

func (ts *TesServer) Dial() drpc.Conn {
	sc, cc := net.Pipe()
	go ts.Server.ServeOne(context.Background(), sc)
	return drpcconn.New(cc)
}
