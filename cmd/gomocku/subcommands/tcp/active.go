package tcp

import (
	"github.com/spf13/cobra"
)

func InitActiveCmd() *cobra.Command {
    activeCmd := cobra.Command{
        Use:   "active",
        Short: "initiate connection to the gomokurs environment manager",
        Run:    activeExecute,
    }

	activeCmd.Flags().String("address", "localhost:49912", "The listening address")

	return &activeCmd
}

func activeExecute(cmd *cobra.Command, args []string) {
	_, _ = cmd.Flags().GetString("address")
}
