package common

import (
	"github.com/manojkva/go-redfish-plugin/proto"
	"golang.org/x/net/context"
)

type GRPCClient struct{ client proto.RedfishClient }

func (m *GRPCClient) GetGUUID() ([]byte, error) {
	resp, err := m.client.GetGUUID(context.Background(), &proto.Empty{})
	if err != nil {
		return nil, err
	}

	return resp.Value, nil
}

func (m *GRPCClient)  UpdateFirmware() error {
	 return nil
}
func (m *GRPCClient)  ConfigureRAID() error {
	 return nil
}
func (m *GRPCClient)  DeployISO() error {
	 return nil
}

type GRPCServer struct {
	Impl Redfish
}

func (m *GRPCServer) GetGUUID(ctx context.Context, req *proto.Empty) (*proto.Response, error) {
	v, err := m.Impl.GetGUUID()
	return &proto.Response{Value: v}, err
}
func (m *GRPCServer)  UpdateFirmware(ctx context.Context, req *proto.Empty) (*proto.Empty,error) {
	 return nil,nil
}
func (m *GRPCServer)  ConfigureRAID(ctx context.Context, req *proto.Empty) (*proto.Empty,error) {
	 return nil,nil
}
func (m *GRPCServer)  DeployISO(ctx context.Context, req *proto.Empty) (*proto.Empty,error) {
	 return nil,nil
}
