package main

import (
	"BlackWings/cmd/search"
	"BlackWings/internal"
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/samber/do"
	"github.com/spf13/cobra"
)

const AppName = "blackwings"

func main() {
	ctx := context.Background()
	injector := do.DefaultInjector
	setup(injector)

	var format string

	coreCommand := &cobra.Command{
		Use:   AppName,
		Short: "Search everywhere",
		Long:  `A command line utility to search everywhere.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Use '%s --help' for more information.\n", AppName)
		},
	}

	coreCommand.PersistentFlags().StringVarP(&format, "format", "f", "json", "Data format to use (default: json)")

	searchCommand := search.NewSearchCommand(ctx, format, injector)

	coreCommand.AddCommand(searchCommand)

	err := coreCommand.Execute()
	if err != nil {
		fmt.Println(color.RedString("Error: %v\n", err))
		os.Exit(1)
	}
}

func setup(injector *do.Injector) {
	internal.WireDependencies(injector)
}
