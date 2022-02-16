---
page_title: "ed25519_key_pair Resource - terraform-provider-sshkey"
subcategory: ""
description: |-
  The ed25519_key_pair resource generates an ED25519 key pair.
---

# Resource: sshkey_ed25519_key_pair

The `sshkey_ed25519_key_pair` resource generates an ED25519 key pair.

## Example Usage

```terraform
resource "sshkey_ed25519_key_pair" "example" {}
```

## Attributes Reference

The following attributes are exported.

- `private_key_pem` - (Sensitive) The RSA private key in PEM format.
- `public_key` - The public key value in SSH2/`authorized_keys` format.
- `fingerprint_sha256` - SHA256 fingerprint of the key.
- `fingerprint_md5` - Legacy MD5 fingerprint of the key.
- `id` - The SHA256 fingerprint of the key.
