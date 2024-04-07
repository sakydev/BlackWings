package account

import (
	"black-wings/internal/services"
	"black-wings/internal/types"

	"github.com/spf13/cobra"
)

const deleteCommandName = "delete"

func NewAccountDeleteCommand(cmdConfig types.CommandConfiguration) *cobra.Command {
	var accountIdentifier string

	listCommand := &cobra.Command{
		Use:   deleteCommandName,
		Short: "List connected accounts",
		Long:  `List all connected accounts or filter by app.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			accountService, err := services.InjectAccountService(cmdConfig.Injector)
			if err != nil {
				return throwError("failed to account app service", err)
			}

			accountID, err := accountService.GetIDByIdentifier(cmdConfig.Context, cmdConfig.Database, accountIdentifier)
			if err != nil {
				if err.Error() == "no rows in result set" {
					return throwError("no account found with that identifier", err)
				}

				return throwError("failed to get account ID by name", err)
			}

			err = accountService.Delete(cmdConfig.Context, cmdConfig.Database, accountID)
			if err != nil {
				return throwError("failed to delete accounts", err)
			}

			err = displayOutput(cmdConfig.Format, "Account deleted successfully")
			if err != nil {
				return throwError("failed to display output", err)
			}
			return nil
		},
	}

	listCommand.Flags().StringVarP(&accountIdentifier, "identifier", "a", "", "Apps to filter by")

	return listCommand
}
