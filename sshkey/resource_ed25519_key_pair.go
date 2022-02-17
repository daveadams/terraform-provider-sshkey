package sshkey

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type resourceED25519KeyPairType struct{}

func (r resourceED25519KeyPairType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:     types.StringType,
				Computed: true,
			},
			"comment": {
				Type:     types.StringType,
				Optional: true,
			},
			"private_key_pem": {
				Type:      types.StringType,
				Computed:  true,
				Sensitive: true,
			},
			"public_key": {
				Type:     types.StringType,
				Computed: true,
			},
			"fingerprint_md5": {
				Type:     types.StringType,
				Computed: true,
			},
			"fingerprint_sha256": {
				Type:     types.StringType,
				Computed: true,
			},
		},
	}, nil
}

func (r resourceED25519KeyPairType) NewResource(ctx context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceED25519KeyPair{}, nil
}

type resourceED25519KeyPair struct{}

func (r resourceED25519KeyPair) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var plan ED25519KeyPair
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	key, err := generateED25519KeyPair()
	if err != nil {
		resp.Diagnostics.AddError("Unable to generate a new ED25519 key", err.Error())
		return
	}

	privateKeyPEM, err := key.PrivateKeyPEM()
	if err != nil {
		resp.Diagnostics.AddError("Unable to render private key to PEM format", err.Error())
		return
	}

	// NOTE: Leaves a trailing space in the public key content if no comment is
	// specified, to match the behavior of `ssh-keygen -t ed25519 -C ''` and avoid
	// validation failure by AWS.
	renderedPublicKey := fmt.Sprintf("%s %s", key.PublicKey(), plan.Comment.Value)
	result := ED25519KeyPair{
		ID:                types.String{Value: key.FingerprintSHA256()},
		Comment:           plan.Comment,
		PrivateKeyPEM:     types.String{Value: privateKeyPEM},
		PublicKey:         types.String{Value: renderedPublicKey},
		FingerprintMD5:    types.String{Value: key.FingerprintMD5()},
		FingerprintSHA256: types.String{Value: key.FingerprintSHA256()},
	}

	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r resourceED25519KeyPair) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	// no need to support Read at the moment since the resource is fully within state
	// NOTE: if we support sourcing from files on disk in the future, this will have to be implemented
}

func (r resourceED25519KeyPair) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	// no need to support Update at the moment, since the only possible changes require replacement
	// NOTE: if we need to support changes later for comments, etc, this will need to be implemented
}

func (r resourceED25519KeyPair) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var state ED25519KeyPair
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r resourceED25519KeyPair) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStateNotImplemented(ctx, "", resp)
}
