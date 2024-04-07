package search

import (
	"black-wings/internal/services"
	"black-wings/internal/types"
	"black-wings/internal/utils"
	"fmt"

	"github.com/spf13/cobra"
)

const mainCommand = "search"

func NewSearchCommand(cmdConfig types.CommandConfiguration) *cobra.Command {
	var flags types.SearchFlags

	searchCommand := &cobra.Command{
		Use:   mainCommand,
		Short: "A command line utility to search everywhere.",
		Long:  `A command line utility to search everywhere.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			validationError := services.ValidateSearchFlags(flags)
			if validationError != nil {
				return throwError("failed to validate search flags", validationError)
			}

			searchService, err := services.InjectSearchService(cmdConfig.Injector)
			if err != nil {
				return throwError("failed to inject search service", err)
			}

			results, err := searchService.Search(cmdConfig.Context, cmdConfig.Database, flags)
			if err != nil {
				return throwError("failed to search", err)
			}

			err = displayOutput(cmdConfig.Format, results)
			if err != nil {
				return throwError("failed to display output", err)
			}
			return nil
		},
	}

	createSearchFlags(searchCommand, &flags)

	return searchCommand
}

func createSearchFlags(command *cobra.Command, flags *types.SearchFlags) {
	command.Flags().StringVarP(&flags.Query, "query", "q", "", "Search query")
	command.Flags().StringSliceVarP(&flags.Apps, "apps", "a", []string{}, "Apps to search")
	command.Flags().StringVarP(&flags.Include, "include", "i", "", "Results must include")
	command.Flags().StringVarP(&flags.Exclude, "exclude", "e", "", "Results must exclude")
	command.Flags().StringSliceVarP(&flags.FileTypes, "file-types", "t", []string{}, "File types to search")
	command.Flags().StringVar(&flags.Before, "before", "", "Results before date")
	command.Flags().StringVar(&flags.After, "after", "", "Results after date")
	command.Flags().StringVarP(&flags.Sort, "sort", "s", "relevance", "Sort results by")
	command.Flags().StringVarP(&flags.Order, "order", "o", "desc", "Order results by")
	command.Flags().Int64VarP(&flags.Limit, "limit", "l", 20, "Results limit")
	command.Flags().Int64VarP(&flags.Offset, "offset", "", 0, "Results offset")

	command.MarkFlagRequired("query")
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
