package sshkey

import (
	"strings"

	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"golang.org/x/crypto/ssh"
)

type rsaKeyPair struct {
	privateKey *rsa.PrivateKey
	publicKey  ssh.PublicKey
}

func generateRSAKeyPair(bits int) (*rsaKeyPair, error) {
	key, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}

	err = key.Validate()
	if err != nil {
		return nil, err
	}

	pub, err := ssh.NewPublicKey(key.Public())
	if err != nil {
		return nil, err
	}

	return &rsaKeyPair{
		privateKey: key,
		publicKey:  pub,
	}, nil
}

func (k *rsaKeyPair) PrivateKeyPEM() string {
	return string(pem.EncodeToMemory(
		&pem.Block{
			Type:  "OPENSSH PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(k.privateKey),
		}),
	)
}

func (k *rsaKeyPair) PublicKey() string {
	// we have to trim off the trailing newline that this function unhelpfully adds
	return strings.TrimSpace(string(ssh.MarshalAuthorizedKey(k.publicKey)))
}

func (k *rsaKeyPair) FingerprintMD5() string {
	return ssh.FingerprintLegacyMD5(k.publicKey)
}

func (k *rsaKeyPair) FingerprintSHA256() string {
	return ssh.FingerprintSHA256(k.publicKey)
}
