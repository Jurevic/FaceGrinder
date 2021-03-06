package cmd

import (
	"github.com/jurevic/facegrinder/pkg/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start http server with configured api",
	Long:  `Starts a http server and serves the configured api`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Run()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	// ENV CONFIGURATION
	viper.SetEnvPrefix("fg")
	viper.AutomaticEnv()

	// SERVER
	viper.SetDefault("http_listen", ":80")

	// TLS
	// If use_tls is set to no tls is not used and other settings below are not used
	viper.SetDefault("use_tls", "no")
	viper.SetDefault("https_listen", ":443")
	viper.SetDefault("cert_path", "/etc/ssl/certs/cert.crt")
	viper.SetDefault("key_path", "/etc/ssl/certs/cert.key")

	// LOG
	viper.SetDefault("log_level", "release")

	// DB
	viper.SetDefault("db_username", "facegrinder")
	viper.SetDefault("db_password", "password")
	viper.SetDefault("db_name", "facegrinder_db")
	viper.SetDefault("db_host", "localhost:5432")

	// AUTH
	viper.SetDefault("jwt_public_key_path", "/usr/src/keys/jwtRS256.key.pub")
	viper.SetDefault("jwt_private_key_path", "/usr/src/keys/jwtRS256.key")
}
