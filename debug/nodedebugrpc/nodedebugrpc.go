package nodedebugrpc

import (
	"context"
	"github.com/anytypeio/any-sync-node/debug/nodedebugrpc/nodedebugrpcproto"
	"github.com/anytypeio/any-sync-node/nodespace"
	nodestorage "github.com/anytypeio/any-sync-node/nodestorage"
	"github.com/anytypeio/any-sync-node/nodesync"
	"github.com/anytypeio/any-sync/app"
	"github.com/anytypeio/any-sync/app/logger"
	"github.com/anytypeio/any-sync/commonspace/object/treegetter"
	"github.com/anytypeio/any-sync/commonspace/spacestorage"
	"github.com/anytypeio/any-sync/net"
	"github.com/anytypeio/any-sync/net/rpc/server"
	"github.com/anytypeio/any-sync/net/secureservice"
	"github.com/anytypeio/any-sync/nodeconf"
	"storj.io/drpc"
)

const CName = "node.debug.nodedebugrpc"

var log = logger.NewNamed(CName)

func New() NodeDebugRpc {
	return &nodeDebugRpc{BaseDrpcServer: server.NewBaseDrpcServer()}
}

type configGetter interface {
	GetDebugNet() net.Config
}

type NodeDebugRpc interface {
	app.ComponentRunnable
	drpc.Mux
}

type nodeDebugRpc struct {
	transport      secureservice.SecureService
	cfg            net.Config
	treeCache      treegetter.TreeGetter
	spaceService   nodespace.Service
	storageService nodestorage.NodeStorage
	nodeSync       nodesync.NodeSync
	nodeConf       nodeconf.Service
	*server.BaseDrpcServer
}

func (s *nodeDebugRpc) Init(a *app.App) (err error) {
	s.treeCache = a.MustComponent(treegetter.CName).(treegetter.TreeGetter)
	s.spaceService = a.MustComponent(nodespace.CName).(nodespace.Service)
	s.storageService = a.MustComponent(spacestorage.CName).(nodestorage.NodeStorage)
	s.cfg = a.MustComponent("config").(configGetter).GetDebugNet()
	s.transport = a.MustComponent(secureservice.CName).(secureservice.SecureService)
	s.nodeSync = a.MustComponent(nodesync.CName).(nodesync.NodeSync)
	s.nodeConf = a.MustComponent(nodeconf.CName).(nodeconf.Service)
	return nil
}

func (s *nodeDebugRpc) Name() (name string) {
	return CName
}

func (s *nodeDebugRpc) Run(ctx context.Context) (err error) {
	params := server.Params{
		BufferSizeMb:  s.cfg.Stream.MaxMsgSizeMb,
		TimeoutMillis: s.cfg.Stream.TimeoutMilliseconds,
		ListenAddrs:   s.cfg.Server.ListenAddrs,
		Wrapper: func(handler drpc.Handler) drpc.Handler {
			return handler
		},
	}
	err = s.BaseDrpcServer.Run(ctx, params)
	if err != nil {
		return
	}
	return nodedebugrpcproto.DRPCRegisterNodeApi(s, &rpcHandler{s: s})
}

func (s *nodeDebugRpc) Close(ctx context.Context) (err error) {
	return s.BaseDrpcServer.Close(ctx)
}
