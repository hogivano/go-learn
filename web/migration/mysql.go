package migration

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
)

type MySQL struct {
	database Database
}

func (sql *MySQL) MigrateUp() {
	instance, err := mysql.WithInstance(&sql.database.db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file:///migrations", "mysql", instance)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}

func (sql *MySQL) MigrateDown() {
	instance, err := mysql.WithInstance(&sql.database.db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file:///migrations", "mysql", instance)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Down(); err != nil {
		log.Fatal(err)
	}
}
