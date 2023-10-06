package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"remote-buddies/server/internal/config"
	"remote-buddies/server/internal/db"
	"remote-buddies/server/internal/routes"

	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := pgx.Connect(context.Background(), config.DATABASE_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	goth.UseProviders(github.New(config.GITHUB_CLIENT_ID, config.GITHUB_CLIENT_SECRET, config.GITHUB_CLIENT_CALLBCK, config.GITHUB_CLIENT_SCOPE))
	gothic.Store = sessions.NewCookieStore([]byte(config.SESSION_SECRET))

	query := db.New(conn)

	router := routes.NewRouter(query)

	if err := router.Start(fmt.Sprintf(":%s", config.PORT)); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
