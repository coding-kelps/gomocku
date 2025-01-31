package main

import (
	"fmt"
	"os"
	"context"
	"time"

	"github.com/spf13/cobra"
	"github.com/rs/zerolog"
	
	"github.com/coding-kelps/gomocku/cmd/gomocku/subcommands"
	gomockuContext "github.com/coding-kelps/gomocku/cmd/gomocku/context"
)

func initRootCmd() *cobra.Command {
	rootCmd := cobra.Command{
		Use: "gomocku",
		Short: "gomocku - a testing client for the gomokurs environment",
		Version: "0.1.0",
		CompletionOptions: cobra.CompletionOptions{
			HiddenDefaultCmd: true,
		},
		PersistentPreRunE: rootCmdPersistentPreRunE,
	}

	rootCmd.PersistentFlags().String("log-level", "INFO", "the logging level of the application")

	rootCmd.AddCommand(subcommands.InitStdioCmd())
	rootCmd.AddCommand(subcommands.InitTcpCmd())

	return &rootCmd
}

func rootCmdPersistentPreRunE(cmd *cobra.Command, args []string) error {
	logLevel, err := cmd.Flags().GetString("log-level")
	if err != nil {
		return fmt.Errorf("zerolog logger initialization failed: %v", err)
	}

	lvl, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		return fmt.Errorf("invalid log level '%s', defaulting to info: %v", logLevel, err)
	}

	consoleWriter := zerolog.ConsoleWriter{
		Out:		os.Stderr,
		TimeFormat: time.RFC3339,
	}

	logger := zerolog.New(consoleWriter).
		Level(lvl).
		With().
		Timestamp().
		Logger()

	ctx := context.WithValue(cmd.Context(), gomockuContext.LoggerKey, logger)
	cmd.SetContext(ctx)

	return nil
}

func main() {
	rootCmd := initRootCmd()

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "FATAL - %v\n", err)
	}
}
