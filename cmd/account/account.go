package account

import (
	"BlackWings/internal/services"
	"BlackWings/internal/types"
	"BlackWings/internal/utils"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

const mainCommandName = "accounts"
const connectCommandName = "connect"

func NewAccountCommands(cmdConfig types.CommandConfiguration) *cobra.Command {
	//var flags types.AppFlags

	accountCommands := &cobra.Command{
		Use:   mainCommandName,
		Short: "Add, remove, or list apps",
		Long:  `Add, remove, or list apps.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Use '%s --help' for more information.\n", mainCommandName)
		},
	}

	connectCommand := &cobra.Command{
		Use:   connectCommandName,
		Short: "Connect a new account",
		Long:  `Connect new account for searching.`,
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

			selectedApp, err := askAppSelection(availableApps)
			if err != nil {
				return throwError("failed to ask app selection", err)
			}

			accountDetails, err := askAccountDetails(selectedApp)
			if err != nil {
				return throwError("failed to ask account details", err)
			}

			results, err := accountService.Connect(cmdConfig.Context, cmdConfig.Database, selectedApp, accountDetails)
			if err != nil {
				return throwError("failed to connect app", err)
			}

			err = displayOutput(cmdConfig.Format, fmt.Sprintf("Account connected successfully with ID: %d", results))
			if err != nil {
				return throwError("failed to display output", err)
			}
			return nil
		},
	}

	accountCommands.AddCommand(connectCommand)

	return accountCommands
}

func askAppSelection(apps map[string]types.App) (types.App, error) {
	availableAppsNames := make([]string, 0, len(apps))
	for name := range apps {
		availableAppsNames = append(availableAppsNames, name)
	}

	var selectedAppName string
	appSelectionPrompt := &survey.Select{
		Message: "Which app do you want to connect?",
		Options: availableAppsNames,
	}

	err := survey.AskOne(appSelectionPrompt, &selectedAppName, survey.WithIcons(func(icons *survey.IconSet) {
		icons.SelectFocus.Text = "\033[32m✔\033[0m" // green checkmark
	}))
	if err != nil {
		return types.App{}, throwError("failed to select app", err)
	}

	return apps[selectedAppName], nil
}

func askAccountDetails(app types.App) (types.Account, error) {
	accountDetails := types.Account{}

	prompts := []*survey.Question{
		{
			Name: "clientID",
			Prompt: &survey.Input{
				Message: "Enter Client ID:",
			},
			Validate: survey.Required,
		},
		{
			Name: "clientSecret",
			Prompt: &survey.Input{
				Message: "Enter Client Secret:",
			},
			Validate: survey.Required,
		},
		{
			Name: "raw",
			Prompt: &survey.Input{
				Message: "Enter Raw Data (if any):",
			},
		},
	}

	// Ask the questions sequentially
	err := survey.Ask(prompts, &accountDetails)
	if err != nil {
		return types.Account{}, throwError("failed to ask account details", err)
	}

	return accountDetails, nil
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
