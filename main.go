package main

import (
	"github.com/higebu/packer-builder-vmware-esxi/builder/vmware/esxi"
	"github.com/mitchellh/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(esxi.Builder))
	server.Serve()
}
