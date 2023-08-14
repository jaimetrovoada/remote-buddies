package main

import (
	"context"
	"fmt"
	"os"
	"remote-buddies/server/internal/db"
	"remote-buddies/server/internal/routes"

	"github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	dbUrl := os.Getenv("DATABASE_URL")
	fmt.Printf("%s\n", dbUrl)

	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	query := db.New(conn)

	router := routes.NewRouter(query)

	router.Listen(":8000")

}
