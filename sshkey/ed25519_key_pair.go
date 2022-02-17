package sshkey

import (
	"strings"

	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"golang.org/x/crypto/ssh"
)

type ed25519KeyPair struct {
	privateKey ed25519.PrivateKey
	publicKey  ssh.PublicKey
}

func generateED25519KeyPair() (*ed25519KeyPair, error) {
	rawPubKey, key, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	pub, err := ssh.NewPublicKey(rawPubKey)
	if err != nil {
		return nil, err
	}

	return &ed25519KeyPair{
		privateKey: key,
		publicKey:  pub,
	}, nil
}

func (k *ed25519KeyPair) PrivateKeyPEM() (string, error) {
	keyBytes, err := x509.MarshalPKCS8PrivateKey(k.privateKey)
	if err != nil {
		return "", err
	}

	return string(pem.EncodeToMemory(
		&pem.Block{
			Type:    "OPENSSH PRIVATE KEY",
			Headers: nil,
			Bytes:   keyBytes,
		}),
	), nil
}

func (k *ed25519KeyPair) PublicKey() string {
	// we have to trim off the trailing newline that this function unhelpfully adds
	return strings.TrimSpace(string(ssh.MarshalAuthorizedKey(k.publicKey)))
}

func (k *ed25519KeyPair) FingerprintMD5() string {
	return ssh.FingerprintLegacyMD5(k.publicKey)
}

func (k *ed25519KeyPair) FingerprintSHA256() string {
	return ssh.FingerprintSHA256(k.publicKey)
}
