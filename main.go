package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"

	"github.com/daveadams/terraform-provider-sshkey/sshkey"
)

func main() {
	opts := tfsdk.ServeOpts{
		Name: "sshkey",
	}
	tfsdk.Serve(context.Background(), sshkey.Provider, opts)
}
