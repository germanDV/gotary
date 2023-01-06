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
