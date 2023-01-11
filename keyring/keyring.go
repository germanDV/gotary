package keyring

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"errors"

	"github.com/tyler-smith/go-bip39"
)

type Public struct {
	Key ed25519.PublicKey `json:"key"`
	Hex string            `json:"hex"`
}

// Keyring holds a Private and a Public key-pair.
type Keyring struct {
	Private ed25519.PrivateKey
	Public  Public
}

// New generates a new key-pair and returns a Keyring.
func New() (*Keyring, error) {
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

// FromSeed creates a new Keyring from a seed.
func FromSeed(seed []byte) (*Keyring, error) {
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
func VerifySignature(publicKey ed25519.PublicKey, data []byte, signature []byte) bool {
	return ed25519.Verify(publicKey, data, signature)
}

// VerifyOwnSignature checks if the data has been signed by the current Keyring.
func (k *Keyring) VerifyOwnSignature(data []byte, signature []byte) bool {
	return ed25519.Verify(k.Public.Key, data, signature)
}

// GenerateMnemonic creates a new 12-word mnemonic.
func GenerateMnemonic() (string, error) {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return "", err
	}
	return bip39.NewMnemonic(entropy)
}

// MnemonicToSeed converts a 12-word mnemonic into a 32-byte seed.
func MnemonicToSeed(mnemonic string) ([]byte, error) {
	password := ""
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, password)
	if err != nil {
		return []byte{}, err
	}
	// We want 32 bytes so that we can use it as a seed for ed25519 keys.
	return seed[0:32], nil
}
