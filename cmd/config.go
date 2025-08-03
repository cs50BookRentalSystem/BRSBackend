package cmd

import (
	"errors"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Generate a default configuration file",
	Run:   generateConfig,
}

func init() {
	rootCmd.AddCommand(configCmd)
}

func generateConfig(cmd *cobra.Command, args []string) {
	v := viper.New()
	v.SetDefault("server.port", "8080")
	v.SetDefault("database.dsn", "./brs.sqlite")
	v.SetDefault("librarian.user", "admin")
	v.SetDefault("librarian.pass", "securePasswd")
	v.SetDefault("rent.overdue_period", 7)

	if err := v.SafeWriteConfigAs("config.yaml"); err != nil {
		var configFileAlreadyExistsError viper.ConfigFileAlreadyExistsError
		if errors.As(err, &configFileAlreadyExistsError) {
			log.Println("Config file already exists.")
		}
	} else {
		log.Println("Config file generated successfully: config.yaml")
	}
}
