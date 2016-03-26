package main

import (
  "./foobar"
  "github.com/hashicorp/terraform/plugin"
)

func main() {
  plugin.Serve(&plugin.ServeOpts{
    ProviderFunc: foobar.Provider,
  })
}
