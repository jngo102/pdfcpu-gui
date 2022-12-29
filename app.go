package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

const (
	// Name of config folder
	configFolder = "pdfcpu-gui"
	// Name of log file
	logFile = "pdfcpu-gui.log"
)

// App struct
type App struct {
	ctx      context.Context
	settings Settings
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	configDir, err := os.UserConfigDir()
	if err != nil {
		panic("Failed to get user config dir")
	}
	configPath := filepath.Join(configDir, configFolder)
	if os.MkdirAll(configPath, os.ModePerm) != nil {
		panic(fmt.Sprint("Failed to create config dir: ", err))
	}

	logPath := filepath.Join(configPath, logFile)
	logFile, err := os.OpenFile(logPath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(fmt.Sprint("Failed to open/create log file: ", err))
	}
	log.SetOutput(logFile)

	settings, err := LoadSettings()
	if err != nil {
		log.Println("Failed to load settings file, creating a new one")
		a.settings = NewSettings()
	} else {
		a.settings = settings
	}
}

func (a *App) shutdown(ctx context.Context) {
	log.Println("Saving settings on quit")
	a.settings.Save()
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
