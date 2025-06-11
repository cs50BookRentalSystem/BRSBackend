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
	middleware "github.com/oapi-codegen/nethttp-middleware"
	"github.com/spf13/cobra"

	"BRSBackend/pkg/api"
	"BRSBackend/pkg/config"
	"BRSBackend/pkg/handlers"
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
	db, err := config.NewDatabase("./")
	if err != nil {
		panic("Failed to connect database")
	}
	defer db.Close()

	repo := sqlite.NewRepository(db.DB)
	svc := services.NewBookService(repo.Book)
	h := handlers.NewBookHandler(svc)

	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	r := chi.NewRouter()
	r.Use(middleware.OapiRequestValidator(swagger))
	api.HandlerFromMux(h, r)

	s := &http.Server{
		Handler: r,
		Addr:    net.JoinHostPort("0.0.0.0", "8080"),
	}

	log.Fatal(s.ListenAndServe())
}
