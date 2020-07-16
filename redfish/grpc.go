package redfish

import (
	"github.com/manojkva/go-redfish-plugin/proto"
	"golang.org/x/net/context"
)

type GRPCClient struct{ client proto.RedfishClient }

func (m *GRPCClient) GetGUIID() ([]byte, error) {
	resp, err := m.client.GetGUIID(context.Background(), &proto.Empty{})
	if err != nil {
		return nil, err
	}

	return resp.Value, nil
}

type GRPCServer struct {
	Impl Redfish
}

func (m *GRPCServer) GetGUIID(ctx context.Context, req *proto.Empty) (*proto.Response, error) {
	v, err := m.Impl.GetGUIID()
	return &proto.Response{Value: v}, err
}
