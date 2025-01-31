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

func InitPassiveCmd() *cobra.Command {
	passiveCmd := cobra.Command{
		Use:   "passive",
		Short: "wait for the gomokurs environment manager to initiate connection",
		Run: 	passiveExecute,
	}

	passiveCmd.Flags().String("address", "localhost:49912", "the listening address")

	return &passiveCmd
}

func passiveExecute(cmd *cobra.Command, args []string) {
	ctx := cmd.Context()
	logger, ok := ctx.Value(context.LoggerKey).(zerolog.Logger)
	if !ok {
		fmt.Fprintf(os.Stderr, "FATAL - could not retrieve logger\n")

		return
	}
	address, _ := cmd.Flags().GetString("address")

	listener, err := net.Listen("tcp", address)
	if err != nil {
        logger.Fatal().
			Err(err).
			Msg("manager tcp dialing")

		return
	}

    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
		defer conn.Close()

		tcp, err := adapters.NewTCPManagerInterface(conn, logger.With().Str("adapter", "tcp").Logger())
		if err != nil {
			logger.Fatal().
				Err(err).
				Msg("tcp manager interface creation failed")
			
			return
		}

		ai := ai.NewRandomAI(logger.With().Str("service", "ai").Logger())
		coord := coordinator.NewCoordinator(tcp, ai, logger.With().Str("service", "coordinator").Logger())
	
		err = coord.Serve()
		if err != nil {
			logger.Fatal().
				Err(err).
				Msg("coordinator failed")

			return
		}

		return
	}
}
