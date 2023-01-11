package keyring

import (
	"testing"
)

func noErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("expected no error, got: %s", err)
	}
}

func generateKeyring(t *testing.T, mnemonic string) ([]byte, *Keyring) {
	t.Helper()
	seed, err := MnemonicToSeed(mnemonic)
	noErr(t, err)
	keys, err := FromSeed(seed)
	noErr(t, err)
	return seed, keys
}

func TestKeyringFromSeed(t *testing.T) {
	seed, k1 := generateKeyring(t, "seven almost tail nose vague drink wonder worry school broken bright slush")

	t.Run("GeneratesSameKeysFromSameSeed", func(t *testing.T) {
		t.Parallel()
		k2, err := FromSeed(seed)
		noErr(t, err)
		if !k1.Private.Equal(k2.Private) || !k1.Public.Key.Equal(k2.Public.Key) {
			t.Error("Keys generated from the same seed are not equal")
		}
	})

	t.Run("GeneratesDifferentKeysFromDifferentSeed", func(t *testing.T) {
		t.Parallel()
		_, k2 := generateKeyring(t, "paddle battle invest cactus curious ball cargo rice gift rack burger permit")
		if k1.Private.Equal(k2.Private) || k1.Public.Key.Equal(k2.Public.Key) {
			t.Error("Keys generated from different seeds are equal")
		}
	})
}

func TestSignatureAndVerification(t *testing.T) {
	_, alice := generateKeyring(t, "seven almost tail nose vague drink wonder worry school broken bright slush")
	_, bob := generateKeyring(t, "paddle battle invest cactus curious ball cargo rice gift rack burger permit")

	t.Run("GoodSignature", func(t *testing.T) {
		t.Parallel()
		message := []byte("lorem ipsum dolor")
		aliceSignature := alice.Sign(message)
		if !VerifySignature(alice.Public.Key, message, aliceSignature) {
			t.Error("signature verification failed")
		}
	})

	t.Run("AlteredMessage", func(t *testing.T) {
		t.Parallel()
		originalMessage := []byte("original message.")
		aliceSignature := alice.Sign(originalMessage)
		alteredMessage := []byte("this is not the original message.")
		if VerifySignature(alice.Public.Key, alteredMessage, aliceSignature) {
			t.Error("signature verification succeeded for an altered message")
		}
	})

	t.Run("SignedByAnotherPublicKey", func(t *testing.T) {
		t.Parallel()
		message := []byte("lorem ipsum dolor")
		signature := bob.Sign(message)
		if VerifySignature(alice.Public.Key, message, signature) {
			t.Error("signature verification succeeded for message signed with another key")
		}
	})

	t.Run("GoodSignatureGettingPublicKeyFromHex", func(t *testing.T) {
		t.Parallel()
		message := []byte("lorem ipsum dolor")
		myContactSignature := alice.Sign(message)
		myContactPublicKey, err := PublicKeyFromHex(alice.Public.Hex)
		noErr(t, err)
		if !VerifySignature(myContactPublicKey.Key, message, myContactSignature) {
			t.Error("signature verification failed")
		}
	})
}

func TestMnemonicToSeed(t *testing.T) {
	goodMnemonic := "seven almost tail nose vague drink wonder worry school broken bright slush"
	shortMnemonic := "almost tail nose vague drink wonder worry school broken bright slush"
	longMnemonic := "seven almost tail nose vague drink wonder worry school broken bright slush cargo"
	nonDictionaryMnemonic := "seven almost tail nose vague drink wonder worry school broken bright slash"

	t.Run("ReturnsCorrectLengthSeed", func(t *testing.T) {
		t.Parallel()
		seed, err := MnemonicToSeed(goodMnemonic)
		noErr(t, err)
		if len(seed) != 32 {
			t.Errorf("expected seed length of 32 bytes, got: %d", len(seed))
		}
	})

	t.Run("FailsWithShortMnemonic", func(t *testing.T) {
		t.Parallel()
		_, err := MnemonicToSeed(shortMnemonic)
		if err == nil {
			t.Error("expected error when mnemonic is missing words")
		}
	})

	t.Run("FailsWithLongMnemonic", func(t *testing.T) {
		t.Parallel()
		_, err := MnemonicToSeed(longMnemonic)
		if err == nil {
			t.Error("expected error when mnemonic has additional words")
		}
	})

	t.Run("FailsWithMnemonicWithNonDictionaryWord", func(t *testing.T) {
		t.Parallel()
		_, err := MnemonicToSeed(nonDictionaryMnemonic)
		if err == nil {
			t.Error("expected error when mnemonic has non-dictionary words")
		}
	})
}
