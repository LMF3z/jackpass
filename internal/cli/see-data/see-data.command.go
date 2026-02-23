package seedata

import (
	seedata "github/lmf3z/jack-pass/internal/services/see-data"

	"github.com/spf13/cobra"
)

func SeeDataCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List enables pass",
		Aliases: []string{"ps"},
		RunE:    seedata.RunSeeDataCommand,
	}

	return cmd
}
