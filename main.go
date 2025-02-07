package main

import (
	_ "cats/docs"
	"cats/internal/domain/usecase"
	"cats/internal/infrastructure/api"
	"cats/internal/infrastructure/database/service"
	"cats/internal/rest"
	"context"
	"embed"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"log"
	"os"
	"strings"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

// @title Spy Cats API
// @version 1.0
// @host 127.0.0.1:3000
// @BasePath /
func main() {
	db := os.Getenv("DATABASE_URL")
	if db == "" {
		fmt.Println("DATABASE_URL environment variable not set, using the default value")
		db = "user=postgres password=postgres host=localhost port=5432 dbname=postgres sslmode=disable pool_max_conns=10"
	}

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("PORT environment variable not set, using the default value")
		port = ":3000"
	}

	migrate := os.Getenv("MIGRATE")

	dbpool, err := pgxpool.New(context.Background(), db)
	if err != nil {
		log.Fatal("Unable to create connection pool: %v\n", err)
	}
	defer dbpool.Close()

	if strings.ToUpper(migrate) == "TRUE" {
		goose.SetBaseFS(embedMigrations)

		if err := goose.SetDialect("postgres"); err != nil {
			panic(err)
		}

		db := stdlib.OpenDBFromPool(dbpool)
		if err := goose.Up(db, "migrations"); err != nil {
			panic(err)
		}
		if err := db.Close(); err != nil {
			panic(err)
		}

	}

	bs := api.NewBreedsService()

	cs := service.NewCatsService(dbpool)
	ms := service.NewMissionsService(dbpool)

	cuc := usecase.NewCatsUseCase(bs, cs)
	muc := usecase.NewMissionsUseCase(ms, cs)

	if err := rest.New(cuc, muc).Run(port); err != nil {
		log.Fatal(err)
	}
}
