package secure

import (
	"context"
	commonaccount "github.com/anytypeio/go-anytype-infrastructure-experiments/common/account"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/common/app"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/common/app/logger"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/common/config"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/common/util/keys"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/sec"
	libp2ptls "github.com/libp2p/go-libp2p/p2p/security/tls"
	"go.uber.org/zap"
	"net"
)

type HandshakeError error

const CName = "common.net.secure"

var log = logger.NewNamed(CName)

func New() Service {
	return &service{}
}

type Service interface {
	TLSListener(lis net.Listener) ContextListener
	TLSConn(ctx context.Context, conn net.Conn) (sec.SecureConn, error)
	app.Component
}

type service struct {
	key crypto.PrivKey
}

func (s *service) Init(a *app.App) (err error) {
	account := a.MustComponent(config.CName).(commonaccount.ConfigGetter).GetAccount()
	pkb, err := keys.DecodeBytesFromString(account.PeerKey)
	if err != nil {
		return
	}
	if s.key, err = crypto.UnmarshalEd25519PrivateKey(pkb); err != nil {
		return
	}

	log.Info("secure service init", zap.String("peerId", account.PeerId))

	return nil
}

func (s *service) Name() (name string) {
	return CName
}

func (s *service) TLSListener(lis net.Listener) ContextListener {
	return newTLSListener(s.key, lis)
}

func (s *service) TLSConn(ctx context.Context, conn net.Conn) (sec.SecureConn, error) {
	tr, err := libp2ptls.New(s.key)
	if err != nil {
		return nil, err
	}
	return tr.SecureOutbound(ctx, conn, "")
}
