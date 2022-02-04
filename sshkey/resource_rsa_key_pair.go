package sshkey

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type resourceRSAKeyPairType struct{}

func (r resourceRSAKeyPairType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:     types.StringType,
				Computed: true,
			},
			"bits": {
				Type:     types.NumberType,
				Optional: true,
				Computed: true,
				Validators: []tfsdk.AttributeValidator{
					rsaKeyPairBitsValidator{},
				},
				PlanModifiers: []tfsdk.AttributePlanModifier{
					tfsdk.RequiresReplace(),
				},
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

func (r resourceRSAKeyPairType) NewResource(ctx context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceRSAKeyPair{}, nil
}

type resourceRSAKeyPair struct{}

func (r resourceRSAKeyPair) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var plan RSAKeyPair
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	bits, _ := plan.Bits.Value.Int64()
	key, err := generateRSAKeyPair(int(bits))
	if err != nil {
		resp.Diagnostics.AddError("Unable to generate a new RSA key", err.Error())
		return
	}

	result := RSAKeyPair{
		ID:                types.String{Value: key.FingerprintSHA256()},
		Bits:              plan.Bits,
		PrivateKeyPEM:     types.String{Value: key.PrivateKeyPEM()},
		PublicKey:         types.String{Value: key.PublicKey()},
		FingerprintMD5:    types.String{Value: key.FingerprintMD5()},
		FingerprintSHA256: types.String{Value: key.FingerprintSHA256()},
	}

	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r resourceRSAKeyPair) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	// no need to support Read at the moment since the resource is fully within state
	// NOTE: if we support sourcing from files on disk in the future, this will have to be implemented
}

func (r resourceRSAKeyPair) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	// no need to support Update at the moment, since the only possible changes require replacement
	// NOTE: if we need to support changes later for comments, etc, this will need to be implemented
}

func (r resourceRSAKeyPair) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var state RSAKeyPair
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r resourceRSAKeyPair) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStateNotImplemented(ctx, "", resp)
}
