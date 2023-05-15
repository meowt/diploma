package database

import (
	"errors"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func Setup() (PostgresClient *sqlx.DB, err error) {
	PostgresClient, err = sqlx.Open("postgres", viper.GetString("postgres.postgresDsn"))
	if err != nil {
		return
	}

	if err = PostgresClient.Ping(); err != nil {
		err = errors.Join(fmt.Errorf("Error while pinging database: "), err)
		return
	}
	log.Println("Successfully connected to Postgres")

	if err = Deploy(PostgresClient); err != nil {
		return
	}
	return
}

func Deploy(PostgresClient *sqlx.DB) (err error) {
	for _, command := range viper.GetStringSlice("postgres.deployment") {
		if _, err = PostgresClient.Exec(command); err != nil {
			return errors.Join(fmt.Errorf("Error while deploying database: "), err)
		}
		log.Printf("Success: %s\n", command)
	}
	log.Println("Successfully deployed database")
	return
}
