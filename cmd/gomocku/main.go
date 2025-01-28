package main

import (
	"fmt"

	"github.com/coding-kelps/gomocku/cmd/gomocku/subcommands"

	"github.com/spf13/cobra"
)



func initRootCmd() *cobra.Command {
	rootCmd := cobra.Command{
		Use: "gomocku",
		Short: "gomocku - a testing client for the gomokurs environment",
		Version: "0.1.0",
	}

	rootCmd.AddCommand(subcommands.InitStdioCmd())
	rootCmd.AddCommand(subcommands.InitTcpCmd())

	return &rootCmd
}

func main() {
	rootCmd := initRootCmd()

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("%v\n", err)
	}
}


