package initconfig

import (
	services "github/lmf3z/jack-pass/internal/services/init-config"

	"github.com/spf13/cobra"
)

func InitCondifCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "init",
		Short:   "Start config",
		Aliases: []string{"i"},
		RunE:    services.RunInitCondifCommand,
	}

	// cmd.Flags().String(models.InitConfigNameFlag, "", "password you want to use!")

	// if err := cmd.MarkFlagRequired(models.InitConfigNameFlag); err != nil {
	// 	panic(err)
	// }

	return cmd
}
