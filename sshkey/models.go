package sshkey

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RSAKeyPair struct {
	ID                types.String `tfsdk:"id"`
	Bits              types.Number `tfsdk:"bits"`
	Comment           types.String `tfsdk:"comment"`
	PrivateKeyPEM     types.String `tfsdk:"private_key_pem"`
	PublicKey         types.String `tfsdk:"public_key"`
	FingerprintMD5    types.String `tfsdk:"fingerprint_md5"`
	FingerprintSHA256 types.String `tfsdk:"fingerprint_sha256"`
}

type ED25519KeyPair struct {
	ID                types.String `tfsdk:"id"`
	Comment           types.String `tfsdk:"comment"`
	PrivateKeyPEM     types.String `tfsdk:"private_key_pem"`
	PublicKey         types.String `tfsdk:"public_key"`
	FingerprintMD5    types.String `tfsdk:"fingerprint_md5"`
	FingerprintSHA256 types.String `tfsdk:"fingerprint_sha256"`
}
