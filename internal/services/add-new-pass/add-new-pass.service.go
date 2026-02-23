package addnewpass

import (
	"fmt"
	"github/lmf3z/jack-pass/internal/models/database"
	"github/lmf3z/jack-pass/internal/utils"
	"syscall"

	"golang.org/x/term"
)

func RunAddNewPassCommand(keyEntry string) error {

	exitDbVault := utils.ValidateIfFileExist(database.DbVaultFileName)
	if !exitDbVault {
		return fmt.Errorf("\n Config file not already exist")
	}

	vaultData, err := utils.GetEncrypteData()
	if err != nil {
		return err
	}

	fmt.Printf("Enter your key (%s) password: \n", keyEntry)
	bytePassword, _ := term.ReadPassword(int(syscall.Stdin))
	keyPassword := string(bytePassword)

	vaultData.Vaul.Entries[keyEntry] = keyPassword

	vaultData.Vaul.SaveVault(vaultData.Password)

	fmt.Printf("Password added successfully!\n")

	return nil
}
