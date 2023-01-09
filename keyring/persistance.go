package keyring

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
)

const (
	GotaryDir          = ".gotary"
	GotaryMnemonicFile = "mnemonic.txt"
	GotaryContactsFile = "contacts.json"
)

func mkdirIfNotExist(dir string) error {
	_, err := os.Stat(dir)
	if !errors.Is(err, os.ErrNotExist) {
		return err
	}
	return os.Mkdir(dir, fs.ModePerm)
}

func WriteMnemonic(mnemonic string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	err = mkdirIfNotExist(filepath.Join(home, GotaryDir))
	if err != nil {
		return err
	}

	// TODO: encrypt the file that contains the mnemonic.
	return os.WriteFile(
		filepath.Join(home, GotaryDir, GotaryMnemonicFile),
		[]byte(mnemonic),
		os.FileMode(0600),
	)
}

func GetSavedMnemonic() (string, bool) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", false
	}

	mnemonic, err := os.ReadFile(filepath.Join(home, GotaryDir, GotaryMnemonicFile))
	if err != nil {
		return "", false
	}

	return string(mnemonic), true
}

func DeleteMnemonic() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	return os.Remove(filepath.Join(home, GotaryDir, GotaryMnemonicFile))
}

func ContactsFileExists() bool {
	home, err := os.UserHomeDir()
	if err != nil {
		return false
	}
	file := filepath.Join(home, GotaryDir, GotaryContactsFile)
	_, err = os.Stat(file)
	if err != nil {
		return false
	}
	return true
}

func WriteContacts(contactJsonBytes []byte) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	file := filepath.Join(home, GotaryDir, GotaryContactsFile)
	return os.WriteFile(file, contactJsonBytes, os.FileMode(0600))
}

func ReadContactFile() ([]byte, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	file := filepath.Join(home, GotaryDir, GotaryContactsFile)

	_, err = os.Stat(file)
	if errors.Is(err, os.ErrNotExist) {
		return []byte{}, nil
	}

	return os.ReadFile(file)
}
