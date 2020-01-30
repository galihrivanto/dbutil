package ping

import (
	"fmt"
	"os"
	"time"

	"github.com/galihrivanto/dbutil"
	"github.com/spf13/cobra"
)

var (
	dbSupports = map[string]Pinger{}
)

// Pinger define availability support
type Pinger interface {
	Ping(dbutil.Config) error
}

// New create new plugin sub command
func New(parser dbutil.ConfigParser) *cobra.Command {
	return &cobra.Command{
		Use:   "ping",
		Short: "Database availability check",
		Run: func(cmd *cobra.Command, args []string) {
			config := parser(cmd)

			t := config.GetString("type")
			pinger, supported := dbSupports[t]
			if !supported {
				fmt.Printf("Ping not supported on selected database: %s \n", t)
				os.Exit(1)
			}

			start := time.Now()
			defer func() {
				fmt.Printf("Ping done in %d ms \n", time.Now().Sub(start).Microseconds())
			}()

			if err := pinger.Ping(config); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
}

func init() {
	dbutil.RegisterCommandFactory(New)
}
