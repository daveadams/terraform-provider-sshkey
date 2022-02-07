terraform {
  required_providers {
    sshkey = {
      version = "0.1.0"
      source  = "daveadams/sshkey"
    }
  }
}

resource "sshkey_rsa_key_pair" "example" {
  count = 2
  bits  = 2048
}

resource "sshkey_rsa_key_pair" "big" {
  bits = 4096
}

resource "sshkey_rsa_key_pair" "small" {
  bits = 1024
}

output "example_fingerprint_md5" {
  value = sshkey_rsa_key_pair.example[*].fingerprint_md5
}

output "big_fingerprint_md5" {
  value = sshkey_rsa_key_pair.big.fingerprint_md5
}

output "small_fingerprint_sha256" {
  value = sshkey_rsa_key_pair.small.fingerprint_sha256
}
