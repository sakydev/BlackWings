package account

import (
	"BlackWings/internal/services"
	"BlackWings/internal/types"
	"BlackWings/internal/utils"
	"fmt"

	"github.com/spf13/cobra"
)

const mainCommand = "accounts"
const connectCommand = "connect"

func NewAppCommand(cmdConfig types.CommandConfiguration) *cobra.Command {
	var flags types.AppFlags

	appCommands := &cobra.Command{
		Use:   mainCommand,
		Short: "Add, remove, or list apps",
		Long:  `Add, remove, or list apps.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Use '%s --help' for more information.\n", mainCommand)
		},
	}

	connectCommand := &cobra.Command{
		Use:   connectCommand,
		Short: "Connect a new account",
		Long:  `Add, remove, or list account.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			/*validationError := services.ValidateSearchFlags(flags)
			if validationError != nil {
				return throwError("failed to validate search flags", validationError)
			}*/

			appService, err := services.InjectAppService(cmdConfig.Injector)
			if err != nil {
				return throwError("failed to inject app service", err)
			}

			results, err := appService.Connect(cmdConfig.Context, cmdConfig.Database, flags)
			if err != nil {
				return throwError("failed to connect app", err)
			}

			err = displayOutput(cmdConfig.Format, results)
			if err != nil {
				return throwError("failed to display output", err)
			}
			return nil
		},
	}

	appCommands.AddCommand(connectCommand)

	return appCommands
}

func displayOutput(format string, results interface{}) error {
	switch format {
	case "json":
		err := utils.DisplayJSONOutput(results)
		if err != nil {
			return err
		}
	default:
		fmt.Println(results)
	}

	return nil
}

func throwError(message string, error error) error {
	return fmt.Errorf("cmd: %s: %s: %w", mainCommand, message, error)
}
