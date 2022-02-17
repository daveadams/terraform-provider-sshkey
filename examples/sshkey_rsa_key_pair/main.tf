terraform {
  required_providers {
    sshkey = {
      source = "daveadams/sshkey"
    }
  }
}

resource "sshkey_rsa_key_pair" "example" {
  count = 2
  bits  = 2048
}

resource "sshkey_rsa_key_pair" "small" {
  bits = 1024
}

resource "sshkey_rsa_key_pair" "big" {
  bits    = 4096
  comment = "admin@example.com"
}

output "example_fingerprint_md5" {
  value = sshkey_rsa_key_pair.example[*].fingerprint_md5
}

output "big_public_key" {
  value = sshkey_rsa_key_pair.big.public_key
}

output "small_fingerprint_sha256" {
  value = sshkey_rsa_key_pair.small.fingerprint_sha256
}

output "small_public_key" {
  value = sshkey_rsa_key_pair.small.public_key
}
