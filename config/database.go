package config

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	_ "entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"log"
	"os"
	"todo-list/ent"
	"todo-list/helper"
)

func ConnectDB() *ent.Client {
	err := godotenv.Load(".env")
	helper.ErrorPanic("failed to load .env", err)
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL env variable not set")
	}

	db, err := sql.Open("pgx", databaseURL)
	helper.ErrorPanic("failed to connect to database", err)

	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		helper.ErrorPanic("failed to create database schema", err)
	}
	return ent.NewClient(ent.Driver(drv))

}

func DisconnectDB(client *ent.Client) {
	if err := client.Close(); err != nil {
		helper.ErrorPanic("Failed to close client", err)
	}
}
