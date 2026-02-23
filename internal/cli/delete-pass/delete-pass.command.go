package deletepass

import (
	"fmt"
	deletepass "github/lmf3z/jack-pass/internal/services/delete-pass"

	"github.com/spf13/cobra"
)

func DeletePassCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "del [key]",
		Short:   "Delete password",
		Aliases: []string{"d"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) <= 0 {
				return fmt.Errorf("key to delete is required!")
			}

			return deletepass.RunDeletePassCommnad(args[0])
		},
	}

	return cmd
}
