package tcp

import (
    "fmt"
    "net"
	"os"

	"github.com/spf13/cobra"
	"github.com/rs/zerolog"

	"github.com/coding-kelps/gomocku/cmd/gomocku/context"

	"github.com/coding-kelps/gomocku/pkg/adapters"
	"github.com/coding-kelps/gomocku/pkg/domain/ai"
	"github.com/coding-kelps/gomocku/pkg/domain/coordinator"
)

func InitActiveCmd() *cobra.Command {
    activeCmd := cobra.Command{
        Use:   "active",
        Short: "initiate connection to the gomokurs environment manager",
        Run:    activeExecute,
    }

	activeCmd.Flags().String("address", "localhost:49912", "the listening address")

	return &activeCmd
}

func activeExecute(cmd *cobra.Command, args []string) {
	ctx := cmd.Context()
	logger, ok := ctx.Value(context.LoggerKey).(zerolog.Logger)
	if !ok {
		fmt.Fprintf(os.Stderr, "FATAL - could not retrieve logger\n")
		
		return
	}
	address, _ := cmd.Flags().GetString("address")

    conn, err := net.Dial("tcp", address)
    if err != nil {
        logger.Fatal().
			Err(err).
			Msg("manager tcp dialing")
		
        return
    }
	defer conn.Close()

	tcp, err := adapters.NewTCPManagerInterface(conn, logger)
	if err != nil {
        logger.Fatal().
			Err(err).
			Msg("tcp manager interface creation failed")

		return
	}

	ai := ai.NewRandomAI(logger)
	coord := coordinator.NewCoordinator(tcp, ai, logger)
	
	err = coord.Serve()
	if err != nil {
        logger.Fatal().
			Err(err).
			Msg("coordinator failed")
		
		return
	}
}
