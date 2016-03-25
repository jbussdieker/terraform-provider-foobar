package main

import (
  "github.com/jbussdieker/terraform-plugin-test/foobar"
  "github.com/hashicorp/terraform/plugin"
)

func main() {
  plugin.Serve(&plugin.ServeOpts{
    ProviderFunc: btmp.Provider,
  })
}
