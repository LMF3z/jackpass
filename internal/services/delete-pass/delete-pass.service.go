package deletepass

import (
	"fmt"
	"github/lmf3z/jack-pass/internal/models/database"
	"github/lmf3z/jack-pass/internal/utils"
)

func RunDeletePassCommnad(keyEntry string) error {
	exitDbVault := utils.ValidateIfFileExist(database.DbVaultFileName)
	if !exitDbVault {
		return fmt.Errorf("\n Config file not already exist")
	}

	vaultData, err := utils.GetEncrypteData()
	if err != nil {
		return err
	}

	delete(vaultData.Vaul.Entries, keyEntry)

	vaultData.Vaul.SaveVault(vaultData.Password)

	fmt.Printf("Password deleted successfully!\n")

	return nil
}
