package cmd

import (
	addnewpass "github/lmf3z/jack-pass/internal/cli/add-new-pass"
	deletepass "github/lmf3z/jack-pass/internal/cli/delete-pass"
	getonepass "github/lmf3z/jack-pass/internal/cli/get-one-pass"
	initproject "github/lmf3z/jack-pass/internal/cli/init-config"
	seedata "github/lmf3z/jack-pass/internal/cli/see-data"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "jackpass",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(initproject.InitCondifCommand())
	rootCmd.AddCommand(seedata.SeeDataCommand())
	rootCmd.AddCommand(getonepass.GetOnePassCommand())
	rootCmd.AddCommand(addnewpass.AddNewPassCommand())
	rootCmd.AddCommand(deletepass.DeletePassCommand())
}
