package keyring

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"errors"
)

type Public struct {
	Key ed25519.PublicKey `json:"key"`
	Hex string            `json:"hex"`
}

// Keyring holds a Private and a Public key.
type Keyring struct {
	Private ed25519.PrivateKey
	Public  Public
}

// NewKeyring generates a new key-pair and returns a Keyring.
func NewKeyring() (*Keyring, error) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	return &Keyring{
		Private: privateKey,
		Public: Public{
			Key: publicKey,
			Hex: publicKeyToHex(publicKey),
		},
	}, nil
}

// KeyringFromSeed creates a new Keyring from a seed.
func KeyringFromSeed(seed []byte) (*Keyring, error) {
	if len(seed) != ed25519.SeedSize {
		return nil, errors.New("invalid seed size for keyring generation")
	}

	privateKey := ed25519.NewKeyFromSeed(seed)
	cryptoPublicKey := privateKey.Public()
	publicKey, ok := cryptoPublicKey.(ed25519.PublicKey)
	if !ok {
		return nil, errors.New("cannot cast public key to ed25519.PublicKey")
	}

	return &Keyring{
		Private: privateKey,
		Public: Public{
			Key: publicKey,
			Hex: publicKeyToHex(publicKey),
		},
	}, nil
}

// publicKeyToHex returns the public key encoded as a hex string.
func publicKeyToHex(publicKey []byte) string {
	return hex.EncodeToString(publicKey)
}

// PublicKeyFromHex parses a hex string into a Public Key.
func PublicKeyFromHex(hs string) (*Public, error) {
	bytes, err := hex.DecodeString(hs)
	if err != nil {
		return nil, err
	}

	if len(bytes) != ed25519.PublicKeySize {
		return nil, errors.New("invalid public key size.")
	}

	return &Public{
		Key: bytes,
		Hex: hs,
	}, nil
}

// Sign signs the data with the Keyring's private key and returns the signature.
func (k *Keyring) Sign(data []byte) []byte {
	return ed25519.Sign(k.Private, data)
}

// VerifySignature checks if the data has been signed by the given public key.
func (k *Keyring) VerifySignature(publicKey ed25519.PublicKey, data []byte, signature []byte) bool {
	return ed25519.Verify(publicKey, data, signature)
}

// VerifyOwnSignature checks if the data has been signed by the current Keyring.
func (k *Keyring) VerifyOwnSignature(data []byte, signature []byte) bool {
	return ed25519.Verify(k.Public.Key, data, signature)
}
