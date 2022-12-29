package main

import (
	"database/sql"
	"embed"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/rs/cors"
	"go.uber.org/zap"

	config "backend/config"
	routes "backend/routes"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	runMigration := flag.String("migration", "", "Flag to check if Migrations need to Run")

	flag.Parse()
	if runMigration != nil {
		runMigrations()
	}
	// Connect DB
	config.Connect()

	// Init Router
	router := gin.Default()

	c:=cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:3000"},
        AllowCredentials: true,
        AllowedMethods: []string{"GET","POST","PUT","DELETE"},
    })
	routes.Routes(router)
    handler:=c.Handler(router)	

	log.Fatal((http.ListenAndServe(":9090",handler)))
    http.Handle("/",router)
	
	// Route Handlers / Endpoints



}

func runMigrations() {

	migrationUserDbUrl := "postgres://postgres:postgres@localhost:5432/cargo?sslmode=disable"
	if strings.TrimSpace(migrationUserDbUrl) == "" {
		fmt.Println("MIGRATIONS_USER_DB_URL is not provided")
		os.Exit(1)
	}
	db, err := sql.Open("postgres", migrationUserDbUrl)
	if err != nil {
		fmt.Println("PG DB Connection Failed", zap.Error(err))
		panic(err)
	}
	defer db.Close()
	// setup database
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		fmt.Println("Setting Goose Postgres Dialect Failed", zap.Error(err))
		panic(err)
	}
	if err := goose.Up(db, "migrations"); err != nil {
		fmt.Println("Goose Up Failed", zap.Error(err))
		panic(err)
	}
}
