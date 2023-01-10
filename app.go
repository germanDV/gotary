package main

import (
	"context"
	"encoding/hex"
	"errors"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gitlab.com/germandv/gotary/keyring"
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
	mnemonic, ok := keyring.GetSavedMnemonic()
	if ok {
		a.Login(mnemonic)
	}
}

func (a *App) Login(mnemonic string) error {
	seed, err := keyring.MnemonicToSeed(mnemonic)
	if err != nil {
		return err
	}

	keys, err := keyring.KeyringFromSeed(seed)
	if err != nil {
		return err
	}

	a.keys = keys
	return keyring.WriteMnemonic(mnemonic)
}

func (a *App) IsLoggedIn() bool {
	return a.keys != nil
}

func (a *App) Logout() error {
	err := keyring.DeleteMnemonic()
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

func (a *App) CopyToClipboard(text string) error {
	err := clipboard.Init()
	if err != nil {
		return err
	}
	clipboard.Write(clipboard.FmtText, []byte(text))
	return nil
}

func (a *App) GetContacts() ([]*keyring.Contact, error) {
	return keyring.GetContacts()
}

func (a *App) AddContact(name string, publicKeyHex string) error {
	return keyring.NewContact(name, publicKeyHex)
}

func (a *App) DeleteContact(name string) error {
	return keyring.DeleteContact(name)
}
