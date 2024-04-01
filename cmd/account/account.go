package account

import (
	"BlackWings/internal/types"
	"BlackWings/internal/utils"
	"fmt"

	"github.com/spf13/cobra"
)

const mainCommandName = "accounts"

func NewAccountCommands(cmdConfig types.CommandConfiguration) *cobra.Command {
	accountCommands := &cobra.Command{
		Use:   mainCommandName,
		Short: "Add, remove, or list apps",
		Long:  `Add, remove, or list apps.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Use '%s --help' for more information.\n", mainCommandName)
		},
	}

	connectCommand := NewAccountConnectCommand(cmdConfig)
	listCommand := NewAccountListCommand(cmdConfig)
	deleteCommand := NewAccountDeleteCommand(cmdConfig)

	accountCommands.AddCommand(connectCommand)
	accountCommands.AddCommand(listCommand)
	accountCommands.AddCommand(deleteCommand)

	return accountCommands
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
	return fmt.Errorf("cmd: %s: %s: %w", mainCommandName, message, error)
}
