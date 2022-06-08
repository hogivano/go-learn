package cmd

import (
	"github.com/hogivano/go-learn.git/web/migration"
	"github.com/spf13/cobra"
)

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Database commands",
	Long:  "All database related commands",
	Run: func(cmd *cobra.Command, args []string) {
		if migrate {
			if migrateSource == "sqlite" {
				migration.NewMigration(migrateSource).MigrateUp()
			} else if migrateSource == "mysql" {
				migration.NewMigration((migrateSource)).MigrateUp()
			}
		}
	},
}
