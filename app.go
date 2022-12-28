package main

import (
	"context"
	"fmt"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) DecryptPDF(pdfPath string) string {
	fmt.Println("Decrypting", pdfPath)
	conf := pdfcpu.NewAESConfiguration("upw", "opw", 256)
	err := api.DecryptFile(pdfPath, pdfPath, conf)
	if err != nil {
		return fmt.Sprintf("Failed to decrypt %s: %v", pdfPath, err)
	}

	return ""
}

func (a *App) EncryptPDF(pdfPath string) string {
	fmt.Println("Encrypting", pdfPath)
	conf := pdfcpu.NewAESConfiguration("upw", "opw", 256)
	err := api.EncryptFile(pdfPath, pdfPath, conf)
	if err != nil {
		return fmt.Sprintf("Failed to encrypt %s: %v", pdfPath, err)
	}

	return ""
}
