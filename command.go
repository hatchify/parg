package flag

import (
	"fmt"
	"os"
	"strings"
)

// Command represents an allowed command in the structure:
// `./<program> <optional -flags (single)> <command (single)> <optional arguments (allows multiple)> <optional -flags (allows multiple)>`
// Lack of commands implies no commands are allowed
// Command with Action="" implies program can run without command, or also with any other available commands
type Command struct {
	// Action is the parsed primary argument
	Action string `json:"action"`

	// Arguments used for defining config values and for returning optional and required trailing args
	Arguments []*Argument `json:"arguments,omitempty"`

	// Flags returned by matched action instance
	Flags map[string]*Flag `json:"flags,omitempty"`

	// Details regarding command usage
	Help string
}

// NewCommand returns a new instance of an empty command
func NewCommand() (cmd *Command) {
	var command Command
	command.Arguments = []*Argument{}
	command.Flags = map[string]*Flag{}

	cmd = &command
	return
}

// Help will return all available commands and flags
func Help() string {
	msg := "\nCommands:\n"
	for _, cmd := range staticParg.AllowedCommands {
		if strings.TrimSpace(cmd.Action) != "" {
			msg += fmt.Sprintf(" %s %s\n  :: ", os.Args[0], cmd.Action)
		} else {
			msg += " <no command>\n  :: "
		}
		msg += cmd.Help
		msg += "\n\n"
	}

	msg += "Flags:\n"
	for _, flag := range staticParg.GlobalFlags {
		msg += fmt.Sprintf(" %s\n  :: %s\n\n", flag.Identifiers, flag.Help)
	}

	return msg
}

// ShowHelp will return a command's help. If help is the command, returns first arg or general help
func (cmd *Command) ShowHelp() string {
	msg := "\nCommand: "
	if cmd.Action == "help" {
		if len(cmd.Arguments) == 0 {
			// Show regular help
			if len(cmd.Flags) == 0 {
				return Help()
			}
		} else {
			for _, argCmd := range staticParg.AllowedCommands {
				if name, ok := cmd.Arguments[0].Value.(string); ok {
					if argCmd.Action == name {
						// Show help for this cmd
						cmd.Action = argCmd.Action
						cmd.Help = argCmd.Help
						break
					}
				}
				if argCmd.Action == cmd.Arguments[0].Name {
					// Show help for this cmd
					cmd.Action = argCmd.Action
					cmd.Help = argCmd.Help
					break
				}
			}
		}
	}

	if cmd.Action == "help" {
		msg = ""
	} else {
		msg += cmd.Action + "\n"
		if strings.TrimSpace(cmd.Action) != "" {
			msg += fmt.Sprintf(" %s %s\n  :: ", os.Args[0], cmd.Action)
		} else {
			msg += " <no command>\n  :: "
		}
		msg += cmd.Help
		msg += "\n"
	}

	if len(cmd.Flags) > 0 {
		msg += "\nFlags:\n"
		for _, flag := range cmd.Flags {
			msg += fmt.Sprintf(" %s\n  :: %s\n\n", flag.Identifiers, flag.Help)
		}
	}

	return msg
}

// Args returns array of arg names
func (cmd *Command) Args() []string {
	args := make([]string, len(cmd.Arguments))
	for i, argument := range cmd.Arguments {
		args[i] = argument.Name
	}

	return args
}

// StringsFrom parses []string from flags["flagIdentifier"]
func (cmd *Command) StringsFrom(flagIdentifier string) (vals []string) {
	flag, ok := cmd.Flags[flagIdentifier]
	if !ok {
		return
	}

	vals, ok = flag.Value.([]string)
	if !ok {
		return
	}
	return
}

// StringFrom parses string from flags["flagIdentifier"]
func (cmd *Command) StringFrom(flagIdentifier string) (val string) {
	flag, ok := cmd.Flags[flagIdentifier]
	if !ok {
		return
	}

	val, ok = flag.Value.(string)
	if !ok {
		return
	}
	return
}

// IntsFrom parses []int from flags["flagIdentifier"]
func (cmd *Command) IntsFrom(flagIdentifier string) (vals []int) {
	flag, ok := cmd.Flags[flagIdentifier]
	if !ok {
		return
	}

	vals, ok = flag.Value.([]int)
	if !ok {
		return
	}
	return
}

// IntFrom parses int from flags["flagIdentifier"]
func (cmd *Command) IntFrom(flagIdentifier string) (val int) {
	flag, ok := cmd.Flags[flagIdentifier]
	if !ok {
		return
	}

	val, ok = flag.Value.(int)
	if !ok {
		return
	}
	return
}

// BoolFrom parses string from flags["flagIdentifier"]
func (cmd *Command) BoolFrom(flagIdentifier string) (val bool) {
	flag, ok := cmd.Flags[flagIdentifier]
	if !ok {
		return
	}

	val, ok = flag.Value.(bool)
	if !ok {
		return
	}
	return
}
