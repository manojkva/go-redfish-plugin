package main

import (
	"encoding/json"
	"fmt"
	"github.com/bm-metamorph/MetaMorph/pkg/db/models/node"
	"github.com/hashicorp/go-plugin"
	"github.com/manojkva/go-redfish-plugin/common"
	"github.com/manojkva/go-redfish-plugin/pkg/drivers/redfish"
	"os"
)

func main() {
	//<TBD> check if stdout need to be restrained..
	//fmt.Printf("%v\n", os.Args[1])
	//fmt.Printf("%v\n", len(os.Args))
	if len(os.Args) != 2 {
		fmt.Println("Usage go-redfish-plugin <uuid>")
		os.Exit(1)
	}
	uuid := os.Args[1]

	//fmt.Printf("uuid %v\n", uuid)

	var bmhnode redfish.BMHNode

	old  :=  os.Stdout

	os.Stdout,_ = os.Open(os.DevNull)

	data, err := node.Describe(uuid)

	if err == nil {

		err = json.Unmarshal(data, &bmhnode)
	}
	if err != nil {

		fmt.Printf("Failed to locate node in DB for uuid %v\n", uuid)
		os.Exit(1)
	}
	os.Stdout = old
	//Get node details from db
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: common.Handshake,
		Plugins: map[string]plugin.Plugin{
//					"redfish": &common.RedfishPlugin{Impl: &redfish.BMHNode{}}},
			"redfish": &common.RedfishPlugin{Impl: &bmhnode}},
		GRPCServer: plugin.DefaultGRPCServer})
}
