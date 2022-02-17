terraform {
  required_providers {
    sshkey = {
      source = "daveadams/sshkey"
    }
  }
}

resource "sshkey_ed25519_key_pair" "no_comment" {}

resource "sshkey_ed25519_key_pair" "with_comment" {
  comment = "admin@example.com"
}

output "no_comment" {
  value = {
    public_key         = sshkey_ed25519_key_pair.no_comment.public_key
    md5_fingerprint    = sshkey_ed25519_key_pair.no_comment.fingerprint_md5
    sha256_fingerprint = sshkey_ed25519_key_pair.no_comment.fingerprint_sha256
  }
}

output "with_comment" {
  value = {
    public_key         = sshkey_ed25519_key_pair.with_comment.public_key
    md5_fingerprint    = sshkey_ed25519_key_pair.with_comment.fingerprint_md5
    sha256_fingerprint = sshkey_ed25519_key_pair.with_comment.fingerprint_sha256
  }
}
