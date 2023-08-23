// Code generated by protoc-gen-go-drpc. DO NOT EDIT.
// protoc-gen-go-drpc version: v0.0.32
// source: nodesync/nodesyncproto/protos/nodesync.proto

package nodesyncproto

import (
	bytes "bytes"
	context "context"
	errors "errors"
	jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	drpc "storj.io/drpc"
	drpcerr "storj.io/drpc/drpcerr"
)

type drpcEncoding_File_nodesync_nodesyncproto_protos_nodesync_proto struct{}

func (drpcEncoding_File_nodesync_nodesyncproto_protos_nodesync_proto) Marshal(msg drpc.Message) ([]byte, error) {
	return proto.Marshal(msg.(proto.Message))
}

func (drpcEncoding_File_nodesync_nodesyncproto_protos_nodesync_proto) Unmarshal(buf []byte, msg drpc.Message) error {
	return proto.Unmarshal(buf, msg.(proto.Message))
}

func (drpcEncoding_File_nodesync_nodesyncproto_protos_nodesync_proto) JSONMarshal(msg drpc.Message) ([]byte, error) {
	var buf bytes.Buffer
	err := new(jsonpb.Marshaler).Marshal(&buf, msg.(proto.Message))
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (drpcEncoding_File_nodesync_nodesyncproto_protos_nodesync_proto) JSONUnmarshal(buf []byte, msg drpc.Message) error {
	return jsonpb.Unmarshal(bytes.NewReader(buf), msg.(proto.Message))
}

type DRPCNodeSyncClient interface {
	DRPCConn() drpc.Conn

	PartitionSync(ctx context.Context, in *PartitionSyncRequest) (*PartitionSyncResponse, error)
	ColdSync(ctx context.Context, in *ColdSyncRequest) (DRPCNodeSync_ColdSyncClient, error)
}

type drpcNodeSyncClient struct {
	cc drpc.Conn
}

func NewDRPCNodeSyncClient(cc drpc.Conn) DRPCNodeSyncClient {
	return &drpcNodeSyncClient{cc}
}

func (c *drpcNodeSyncClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcNodeSyncClient) PartitionSync(ctx context.Context, in *PartitionSyncRequest) (*PartitionSyncResponse, error) {
	out := new(PartitionSyncResponse)
	err := c.cc.Invoke(ctx, "/anyNodeSync.NodeSync/PartitionSync", drpcEncoding_File_nodesync_nodesyncproto_protos_nodesync_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcNodeSyncClient) ColdSync(ctx context.Context, in *ColdSyncRequest) (DRPCNodeSync_ColdSyncClient, error) {
	stream, err := c.cc.NewStream(ctx, "/anyNodeSync.NodeSync/ColdSync", drpcEncoding_File_nodesync_nodesyncproto_protos_nodesync_proto{})
	if err != nil {
		return nil, err
	}
	x := &drpcNodeSync_ColdSyncClient{stream}
	if err := x.MsgSend(in, drpcEncoding_File_nodesync_nodesyncproto_protos_nodesync_proto{}); err != nil {
		return nil, err
	}
	if err := x.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DRPCNodeSync_ColdSyncClient interface {
	drpc.Stream
	Recv() (*ColdSyncResponse, error)
}

type drpcNodeSync_ColdSyncClient struct {
	drpc.Stream
}

func (x *drpcNodeSync_ColdSyncClient) Recv() (*ColdSyncResponse, error) {
	m := new(ColdSyncResponse)
	if err := x.MsgRecv(m, drpcEncoding_File_nodesync_nodesyncproto_protos_nodesync_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *drpcNodeSync_ColdSyncClient) RecvMsg(m *ColdSyncResponse) error {
	return x.MsgRecv(m, drpcEncoding_File_nodesync_nodesyncproto_protos_nodesync_proto{})
}

type DRPCNodeSyncServer interface {
	PartitionSync(context.Context, *PartitionSyncRequest) (*PartitionSyncResponse, error)
	ColdSync(*ColdSyncRequest, DRPCNodeSync_ColdSyncStream) error
}

type DRPCNodeSyncUnimplementedServer struct{}

func (s *DRPCNodeSyncUnimplementedServer) PartitionSync(context.Context, *PartitionSyncRequest) (*PartitionSyncResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCNodeSyncUnimplementedServer) ColdSync(*ColdSyncRequest, DRPCNodeSync_ColdSyncStream) error {
	return drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

type DRPCNodeSyncDescription struct{}

func (DRPCNodeSyncDescription) NumMethods() int { return 2 }

func (DRPCNodeSyncDescription) Method(n int) (string, drpc.Encoding, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/anyNodeSync.NodeSync/PartitionSync", drpcEncoding_File_nodesync_nodesyncproto_protos_nodesync_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCNodeSyncServer).
					PartitionSync(
						ctx,
						in1.(*PartitionSyncRequest),
					)
			}, DRPCNodeSyncServer.PartitionSync, true
	case 1:
		return "/anyNodeSync.NodeSync/ColdSync", drpcEncoding_File_nodesync_nodesyncproto_protos_nodesync_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return nil, srv.(DRPCNodeSyncServer).
					ColdSync(
						in1.(*ColdSyncRequest),
						&drpcNodeSync_ColdSyncStream{in2.(drpc.Stream)},
					)
			}, DRPCNodeSyncServer.ColdSync, true
	default:
		return "", nil, nil, nil, false
	}
}

func DRPCRegisterNodeSync(mux drpc.Mux, impl DRPCNodeSyncServer) error {
	return mux.Register(impl, DRPCNodeSyncDescription{})
}

type DRPCNodeSync_PartitionSyncStream interface {
	drpc.Stream
	SendAndClose(*PartitionSyncResponse) error
}

type drpcNodeSync_PartitionSyncStream struct {
	drpc.Stream
}

func (x *drpcNodeSync_PartitionSyncStream) SendAndClose(m *PartitionSyncResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_nodesync_nodesyncproto_protos_nodesync_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCNodeSync_ColdSyncStream interface {
	drpc.Stream
	Send(*ColdSyncResponse) error
}

type drpcNodeSync_ColdSyncStream struct {
	drpc.Stream
}

func (x *drpcNodeSync_ColdSyncStream) Send(m *ColdSyncResponse) error {
	return x.MsgSend(m, drpcEncoding_File_nodesync_nodesyncproto_protos_nodesync_proto{})
}
