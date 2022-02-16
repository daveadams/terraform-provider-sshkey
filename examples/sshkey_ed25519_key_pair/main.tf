terraform {
  required_providers {
    sshkey = {
      source = "daveadams/sshkey"
    }
  }
}

resource "sshkey_ed25519_key_pair" "example" {
  count = 2
}

output "example_public_keys" {
  value = sshkey_ed25519_key_pair.example[*].public_key
}

output "example_fingerprints_md5" {
  value = sshkey_ed25519_key_pair.example[*].fingerprint_md5
}

output "example_fingerprints_sha256" {
  value = sshkey_ed25519_key_pair.example[*].fingerprint_sha256
}
