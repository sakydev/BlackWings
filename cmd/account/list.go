package account

import (
	"black-wings/internal/services"
	"black-wings/internal/types"

	"github.com/spf13/cobra"
)

const listCommandName = "list"

func NewAccountListCommand(cmdConfig types.CommandConfiguration) *cobra.Command {
	var appNames []string
	var appIDs []int64

	listCommand := &cobra.Command{
		Use:   listCommandName,
		Short: "List connected accounts",
		Long:  `List all connected accounts or filter by app.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			accountService, err := services.InjectAccountService(cmdConfig.Injector)
			if err != nil {
				return throwError("failed to account app service", err)
			}

			appService, err := services.InjectAppService(cmdConfig.Injector)
			if err != nil {
				return throwError("failed to inject app service", err)
			}

			availableApps, err := appService.List(cmdConfig.Context, cmdConfig.Database)
			if err != nil {
				return throwError("failed to list apps", err)
			}

			if len(appNames) != 0 {
				appIDs, err = appService.MapNamesToIDs(appNames, availableApps)
				if err != nil {
					return throwError("failed to map app names to IDs", err)
				}
			}

			results, err := accountService.List(cmdConfig.Context, cmdConfig.Database, appIDs)
			if err != nil {
				return throwError("failed to list accounts", err)
			}

			err = displayOutput(cmdConfig.Format, results)
			if err != nil {
				return throwError("failed to display output", err)
			}
			return nil
		},
	}

	listCommand.Flags().StringSliceVarP(&appNames, "apps", "a", []string{}, "Apps to filter by")

	return listCommand
}
