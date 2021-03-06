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

	handler func(cmd *Command) (err error)

	// Details regarding command usage
	helpDetails string
}

// NewCommand returns a new instance of an empty command
func NewCommand() (cmd *Command) {
	var command Command
	command.Arguments = []*Argument{}
	command.Flags = map[string]*Flag{}
	command.handler = nil

	cmd = &command
	return
}

// Help will return all available commands and flags
func Help(markdown bool) string {
	var prefix = ""
	if markdown {
		prefix = "#"
	}

	var doublePrefix = prefix + prefix
	var triplePrefix = doublePrefix + prefix

	msg := "\n" + doublePrefix + " Commands\n\n"
	for _, cmd := range staticParg.AllowedCommands {
		if strings.TrimSpace(cmd.Action) != "" {
			msg += fmt.Sprintf("%s %s %s\n  :: ", triplePrefix, os.Args[0], cmd.Action)
		} else {
			msg += triplePrefix + " " + os.Args[0] + "\n  :: "
		}
		msg += cmd.helpDetails
		msg += "\n\n"
	}

	if len(staticParg.GlobalFlags) > 0 {
		msg += doublePrefix + " Flags\n"
		for _, flag := range staticParg.GlobalFlags {
			msg += fmt.Sprintf("\n%s %s\n  :: %s\n", triplePrefix, flag.Identifiers, flag.Help)
		}
	}

	return msg
}

// Help will return a command's help. If help is the command, returns first arg or general help
func (cmd *Command) Help(markdown bool) string {
	var prefix = ""
	if markdown {
		prefix = "#"
	}

	var doublePrefix = prefix + prefix
	var triplePrefix = doublePrefix + prefix

	msg := "\n" + doublePrefix + " Command: "
	if cmd.Action == "help" {
		if len(cmd.Arguments) == 0 {
			// Show regular help
			if len(cmd.Flags) == 0 {
				return Help(true)
			}
		} else {
			for _, argCmd := range staticParg.AllowedCommands {
				if name, ok := cmd.Arguments[0].Value.(string); ok {
					if argCmd.Action == name {
						// Show help for this cmd
						cmd.Action = argCmd.Action
						cmd.helpDetails = argCmd.helpDetails
						break
					}
				}
				if argCmd.Action == cmd.Arguments[0].Name {
					// Show help for this cmd
					cmd.Action = argCmd.Action
					cmd.helpDetails = argCmd.helpDetails
					break
				}
			}
		}
	}

	if cmd.Action == "help" {
		if len(cmd.Flags) == 0 {
			msg = Help(true)
		} else {
			msg = ""
		}
	} else {
		msg += cmd.Action + "\n\n"
		msg += fmt.Sprintf("%s %s %s\n  :: ", triplePrefix, os.Args[0], cmd.Action)
		msg += cmd.helpDetails
		msg += "\n"
	}

	if len(cmd.Flags) > 0 {
		msg += "\n" + doublePrefix + " Flags: "
		output := ""
		for _, flag := range cmd.Flags {
			msg += flag.Name + " "
			output += fmt.Sprintf("\n%s %s\n  :: %s\n", triplePrefix, flag.Identifiers, flag.Help)
		}
		msg += "\n" + output
	}

	if cmd.Action == "help" {
		if len(cmd.Arguments) > 0 {
			msg += "\nError parsing arguments: invalid command <" + cmd.Arguments[0].Name + "> encountered"
		}
	}

	return msg
}

// Exec will run handler
func (cmd *Command) Exec() (err error) {
	if cmd.handler == nil {
		err = fmt.Errorf("unable to exec cmd \"%s\": no handler set", cmd.Action)
	}

	return cmd.handler(cmd)
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
