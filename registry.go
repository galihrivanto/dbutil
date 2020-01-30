package dbutil

import (
	"github.com/spf13/cobra"
)

var (
	factories []CommandFactory
)

// ConfigParser provide function to resolve config
// on runtime
type ConfigParser func(*cobra.Command) Config

// CommandFactory define abstraction of sub command creation
type CommandFactory func(ConfigParser) *cobra.Command

// RegisterCommandFactory add command factory
func RegisterCommandFactory(factory CommandFactory) {
	factories = append(factories, factory)
}
