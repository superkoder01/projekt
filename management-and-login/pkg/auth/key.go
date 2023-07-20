package auth

import (
	"crypto/ed25519"
	"encoding/asn1"
	"encoding/pem"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/file_utils"
)

type KeyLoader interface {
	GetPrivateKey() (priv ed25519.PrivateKey, err error)
	GetPublicKey() (pub ed25519.PublicKey, err error)
}

type keyLoader struct {
	KeyPath string
}

type ed25519PrivKey struct {
	Version          int
	ObjectIdentifier struct {
		ObjectIdentifier asn1.ObjectIdentifier
	}
	PrivateKey []byte
}

func newKeyLoader(keyPath string) *keyLoader {
	return &keyLoader{
		KeyPath: keyPath,
	}
}

func (k *keyLoader) GetPrivateKey() (ed25519.PrivateKey, error) {
	key, err := loadKeyFromPEMFile(k.KeyPath)
	if err != nil {
		return nil, err
	}

	var block *pem.Block
	block, _ = pem.Decode(key)

	var asn1PrivKey ed25519PrivKey
	_, err = asn1.Unmarshal(block.Bytes, &asn1PrivKey)
	if err != nil {
		return nil, err
	}

	return ed25519.NewKeyFromSeed(asn1PrivKey.PrivateKey[2:]), nil
}

func (k *keyLoader) GetPublicKey() (ed25519.PublicKey, error) {
	priv, err := k.GetPrivateKey()
	if err != nil {
		return nil, err
	}

	pub, ok := priv.Public().(ed25519.PublicKey)
	if !ok {
		return nil, ErrNotEd25519PublicKey
	}

	return pub, nil
}

func loadKeyFromPEMFile(keyPath string) ([]byte, error) {
	return utils.ReadFileContent(keyPath)
}
