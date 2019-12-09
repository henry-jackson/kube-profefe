package cmd

import (
	"github.com/spf13/cobra"
)

func NewGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	cmd.AddCommand(NewGetProfilesCmd())
	cmd.AddCommand(NewGetProfileTypesCmd())
	return cmd
}
