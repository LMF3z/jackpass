package addnewpass

import (
	"fmt"
	addnewpass "github/lmf3z/jack-pass/internal/services/add-new-pass"

	"github.com/spf13/cobra"
)

func AddNewPassCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [key]",
		Short: "Add password",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) <= 0 {
				return fmt.Errorf("key to save is required!")
			}

			return addnewpass.RunAddNewPassCommand(args[0])
		},
	}

	return cmd
}
