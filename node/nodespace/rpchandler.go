package nodespace

import (
	"context"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/common/commonspace/spacesyncproto"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/common/commonspace/storage"
)

type rpcHandler struct {
	s *service
}

func (r *rpcHandler) PushSpace(ctx context.Context, req *spacesyncproto.PushSpaceRequest) (resp *spacesyncproto.PushSpaceResponse, err error) {
	_, err = r.s.GetSpace(ctx, req.SpaceId)
	if err == nil {
		resp = &spacesyncproto.PushSpaceResponse{}
		return
	}
	payload := storage.SpaceStorageCreatePayload{
		RecWithId:   req.AclRoot,
		SpaceHeader: req.SpaceHeader,
		Id:          req.SpaceId,
	}
	_, err = r.s.spaceStorageProvider.CreateSpaceStorage(payload)
	if err != nil {
		err = spacesyncproto.ErrUnexpected
		return
	}
	return
}

func (r *rpcHandler) HeadSync(ctx context.Context, req *spacesyncproto.HeadSyncRequest) (*spacesyncproto.HeadSyncResponse, error) {
	sp, err := r.s.GetSpace(ctx, req.SpaceId)
	if err != nil {
		return nil, spacesyncproto.ErrSpaceMissing
	}
	return sp.SpaceSyncRpc().HeadSync(ctx, req)
}

func (r *rpcHandler) Stream(stream spacesyncproto.DRPCSpace_StreamStream) error {
	msg, err := stream.Recv()
	if err != nil {
		return err
	}
	sp, err := r.s.GetSpace(stream.Context(), msg.SpaceId)
	if err != nil {
		return spacesyncproto.ErrSpaceMissing
	}
	return sp.SpaceSyncRpc().Stream(stream)
}
