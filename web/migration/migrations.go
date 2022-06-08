package migration

import (
	"database/sql"
	"log"
	"os"
)

type Database struct {
	db sql.DB
}

type Migrate interface {
	MigrateUp()
	MigrateDown()
}

func NewMigration(builderType string) Migrate {
	if builderType == "sqlite" {
		conn := sqliteConn()
		return &Sqlite{
			database: Database{
				db: conn,
			},
		}
	} else if builderType == "mysql" {
		conn := mysqlConn()
		return &MySQL{
			database: Database{
				db: conn,
			},
		}
	}
	return nil
}

func sqliteConn() sql.DB {
	db, err := sql.Open("sqlite3", "./rubboto_go.sql")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	return *db
}

func mysqlConn() sql.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	db, err := sql.Open("mysql", "host")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	return *db
}
