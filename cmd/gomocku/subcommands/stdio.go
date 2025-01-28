package subcommands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/coding-kelps/gomocku/pkg/adapters"
	"github.com/coding-kelps/gomocku/pkg/domain/ai"
	"github.com/coding-kelps/gomocku/pkg/domain/listener"
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
	stdio := adapters.NewStdio()

	ai := ai.NewRandomAI()
	listener := listener.NewListener(stdio, ai)

	err := listener.Listen()
	if err != nil {
		fmt.Printf("%e\n", err)
	}
}
