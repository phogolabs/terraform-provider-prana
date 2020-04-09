package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
	"github.com/phogolabs/terraform-provider-prana/prana"

	_ "github.com/lib/pq"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return prana.NewProvider()
		},
	})
}
