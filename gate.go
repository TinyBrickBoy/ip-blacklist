package main

import (
	"github.com/TinyBrickBoy/ip-blacklist/plugins/ipblacklist"
	"go.minekube.com/gate/cmd/gate"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

func main() {
    // Hier registrieren wir unsere Plugins mit dem Proxy.
    proxy.Plugins = append(proxy.Plugins,
        // Andere Plugins...
        ipblacklist.Plugin,
    )

    // Ausführen von Gate...
    gate.Execute()
}
