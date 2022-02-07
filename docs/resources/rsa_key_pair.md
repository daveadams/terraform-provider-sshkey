---
page_title: "rsa_key_pair Resource - terraform-provider-sshkey"
subcategory: ""
description: |-
  The rsa_key_pair resource generates an RSA key pair.
---

# Resource: sshkey_rsa_key_pair

The `sshkey_rsa_key_pair` resource generates an RSA key pair with the specified
number of bits.

## Example Usage

```terraform
resource "sshkey_rsa_key_pair" "small" {
  bits = 1024
}

resource "sshkey_rsa_key_pair" "medium" {
  bits = 2048
}

resource "sshkey_rsa_key_pair" "large" {
  bits = 4096
}
```

## Argument Reference

- `bits` - (Required) Bit size for RSA key. Must be one of 1024, 2048, or 4096.

## Attributes Reference

In addition to the arguments above, the following attributes are exported.

- `private_key_pem` - (Sensitive) The RSA private key in PEM format.
- `public_key` - The public key value in SSH2/`authorized_keys` format.
- `fingerprint_sha256` - SHA256 fingerprint of the key.
- `fingerprint_md5` - Legacy MD5 fingerprint of the key.
- `id` - The SHA256 fingerprint of the key.
