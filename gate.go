package main

import (
	"github.com/minekube/gate-plugin-template/plugins/ipblacklist" // Local import path
	"go.minekube.com/gate/cmd/gate"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

// Es ist ein normales Go-Programm, wir müssen nur
// unsere Plugins registrieren und Gate ausführen.
func main() {
	// Hier registrieren wir unsere Plugins mit dem Proxy.
	proxy.Plugins = append(proxy.Plugins,
		// Unser IP Blacklist Plugin
		ipblacklist.Plugin,

		// Füge weitere Plugins hinzu, falls benötigt
	)

	// Führe Gate aus, als wäre es ein normales Go-Programm.
	// Gate kümmert sich um den Rest für uns,
	// wie Config auto-reloading und Flags wie --debug.
	gate.Execute()
}
