package main

import (
	"github.com/hashicorp/go-plugin"
        "github.com/manojkva/go-redfish-plugin/common"
        "github.com/manojkva/go-redfish-plugin/pkg/drivers/redfish"
)


func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: common.Handshake,
		Plugins: map[string]plugin.Plugin{
			"redfish": &common.RedfishPlugin{Impl: &redfish.BMHNode{}}},
		GRPCServer: plugin.DefaultGRPCServer})
}
