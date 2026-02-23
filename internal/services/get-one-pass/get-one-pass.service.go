package services

import (
	"fmt"
	"github/lmf3z/jack-pass/internal/models/database"
	"github/lmf3z/jack-pass/internal/utils"
	"time"

	"golang.design/x/clipboard"
)

func RunGetOnePassCommand(keypass string) error {

	exitDbVault := utils.ValidateIfFileExist(database.DbVaultFileName)

	if !exitDbVault {
		return fmt.Errorf("\n Config file not already exist")
	}

	vaultData, err := utils.GetEncrypteData()
	if err != nil {
		return err
	}

	var passFound string

	for site, key := range vaultData.Vaul.Entries {
		if keypass == site {
			passFound = key
		}
	}

	if len(passFound) <= 0 {
		return fmt.Errorf("Not found password! to: %s\n", keypass)
	}

	err = clipboard.Init()
	if err != nil {
		panic(err)
	}

	clipboard.Write(clipboard.FmtText, []byte(passFound))
	fmt.Printf("Password paste to clipboard! it will clean in 10s\n")

	time.Sleep(10 * time.Second)
	clipboard.Write(clipboard.FmtText, nil)
	fmt.Printf("Clipboard cleaned\n")
	time.Sleep(1 * time.Second)

	return nil
}
