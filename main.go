package main

import (
  "github.com/jbussdieker/terraform-provider-foobar/foobar"
  "github.com/hashicorp/terraform/plugin"
)

func main() {
  plugin.Serve(&plugin.ServeOpts{
    ProviderFunc: foobar.Provider,
  })
}
