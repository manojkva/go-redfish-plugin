package test

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/hashicorp/go-plugin"
	"github.com/manojkva/go-redfish-plugin/redfish"
)

func TestClientRequest(t *testing.T) {
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig:  redfish.Handshake,
		Plugins:          redfish.PluginMap,
		Cmd:              exec.Command("sh", "-c", "../go-redfish-plugin"),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC}})
	defer client.Kill()

	rpcClient, err := client.Client()

	if err != nil {
		fmt.Printf("Error %v\n", err)
		os.Exit(1)
	}

	raw, err := rpcClient.Dispense("redfish")
	if err != nil {
		fmt.Printf("Error %v\n", err)
		os.Exit(1)

	}
	service := raw.(redfish.Redfish)
  x, err := service.GetGUIID()
  fmt.Printf("%v\n", string(x))
}
