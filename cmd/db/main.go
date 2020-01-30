package main

import (
	"os"

	"github.com/galihrivanto/dbutil"
	_ "github.com/galihrivanto/dbutil/plugin/ping"
)

func main() {
	if err := dbutil.Execute(); err != nil {
		os.Exit(1)
	}
}
