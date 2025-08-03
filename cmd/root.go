package cmd

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/chi/v5"
	middlewareoapi "github.com/oapi-codegen/nethttp-middleware"
	"github.com/spf13/cobra"

	"BRSBackend/pkg/api"
	"BRSBackend/pkg/config"
	"BRSBackend/pkg/handlers"
	"BRSBackend/pkg/middleware"
	"BRSBackend/pkg/repository/sqlite"
	"BRSBackend/pkg/services"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "BRSBackend",
	Short: "A brief description of your application",
	Run:   runCommand,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is config.dev.yaml)")
}

func runCommand(cmd *cobra.Command, args []string) {
	if cfgFile == "" {
		cfgFile = "config.dev.yaml"
	}
	cfg, err := config.LoadConfig(cfgFile)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := initializeDatabase(cfg.Database.DSN)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	repo := sqlite.NewRepository(db.DB)
	svc := services.NewService(repo, cfg.Rent.RentalDays)

	seedData(svc, cfg)

	r := setupRouter(svc)

	go startCleanupRoutine(svc.Auth)

	server := config.NewServer(net.JoinHostPort("0.0.0.0", cfg.Server.Port), r)
	server.Start()
}

func initializeDatabase(dsn string) (*config.Database, error) {
	db, err := config.NewDatabase(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	db.AutoMigrate()
	return db, nil
}

func seedData(svc *services.Service, cfg *config.AppConfig) {
	env := cfg.Server.Env
	if env != "prod" {
		config.SeedData(svc.Book, svc.Student, svc.Rent)
	}

	if cfg.Librarian.User != "" && cfg.Librarian.Pass != "" {
		config.SeedLibrarian(svc.Auth, cfg.Librarian.User, cfg.Librarian.Pass)
	} else if env != "prod" {
		config.SeedDefaultLibrarian(svc.Auth)
	}
}

func setupRouter(svc *services.Service) *chi.Mux {
	h := handlers.NewHandler(svc)

	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}
	swagger.Servers = nil

	authFun := middleware.NewOApiAuthenticationFunc(svc.Auth)

	r := chi.NewRouter()
	r.Use(middleware.Cors())

	r.Get("/swagger/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/swagger/", http.FileServer(http.FS(api.SwaggerUI))).ServeHTTP(w, r)
	})

	r.Group(func(r chi.Router) {
		r.Use(middlewareoapi.OapiRequestValidatorWithOptions(swagger, &middlewareoapi.Options{
			Options: openapi3filter.Options{
				ExcludeRequestBody:    false,
				ExcludeResponseBody:   false,
				IncludeResponseStatus: true,
				AuthenticationFunc:    authFun,
			},
		}))
		api.HandlerFromMux(h, r)
	})

	return r
}

func startCleanupRoutine(authService services.AuthService) {
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			authService.CleanupExpiredSessions()
		}
	}
}
