package subcommands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/rs/zerolog"

	"github.com/coding-kelps/gomocku/cmd/gomocku/context"

	"github.com/coding-kelps/gomocku/pkg/adapters"
	"github.com/coding-kelps/gomocku/pkg/domain/ai"
	"github.com/coding-kelps/gomocku/pkg/domain/coordinator"
)

func InitStdioCmd() *cobra.Command {
	stdioCmd := cobra.Command{
		Use:	"stdio",
		Short:	"run the gomocku mock AI using standard I/O",
		Run:	stdioExecute,
	}

	return &stdioCmd
}

func stdioExecute(cmd *cobra.Command, _ []string) {
	ctx := cmd.Context()
	logger, ok := ctx.Value(context.LoggerKey).(zerolog.Logger)
	if !ok {
		fmt.Fprintf(os.Stderr, "FATAL - could not retrieve logger\n")

		return
	}

	stdio := adapters.NewStdioManagerInterface(logger)

	ai := ai.NewRandomAI(logger)
	coord := coordinator.NewCoordinator(stdio, ai, logger)

	err := coord.Serve()
	if err != nil {
        logger.Fatal().
			Err(err).
			Msg("coordinator failed")
		
		return
	}
}
