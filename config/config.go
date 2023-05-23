package config

import (
	"github.com/anyproto/any-sync-node/nodestorage"
	"github.com/anyproto/any-sync-node/nodesync"
	"github.com/anyproto/any-sync-node/nodesync/hotsync"
	commonaccount "github.com/anyproto/any-sync/accountservice"
	"github.com/anyproto/any-sync/app"
	"github.com/anyproto/any-sync/app/logger"
	"github.com/anyproto/any-sync/commonspace"
	"github.com/anyproto/any-sync/metric"
	"github.com/anyproto/any-sync/net"
	"github.com/anyproto/any-sync/nodeconf"
	"gopkg.in/yaml.v3"
	"os"
)

const CName = "config"

func NewFromFile(path string) (c *Config, err error) {
	c = &Config{}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(data, c); err != nil {
		return nil, err
	}
	return
}

type Config struct {
	GrpcServer               net.Config             `yaml:"grpcServer"`
	Account                  commonaccount.Config   `yaml:"account"`
	APIServer                net.Config             `yaml:"apiServer"`
	Network                  nodeconf.Configuration `yaml:"network"`
	NetworkStorePath         string                 `yaml:"networkStorePath"`
	NetworkUpdateIntervalSec int                    `yaml:"networkUpdateIntervalSec"`
	Space                    commonspace.Config     `yaml:"space"`
	Storage                  nodestorage.Config     `yaml:"storage"`
	Metric                   metric.Config          `yaml:"metric"`
	Log                      logger.Config          `yaml:"log"`
	NodeSync                 nodesync.Config        `yaml:"nodeSync"`
}

func (c Config) Init(a *app.App) (err error) {
	return
}

func (c Config) Name() (name string) {
	return CName
}

func (c Config) GetNet() net.Config {
	return c.GrpcServer
}

func (c Config) GetDebugNet() net.Config {
	return c.APIServer
}

func (c Config) GetAccount() commonaccount.Config {
	return c.Account
}

func (c Config) GetMetric() metric.Config {
	return c.Metric
}

func (c Config) GetSpace() commonspace.Config {
	return c.Space
}

func (c Config) GetStorage() nodestorage.Config {
	return c.Storage
}

func (c Config) GetNodeConf() nodeconf.Configuration {
	return c.Network
}

func (c Config) GetNodeConfStorePath() string {
	return c.NetworkStorePath
}

func (c Config) GetNodeConfUpdateInterval() int {
	return c.NetworkUpdateIntervalSec
}

func (c Config) GetNodeSync() nodesync.Config {
	return c.NodeSync
}

func (c Config) GetHotSync() hotsync.Config {
	return c.NodeSync.HotSync
}
