package getonepass

import (
	"fmt"
	services "github/lmf3z/jack-pass/internal/services/get-one-pass"

	"github.com/spf13/cobra"
)

func GetOnePassCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "get [keypass]",
		Short:   "Get password",
		Aliases: []string{"g"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) <= 0 {
				return fmt.Errorf("key to find is required!")
			}

			return services.RunGetOnePassCommand(args[0])
		},
	}

	return cmd
}
