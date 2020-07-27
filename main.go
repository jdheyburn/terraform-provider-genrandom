package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/jdheyburn/terraform-provider-template/genrandom"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: genrandom.Provider})
}
