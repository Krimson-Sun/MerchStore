package main

import (
	"context"
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"merch-store/internal/app"
	"merch-store/internal/db"
	"merch-store/internal/jwt"
	"merch-store/internal/logger"
	"merch-store/internal/repository"
	"merch-store/internal/service"
	"merch-store/internal/tracer"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func init() {
	logger.Init()
	godotenv.Load()
	log.SetOutput(io.Discard)
}

func loadPostgresURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_SSL_MODE"),
	)
}

func Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tracer.MustSetup(
		ctx,
		tracer.WithServiceName("merch-store"),
		tracer.WithCollectorEndpoint(os.Getenv("JAEGER_COLLECTOR_ENDPOINT")),
	)

	postgresURL := loadPostgresURL()

	pool, err := pgxpool.New(ctx, postgresURL)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		logger.Fatal(err.Error())
	}

	ContextManager := db.NewContextManager(pool)

	Repo := repository.NewPGXRepository(ContextManager)

	jwtSecret := os.Getenv("JWT_SECRET")

	JWTProvider := jwt.NewProvider(
		jwt.WithCredentials(
			jwt.NewSecretCredentials(jwtSecret),
		),
		jwt.WithAccessTTL(
			30*time.Minute,
		),
	)

	Service := service.New(
		ContextManager,
		JWTProvider,
		Repo, // Auth
		Repo, // User
	)

	App := app.New(
		Service,
		Service,
		app.WithHTTPPathPrefix("/api"),
	)

	if err := App.Run(ctx); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		panic(err)
	}
}
