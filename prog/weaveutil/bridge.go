package main

import (
	"fmt"
	"strconv"

	"github.com/weaveworks/weave/common"
	weavenet "github.com/weaveworks/weave/net"
)

func detectBridgeType(args []string) error {
	if len(args) != 2 {
		cmdUsage("detect-bridge-type", "<weave-bridge-name> <datapath-name>")
	}
	bridgeType := weavenet.DetectBridgeType(args[0], args[1])
	fmt.Println(bridgeType.String())
	return nil
}

func createBridge(args []string) error {
	if len(args) != 7 {
		cmdUsage("create-bridge", "<docker-bridge-name> <weave-bridge-name> <datapath-name> <mtu> <port> <no-fastdp> <no-bridged-fastdp>")
	}

	mtu, err := strconv.Atoi(args[3])
	if err != nil && args[3] != "" {
		return err
	}
	port, err := strconv.Atoi(args[4])
	if err != nil {
		return err
	}
	config := common.BridgeConfig{
		DockerBridgeName: args[0],
		WeaveBridgeName:  args[1],
		DatapathName:     args[2],
		MTU:              mtu,
		Port:             port,
		NoFastdp:         args[5] != "",
		NoBridgedFastdp:  args[6] != "",
	}
	bridgeType, err := common.CreateBridge(&config)
	fmt.Println(bridgeType.String())
	return err
}

// TODO: destroy-bridge

func enforceAddrAsign(args []string) error {
	if len(args) != 1 {
		cmdUsage("enforce-bridge-addr-assign-type", "<bridge-name>")
	}
	return common.EnforceAddrAssignType(args[0])
}