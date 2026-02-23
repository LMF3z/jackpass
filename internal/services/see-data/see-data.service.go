package seedata

import (
	"fmt"
	"github/lmf3z/jack-pass/internal/models/database"
	"github/lmf3z/jack-pass/internal/utils"

	"github.com/spf13/cobra"
)

func RunSeeDataCommand(cmd *cobra.Command, args []string) error {
	exitDbVault := utils.ValidateIfFileExist(database.DbVaultFileName)

	if !exitDbVault {
		return fmt.Errorf("\n Config file not already exist")
	}

	vaultData, err := utils.GetEncrypteData()
	if err != nil {
		return err
	}

	if len(vaultData.Vaul.Entries) <= 0 {
		fmt.Printf("There is not data to show!\n")
		return nil
	}

	for site := range vaultData.Vaul.Entries {
		fmt.Printf("Site: %-10s | Key: %s\n", site, "******")
	}

	return nil
}
