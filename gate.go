package main

import (
	"github.com/TinyBrickBoy/ip-blacklist/plugins/blacklist"
	"go.minekube.com/gate/cmd/gate"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

func main() {
    proxy.Plugins = append(proxy.Plugins,
        ipblacklist.Plugin,
    )
    gate.Execute()
}
