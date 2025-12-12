package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	commentapi "github.com/iviv660/wb-CommentTree.git/internal/api/comment"
	commentrepo "github.com/iviv660/wb-CommentTree.git/internal/repository/comment"
	commentservice "github.com/iviv660/wb-CommentTree.git/internal/service/comment"
)

func getEnv(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}

func main() {
	_ = godotenv.Load(".env")

	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPass := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "wb_comment_tree")
	dbSSL := getEnv("DB_SSLMODE", "disable")

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPass, dbHost, dbPort, dbName, dbSSL,
	)

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}
	defer pool.Close()

	repo := commentrepo.NewRepository(pool)
	svc := commentservice.New(repo)

	mux := chi.NewRouter()
	api := commentapi.NewAPI(mux, svc)
	api.RegisterHandler()

	addr := getEnv("APP_PORT", ":8080")
	log.Printf("listening on %s", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
