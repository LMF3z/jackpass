package utils

import (
	"encoding/json"
	"fmt"
	"github/lmf3z/jack-pass/internal/models/database"
	"github/lmf3z/jack-pass/pkg/encrypt"
	"os"
	"syscall"

	"golang.org/x/term"
)

type EncryptedData struct {
	Vaul     *database.DBVault
	Password string
}

func GetEncrypteData() (*EncryptedData, error) {
	encryptedData, _ := os.ReadFile(database.DbVaultFileName)

	fmt.Print("Enter your master password: \n")
	bytePassword, _ := term.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)
	fmt.Println("\n[Verifiying...]")

	plaintText, err := encrypt.DecryptFile(encryptedData, password)
	if err != nil {
		return nil, fmt.Errorf("Error: %v\n", err)

	}

	var vaulDb database.DBVault
	err = json.Unmarshal(plaintText, &vaulDb)
	if err != nil {
		return nil, fmt.Errorf("Error: %v\n", err)
	}

	return &EncryptedData{
		Vaul:     &vaulDb,
		Password: password,
	}, nil
}
