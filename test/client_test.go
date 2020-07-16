package test

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

        hclog "github.com/hashicorp/go-hclog"

	"github.com/hashicorp/go-plugin"
	"github.com/manojkva/go-redfish-plugin/common"
	//"github.com/manojkva/go-redfish-plugin/pkg/drivers/redfish"
)

func TestClientRequest(t *testing.T) {
	logger := hclog.New(&hclog.LoggerOptions{
		  Name: "plugin",
		  Output: os.Stdout,
		  Level: hclog.Debug,})
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig:  common.Handshake,
		Plugins:          common.PluginMap,
		Cmd:              exec.Command("sh", "-c", "../go-redfish-plugin"),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
	        Logger: logger,})
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
	service := raw.(common.Redfish)
	//service := raw.(redfish.BMHNode)
  x, err := service.GetGUUID()
  fmt.Printf("%v\n", string(x))
}
