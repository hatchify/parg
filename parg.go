package flag

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Parg represents expected argument config.
// Add commands and flags to Parg, then call p.Arguments() to get the command or error
// Optionally call static parg.Parse() to get a generically parsed command
type Parg struct {
	// AllowedCommands indicates a command config to validate
	// Empty list implies no commands arguments are allowed
	AllowedCommands []Command
	// GlobalFlags apply to all commands
	GlobalFlags []Flag
}

// Private singleton
var parg Parg

// AddGlobalFlag appends an allowed optional flag for all commands to the existing set
func AddGlobalFlag(flag Flag) {
	parg.GlobalFlags = append(parg.GlobalFlags, flag)
}

// SetGlobalFlags overwrites the allowed optional flags for all commands
func SetGlobalFlags(flags []Flag) {
	parg.GlobalFlags = flags
}

// AddCommand appends an allowed command to expect.
// Empty set enforces no arguments, throws error if argument detected
// Adding Command with name "" allows no arguments, along with any other allowed commands
func AddCommand(command Command) {
	parg.AllowedCommands = append(parg.AllowedCommands, command)
}

// SetCommands overwrites the allowed commands
func SetCommands(commands []Command) {
	parg.AllowedCommands = commands
}

// GetGlobalFlags returns allowed flags for global config
func (p *Parg) GetGlobalFlags() (allowedFlags map[string]*Flag) {
	var flag *Flag
	for flagIndex := range p.GlobalFlags {
		*flag = p.GlobalFlags[flagIndex]
		for identifierIndex := range flag.Identifiers {
			allowedFlags[flag.Identifiers[identifierIndex]] = flag
		}
	}

	return
}

// GetAllowedCommands aggregates configured commands
func (p *Parg) GetAllowedCommands() (allowedCommands map[string]*Command) {
	var command *Command
	for i := range parg.AllowedCommands {
		*command = parg.AllowedCommands[i]
		allowedCommands[command.Action] = command
	}

	return
}

// Validate `p.Arguments()` will return a command for the os.Args provided with the configured Parg instance
// error if fails to validate config
func (p *Parg) Validate() (*Command, error) {
	var argV = os.Args

	return p.validate(argV)
}

// Simple will return a command for the os.Args provided with default parse configuration
func Simple() *Command {
	var argV = os.Args

	return simpleParse(argV)
}

// validate `p.Arguments()` returns parsed command or error if does not match configured values
func (p *Parg) validate(argV []string) (*Command, error) {
	var curCommand *Command
	var action string
	var args []*Argument
	var flags map[string]*Flag

	allowedFlags := p.GetGlobalFlags()
	allowedCommands := p.GetAllowedCommands()

	var curFlag *Flag
	var arg *string

	for i := 1; i < len(argV); i++ {
		arg = &argV[i]

		if strings.HasPrefix(*arg, "-") {
			// Check if allows flag
			if flag, ok := allowedFlags[*arg]; ok {
				if _, ok := flags[flag.Name]; !ok {
					// Add the flag
					flags[flag.Name] = flag
				}

				if flag.Type == BOOL {
					// Existence is sufficient, no trailing args expected
					curFlag = nil
				} else {
					// Set flag and append trailing values
					curFlag = flag
				}
			} else {
				return nil, fmt.Errorf("invalid flag <" + *arg + "> encountered")
			}
		} else {
			if curFlag == nil {
				// No flag set, this is an action or an arg
				if len(action) == 0 {
					// Parse action (or lack thereof)
					if cmd, ok := allowedCommands[*arg]; ok || len(*arg) == 0 && len(allowedCommands) == 0 {
						// Set command
						curCommand = cmd
						action = cmd.Action
					} else {
						return nil, fmt.Errorf("invalid command <" + *arg + "> encountered")
					}

				} else {
					// Argument?
					if len(args) == len(curCommand.Arguments)-1 {
						// We've exceeded our argument limit
						return nil, fmt.Errorf("invalid argument count: no rules for argument <" + *arg + ">")
					}

					argument := curCommand.Arguments[len(args)]
					if err := argument.Parse(*arg); err != nil {
						return nil, err
					}

					args = append(args, argument)
				}

				// No flag set to parse
				continue
			}

			if err := curFlag.Parse(*arg); err != nil {
				return nil, err
			}
		}
	}

	return &Command{action, args, flags}, nil
}

// simpleParse returns a generically parsed argument structure, with default parsing rules:
// 1) first non-flag (not preceded by a '-flagname' arg) is treated as command
// 2) last non-flag is treated as command if not yet set
// 3) multiple matching -flagname arguments are grouped together
// 4) multiple non-flag tokens are grouped with last preceding -flagname, or grouped as command args if preceded directly by command
func simpleParse(argV []string) *Command {
	debug("Got: ", argV)

	command := NewCommand()

	parsedFlags := map[string][]string{}

	gotTrailing := true
	curFlag := ""
	var arg *string
	for i := 1; i < len(argV); i++ {
		arg = &argV[i]
		debug("\nProcessing: ", *arg)

		if strings.HasPrefix(*arg, "-") {
			if !gotTrailing && len(curFlag) > 0 {
				// Append this boolean flag first
				parsedFlags[curFlag] = []string{}
			}

			// Parse flag
			debug("  Parse flag: ", *arg)
			curFlag = *arg
			gotTrailing = false
		} else {
			if gotTrailing && len(command.Action) == 0 {
				// Still need command... Cut last flag trailing args and use this one for command
				curFlag = ""
				debug("    Reset flag, use command: ", *arg)
			}

			// Parse args
			switch curFlag {
			case "":
				if len(command.Action) == 0 {
					// Command
					command.Action = *arg
					debug("  Command: ", *arg)
				} else {
					// Arg
					command.Arguments = append(command.Arguments, &Argument{*arg, DEFAULT, false, *arg})
					debug("  Argument: ", *arg)
				}
			default:
				debug("    flag: "+curFlag+" ", *arg)
				gotTrailing = true
				if flags, ok := parsedFlags[curFlag]; ok {
					parsedFlags[curFlag] = append(flags, *arg)
				} else {
					parsedFlags[curFlag] = []string{*arg}
				}
			}
		}
	}

	for key, val := range parsedFlags {
		var flag Flag
		flag.Name = key
		flag.Identifiers = []string{key}
		flag.Value = val

		switch len(val) {
		case 0:
			flag.Type = BOOL
			flag.Value = true
		case 1:
			flag.Type = DEFAULT
		default:
			flag.Type = STRINGS
		}

		command.Flags[key] = &flag
	}

	debug("\nProcessed:", toString(command))
	return command
}

var shouldPrint = true

func debug(args ...interface{}) {
	if shouldPrint {
		fmt.Println(args...)
	}
}

func toString(obj interface{}) string {
	if !shouldPrint {
		return ""
	}

	if val, ok := obj.(string); ok {
		return val
	}

	if s, err := json.MarshalIndent(obj, "", " "); err == nil {
		return string(s)
	}

	return ""
}
