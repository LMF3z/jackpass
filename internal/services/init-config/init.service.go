package services

import (
	"fmt"
	"github/lmf3z/jack-pass/internal/models/database"
	"github/lmf3z/jack-pass/internal/utils"
	"github/lmf3z/jack-pass/pkg/encrypt"
	"os"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

func RunInitCondifCommand(cmd *cobra.Command, args []string) error {
	exitDbVault := utils.ValidateIfFileExist(database.DbVaultFileName)

	if exitDbVault {
		return fmt.Errorf("\n Config file already exist")
	}

	fmt.Printf("Enter your master password: \n")

	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println("\nError to read password")
		return err
	}

	password := string(bytePassword)

	salt, _ := encrypt.GenerateSalt()
	masterKey, _ := encrypt.GenerateKey(password, salt)

	fmt.Printf("Your master key of %d bits was generated.\n", len(masterKey)*8)

	v := database.NewDbVault()

	plaintText, _ := v.ToBytes()

	encryptedData, err := encrypt.EncryptFile(plaintText, password)
	if err != nil {
		fmt.Printf("Error al cifrar: %v\n", err)
		return err
	}

	// save on disk
	err = os.WriteFile(database.DbVaultFileName, encryptedData, 0600)
	if err != nil {
		fmt.Printf("Error to save file: %v\n", err)
		return err
	}

	fmt.Printf("The config file %s was successfully created!\n", database.DbVaultFileName)

	return nil
}
