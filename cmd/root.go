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
	db, err := config.NewDatabase("/data/identifier.sqlite")
	if err != nil {
		panic("Failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate()

	repo := sqlite.NewRepository(db.DB)
	svc := services.NewService(repo)

	config.SeedDefaultLibrarian(svc.Auth)

	config.SeedData(svc.Book, svc.Student, svc.Rent)

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
	//r.Use(middlewareoapi.OapiRequestValidatorWithOptions(swagger, &middlewareoapi.Options{
	//	Options: openapi3filter.Options{
	//		ExcludeRequestBody:    false,
	//		ExcludeResponseBody:   false,
	//		IncludeResponseStatus: true,
	//		AuthenticationFunc:    authFun,
	//	},
	//}))

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

	//api.HandlerFromMux(h, r)

	go func() {
		ticker := time.NewTicker(1 * time.Hour)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				svc.Auth.CleanupExpiredSessions()
			}
		}
	}()

	s := &http.Server{
		Handler: r,
		Addr:    net.JoinHostPort("0.0.0.0", "8080"),
	}

	log.Fatal(s.ListenAndServe())
}
