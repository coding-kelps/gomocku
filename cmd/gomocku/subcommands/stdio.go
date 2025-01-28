package subcommands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

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

func stdioExecute(_ *cobra.Command, _ []string) {
	stdio := adapters.NewStdioManagerInterface()

	ai := ai.NewRandomAI()
	coord := coordinator.NewCoordinator(stdio, ai)

	err := coord.Serve()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%e\n", err)
	}
}
