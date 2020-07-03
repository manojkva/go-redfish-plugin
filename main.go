package main

import (
	"github.com/hashicorp/go-plugin"
        "github.com/manojkva/go-redfish-plugin/redfish"
)

type Redfish struct{}

func (Redfish) GetGUIID() ([]byte, error) {
	// Write the logic of handing redfish wrapper code here
	return []byte("test"), nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: redfish.Handshake,
		Plugins: map[string]plugin.Plugin{
			"redfish": &redfish.RedfishPlugin{Impl: &Redfish{}}},
		GRPCServer: plugin.DefaultGRPCServer})
}
