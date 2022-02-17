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

resource "sshkey_ed25519_key_pair" with_comment" {
  comment = "user@example.com"
}
```

## Arguments Reference

- `comment` - (Optional) Key comment to appear in the public key serialization.

## Attributes Reference

In addition to the arguments above, the following attributes are exported.

- `private_key_pem` - (Sensitive) The RSA private key in PEM format.
- `public_key` - The public key value in SSH2/`authorized_keys` format.
- `fingerprint_sha256` - SHA256 fingerprint of the key.
- `fingerprint_md5` - Legacy MD5 fingerprint of the key.
- `id` - The SHA256 fingerprint of the key.
