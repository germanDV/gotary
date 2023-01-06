package keyring

import (
	"encoding/hex"

	"github.com/tyler-smith/go-bip39"
)

func GenerateMnemonic() (string, error) {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return "", err
	}
	return bip39.NewMnemonic(entropy)
}

func MnemonicToSeed(mnemonic string) ([]byte, error) {
	password := ""
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, password)
	if err != nil {
		return []byte{}, err
	}
	// We want 32 bytes so that we can use it as a seed for ed25519 keys.
	return seed[0:32], nil
}

func SeedToHex(seed []byte) string {
	return hex.EncodeToString(seed)
}
