/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/cobra"

	"BRSBackend/pkg/api"
	"BRSBackend/pkg/config"
	"BRSBackend/pkg/handlers"
	"BRSBackend/pkg/middleware"
	"BRSBackend/pkg/repository/sqlite"
	"BRSBackend/pkg/services"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "BRSBackend",
	Short: "A brief description of your application",
	Run:   runCommand,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.BRSBackend.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runCommand(cmd *cobra.Command, args []string) {
	db, err := config.NewDatabase("./identifier.sqlite")
	if err != nil {
		panic("Failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate()

	repo := sqlite.NewRepository(db.DB)
	librarianRepo := sqlite.NewLibrarianRepository(db.DB)
	sessionRepo := sqlite.NewSessionRepository(db.DB)

	authService := services.NewAuthService(librarianRepo, sessionRepo)
	svc := services.NewBookService(repo.Book)

	config.SeedDefaultLibrarian(authService)

	h := handlers.NewBookHandler(svc)
	authHandler := handlers.NewAuthHandler(authService)

	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	r := chi.NewRouter()
	r.Post("/login", authHandler.Login)
	r.Post("/logout", authHandler.Logout)
	r.Group(func(r chi.Router) {
		// Apply auth middleware to protect API routes
		r.Use(middleware.AuthMiddleware(authService))
		//r.Use(middleware.OapiRequestValidator(swagger))
		// Register your generated API routes (these require authentication)
		api.HandlerFromMux(h, r)
	})
	//

	s := &http.Server{
		Handler: r,
		Addr:    net.JoinHostPort("0.0.0.0", "8080"),
	}

	log.Fatal(s.ListenAndServe())
}
