# terraform-provider-sshkey

Simple Terraform provider that generates SSH key pairs. Documentation is available
on the [sshkey provider page](https://registry.terraform.io/providers/daveadams/sshkey/latest/docs)
in the Terraform Registry.

## Example Usage

    terraform {
      required_providers {
        sshkey = {
          source = "daveadams/sshkey"
        }
      }
    }

    resource "sshkey_rsa_key_pair" "default" {
      bits = 4096
    }

    resource "sshkey_ed25519_key_pair" "admin" {
      comment = "admin@example.com"
    }

    output "default_key_fingerprint" {
      value = sshkey_rsa_key_pair.default.fingerprint_sha256
    }

    resource "local_file" "default_key" {
      filename        = "id.default"
      content         = sshkey_rsa_key_pair.default.private_key_pem
      file_permission = "0600"
    }

    resource "aws_key_pair" "default" {
      key_name   = "default"
      public_key = sshkey_rsa_key_pair.default.public_key
    }
