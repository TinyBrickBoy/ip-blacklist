package main

import (
	"fmt"
	"github.com/TinyBrickBoy/ip-blacklist/plugins/blacklist"
	"go.minekube.com/gate/cmd/gate"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

func main() {
	fmt.Println("Starting Minekube Gate with IP Blacklist Plugin...")
	proxy.Plugins = append(proxy.Plugins, blacklist.Plugin)
	gate.Execute()
}
