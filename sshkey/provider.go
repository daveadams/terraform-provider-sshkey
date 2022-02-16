package sshkey

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type sshkeyProvider struct {
	configured bool
}

func Provider() tfsdk.Provider {
	return &sshkeyProvider{}
}

func (p *sshkeyProvider) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			// TODO: remove this as of tfsdk 0.6.0
			// see https://github.com/hashicorp/terraform-plugin-framework/issues/250
			"noop": {
				Type:     types.StringType,
				Optional: true,
			},
		},
	}, nil
}

type sshkeyProviderData struct{}

func (p *sshkeyProvider) Configure(ctx context.Context, req tfsdk.ConfigureProviderRequest, resp *tfsdk.ConfigureProviderResponse) {
	p.configured = true
}

func (p *sshkeyProvider) GetResources(_ context.Context) (map[string]tfsdk.ResourceType, diag.Diagnostics) {
	return map[string]tfsdk.ResourceType{
		"sshkey_rsa_key_pair":     resourceRSAKeyPairType{},
		"sshkey_ed25519_key_pair": resourceED25519KeyPairType{},
	}, nil
}

func (p *sshkeyProvider) GetDataSources(ctx context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
	return map[string]tfsdk.DataSourceType{}, nil
}
