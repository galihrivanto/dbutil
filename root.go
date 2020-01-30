package dbutil

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func Execute() error {
	rootCmd := &cobra.Command{
		Use:   filepath.Base(os.Args[0]),
		Short: "DB Utility",
		Long:  "Swiss-Army-Knife Database utility tool",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	rootCmd.PersistentFlags().StringP("type", "t", "postgres", "database type")
	rootCmd.PersistentFlags().String("host", "localhost", "database target hostname")
	rootCmd.PersistentFlags().Int("port", 5432, "database target port")
	rootCmd.PersistentFlags().StringP("user", "u", "user", "user credential")
	rootCmd.PersistentFlags().StringP("password", "w", "", "password credential")
	rootCmd.PersistentFlags().StringP("dbname", "d", "mydb", "database name")

	// get sub command
	for _, factory := range factories {
		rootCmd.AddCommand(factory(configFromCommandFlags))
	}

	return rootCmd.Execute()
}

func configFromCommandFlags(cmd *cobra.Command) Config {
	_type, _ := cmd.Flags().GetString("type")
	host, _ := cmd.Flags().GetString("host")
	port, _ := cmd.Flags().GetInt("port")
	user, _ := cmd.Flags().GetString("user")
	password, _ := cmd.Flags().GetString("password")
	dbname, _ := cmd.Flags().GetString("dbname")

	return Config{
		"type":     _type,
		"host":     host,
		"port":     port,
		"user":     user,
		"password": password,
		"dbname":   dbname,
	}
}
