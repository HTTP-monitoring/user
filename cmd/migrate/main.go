package migrate

import (
	"database/sql"
	"log"
	"user/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/spf13/cobra"
)

func Register(root *cobra.Command, cfg config.Database) {
	c := cobra.Command{
		Use:   "migrate",
		Short: "Manages database, creates and fills tables if don't exist",
		Run: func(cmd *cobra.Command, args []string) {
			database, err := sql.Open("postgres", cfg.Cstring())
			if err != nil {
				log.Fatal(err)
			}

			driver, err := postgres.WithInstance(database, &postgres.Config{})
			if err != nil {
				log.Fatal(err)
			}

			p, err := migrate.NewWithDatabaseInstance("file://./migration", "monitor", driver)
			if err != nil {
				log.Fatal(err)
			}

			if err := p.Up(); err != nil && err != migrate.ErrNoChange {
				log.Fatal(err)
			}
		},
	}

	root.AddCommand(
		&c,
	)
}