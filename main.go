package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/pasqualet/terraform-provider-gogs/gogs"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: gogs.Provider})
}
