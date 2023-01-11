package main

import (
	"context"
	"encoding/hex"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gitlab.com/germandv/gotary/contacts"
	"gitlab.com/germandv/gotary/keyring"
	"gitlab.com/germandv/gotary/repository"
	"golang.design/x/clipboard"
)

type App struct {
	ctx  context.Context
	keys *keyring.Keyring
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	mnemonic, ok := repository.GetSavedMnemonic()
	if ok {
		a.Login(mnemonic)
	}
}

func (a *App) Login(mnemonic string) error {
	seed, err := keyring.MnemonicToSeed(mnemonic)
	if err != nil {
		return err
	}

	keys, err := keyring.FromSeed(seed)
	if err != nil {
		return err
	}

	a.keys = keys
	return repository.WriteMnemonic(mnemonic)
}

func (a *App) IsLoggedIn() bool {
	return a.keys != nil
}

func (a *App) Logout() error {
	err := repository.DeleteMnemonic()
	if err != nil {
		return err
	}
	a.keys = nil
	return nil
}

func (a *App) GetMyPublicKey() (string, error) {
	if a.keys == nil {
		return "", errors.New("no keyring initialized")
	}
	return a.keys.Public.Hex, nil
}

func (a *App) SelectFile() (string, error) {
	path, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select file to sign",
	})
	if err != nil {
		return "", err
	}
	return path, nil
}

func (a *App) SignFile(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	if a.keys == nil {
		return "", errors.New("no keyring initialized")
	}

	signature := a.keys.Sign(content)
	return hex.EncodeToString(signature), nil
}

func (a *App) VerifySignature(filePath, signaturePath, candidateKey string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	signatureHex, err := os.ReadFile(signaturePath)
	if err != nil {
		return err
	}

	signatureStr := strings.TrimSuffix(strings.TrimSuffix(string(signatureHex), "\n"), "\r")
	signature, err := hex.DecodeString(signatureStr)
	if err != nil {
		return err
	}

	key, err := keyring.PublicKeyFromHex(candidateKey)
	if err != nil {
		return err
	}

	if !keyring.VerifySignature(key.Key, data, signature) {
		return errors.New("Bad signature")
	}
	return nil
}

func (a *App) CopyToClipboard(text string) error {
	err := clipboard.Init()
	if err != nil {
		return err
	}
	clipboard.Write(clipboard.FmtText, []byte(text))
	return nil
}

func (a *App) SaveSignatureToDisk(filePath, signature string) (string, error) {
	file := filepath.Base(filePath)
	base := strings.TrimSuffix(filePath, file)
	ext := filepath.Ext(filePath)
	sig := strings.TrimSuffix(file, ext) + ".sig"
	signaturePath := base + sig
	err := repository.WriteSignature(signaturePath, signature)
	if err != nil {
		return "", err
	}
	return sig, nil
}

func (a *App) GetContacts() ([]*contacts.Contact, error) {
	return contacts.GetAll()
}

func (a *App) AddContact(name string, publicKeyHex string) error {
	return contacts.New(name, publicKeyHex)
}

func (a *App) DeleteContact(name string) error {
	return contacts.Delete(name)
}

func (a *App) GenerateMnemonic() (string, error) {
	return keyring.GenerateMnemonic()
}
