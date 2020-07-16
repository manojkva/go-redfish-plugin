package common

import (
	"context"
	"github.com/hashicorp/go-plugin"
	"github.com/manojkva/go-redfish-plugin/proto"
	"google.golang.org/grpc"
)

var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "REDFISH_PLUGIN",
	MagicCookieValue: "1.0",
}

var PluginMap = map[string]plugin.Plugin{
	"redfish": &RedfishPlugin{},
}

type Redfish interface {
	GetGUUID() ([]byte, error)
	DeployISO() (error)
	UpdateFirmware() (error)
	ConfigureRAID()  (error)
}

type RedfishPlugin struct {
	Impl Redfish
	plugin.Plugin
}

func (p *RedfishPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterRedfishServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *RedfishPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: proto.NewRedfishClient(c)}, nil

}
