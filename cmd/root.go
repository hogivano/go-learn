package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	roodCmd = &cobra.Command{
		Use:   "go-learn",
		Short: "Learning Golang Concept",
		Long:  "Learning Golang Concept",
	}
	migrate       bool   = false
	migrateSource string = "sqlite"
)

func init() {
	cobra.OnInitialize()

	dbCmd.Flags().BoolVar(&migrate, "migrate", false, "Migrate Database")
	dbCmd.Flags().StringVar(&migrateSource, "migrate-source", "sqlite", "Migrate Source")
	roodCmd.AddCommand(dbCmd)
}

func Execute() {
	if err := roodCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
