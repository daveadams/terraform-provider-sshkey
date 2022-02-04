package sshkey

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

var validRSABits = []int64{1024, 2048, 4096}

type rsaKeyPairBitsValidator struct{}

func (r rsaKeyPairBitsValidator) Description(ctx context.Context) string {
	return "RSA keypair bits must be 1024, 2048, or 4096."
}

func (r rsaKeyPairBitsValidator) MarkdownDescription(ctx context.Context) string {
	return "RSA keypair bits must be 1024, 2048, or 4096."
}

func (r rsaKeyPairBitsValidator) Validate(ctx context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) {
	var config RSAKeyPair
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	bits, _ := config.Bits.Value.Int64()
	for _, validBits := range validRSABits {
		if bits == validBits {
			// valid!
			return
		}
	}
	resp.Diagnostics.AddAttributeError(
		req.AttributePath,
		"Invalid RSA key bit size",
		"The value of 'bits' must be one of 1024, 2048, or 4096.",
	)
}
