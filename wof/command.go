// Package wof registers and constructs subcommands for the wof tool.
package wof

import (
	"context"
	"sort"
	"strings"

	"github.com/aaronland/go-roster"
)

var command_roster roster.Roster

// Command runs a wof subcommand with the supplied command-line arguments.
type Command interface {
	Run(context.Context, []string) error
}

// CommandInitializationFunc is a function defined by individual command package and used to create
// an instance of that command
type CommandInitializationFunc func(ctx context.Context, cmd string) (Command, error)

// RegisterCommand registers 'scheme' as a key pointing to 'init_func' in an internal lookup table
// used to create new `Command` instances by the `NewCommand` method.
func RegisterCommand(ctx context.Context, scheme string, init_func CommandInitializationFunc) error {

	err := ensureCommandRoster()

	if err != nil {
		return err
	}

	return command_roster.Register(ctx, scheme, init_func)
}

func ensureCommandRoster() error {

	if command_roster == nil {

		r, err := roster.NewDefaultRoster()

		if err != nil {
			return err
		}

		command_roster = r
	}

	return nil
}

// NewCommand constructs the command registered under cmd with [RegisterCommand].
func NewCommand(ctx context.Context, cmd string) (Command, error) {

	i, err := command_roster.Driver(ctx, cmd)

	if err != nil {
		return nil, err
	}

	init_func := i.(CommandInitializationFunc)
	return init_func(ctx, cmd)
}

// Commands returns the list of command that have been registered.
func Commands() []string {

	ctx := context.Background()
	commands := []string{}

	err := ensureCommandRoster()

	if err != nil {
		return commands
	}

	for _, dr := range command_roster.Drivers(ctx) {
		cmd := strings.ToLower(dr)
		commands = append(commands, cmd)
	}

	sort.Strings(commands)
	return commands
}
