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

var staticParg *Parg

// New returns a clean instance of Parg
func New() *Parg {
	var parg Parg
	parg.AllowedCommands = []Command{}
	parg.GlobalFlags = []Flag{}
	staticParg = &parg
	return &parg
}

// AddGlobalFlag appends an allowed optional flag for all commands to the existing set
func (p *Parg) AddGlobalFlag(flag Flag) {
	p.GlobalFlags = append(p.GlobalFlags, flag)
}

// SetGlobalFlags overwrites the allowed optional flags for all commands
func (p *Parg) SetGlobalFlags(flags []Flag) {
	p.GlobalFlags = flags
}

// AddAction is a shortcut for adding an empty command with action
func (p *Parg) AddAction(action string, usage string) {
	var command Command
	command.Action = action
	command.Help = usage
	p.AllowedCommands = append(p.AllowedCommands, command)
}

// AddCommand appends an allowed command to expect.
// Empty set enforces no arguments, throws error if argument detected
// Adding Command with name "" allows no arguments, along with any other allowed commands
func (p *Parg) AddCommand(command Command) {
	p.AllowedCommands = append(p.AllowedCommands, command)
}

// SetCommands overwrites the allowed commands
func (p *Parg) SetCommands(commands []Command) {
	p.AllowedCommands = commands
}

// GetGlobalFlags returns allowed flags for global config
func (p *Parg) GetGlobalFlags() (allowedFlags map[string]*Flag) {
	allowedFlags = map[string]*Flag{}
	var flag *Flag
	for flagIndex := range p.GlobalFlags {
		flag = &p.GlobalFlags[flagIndex]
		for identifierIndex := range flag.Identifiers {
			allowedFlags[flag.Identifiers[identifierIndex]] = flag
		}
	}

	return
}

// GetAllowedCommands aggregates configured commands
func (p *Parg) GetAllowedCommands() (allowedCommands map[string]*Command) {
	allowedCommands = map[string]*Command{}
	var command *Command
	for i := range p.AllowedCommands {
		command = &p.AllowedCommands[i]
		allowedCommands[command.Action] = command
	}

	return
}

// Validate `p.Arguments()` will return a command for the os.Args provided with the configured Parg instance
// error if fails to validate config
func Validate() (*Command, error) {
	var argV = os.Args

	if staticParg == nil {
		staticParg = New()
	}

	return staticParg.validate(argV)
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
	var args = []*Argument{}
	var flags = map[string]*Flag{}
	var help = ""

	allowedFlags := p.GetGlobalFlags()
	allowedCommands := p.GetAllowedCommands()

	if cmd, ok := allowedCommands[""]; ok {
		help = cmd.Help
	} else {
		help = Help()
	}

	var curFlag *Flag
	var arg *string

	for i := 1; i < len(argV); i++ {
		arg = &argV[i]

		if strings.HasPrefix(*arg, "-") {
			// Check if allows flag
			var newFlag *Flag
			for _, allowedFlag := range allowedFlags {
				for index := range allowedFlag.Identifiers {
					if allowedFlag.Identifiers[index] == *arg {
						// We have a winner!
						// newFlag is old, use old flag
						// Create new flag instance
						var ok bool
						if newFlag, ok = flags[allowedFlag.Name]; !ok {
							newFlag = &Flag{
								Name:        allowedFlag.Name,
								Identifiers: allowedFlag.Identifiers,
								Type:        allowedFlag.Type,
								Help:        allowedFlag.Help,
							}
						}
						break
					}
				}

				if newFlag != nil {
					// We already got one!
					break
				}
			}

			if newFlag == nil {
				// Miss job
				return nil, fmt.Errorf("invalid flag <" + *arg + "> encountered")
			}

			// Add the flag
			flags[newFlag.Name] = newFlag

			if newFlag.Type == BOOL {
				// Existence is sufficient, no trailing args expected
				newFlag.Value = true
				curFlag = nil
			} else {
				// Set flag and append trailing values
				curFlag = newFlag
			}
		} else {
			if curFlag != nil {
				// Flag set, but check if this is an action
				shouldParse := true
				if len(action) == 0 {
					// Parse action (or lack thereof)
					if _, ok := allowedCommands[*arg]; ok {
						// This is an allowed action, check for other candidates
						foundAction := false
						for x := i + 1; x < len(argV); x++ {
							if _, ok := allowedCommands[argV[x]]; ok {
								// We have another candidate, we're probably ok to treat this as a param
								foundAction = true
								break
							}
						}
						if !foundAction {
							shouldParse = false
						}
					} else {
						// Not an action.. probably a flag param
					}
				}

				if !shouldParse {
					// This is probably a command, let's skip parsing this arg
					curFlag = nil
				} else if err := curFlag.Parse(*arg); err == nil {
					// We parsed this arg!
					continue
				} else {
					// We can't parse this arg... fall through
					curFlag = nil
				}
			}

			// No flag set, this is an action or an arg
			if len(action) == 0 {
				// Parse action (or lack thereof)
				if cmd, ok := allowedCommands[*arg]; ok || len(*arg) == 0 && len(allowedCommands) == 0 {
					// Set command
					curCommand = cmd
					action = cmd.Action
					help = cmd.Help
				} else {
					return nil, fmt.Errorf("invalid command <" + *arg + "> encountered")
				}

			} else {
				// Argument?
				var argument *Argument
				argCount := len(args)
				if curCommand.Arguments != nil {
					if argCount >= len(curCommand.Arguments) {
						// We've exceeded our argument limit
						return nil, fmt.Errorf("invalid argument count: no rules for argument <" + *arg + ">")
					}

					argument = curCommand.Arguments[len(args)]
				} else {
					if argCount > 0 && args[argCount-1].Type == STRINGS {
						argument = args[argCount-1]
					} else {
						argument = &Argument{Name: *arg, Type: DEFAULT}
					}
				}

				if err := argument.Parse(*arg); err != nil {
					return nil, err
				}

				args = append(args, argument)
			}
		}
	}

	if _, ok := allowedCommands[action]; ok || len(action) == 0 && len(allowedCommands) == 0 {
		// Command allowed
	} else {
		return nil, fmt.Errorf("invalid command <" + action + "> encountered")
	}
	return &Command{action, args, flags, help}, nil
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
			flag.Value = val[0]
		default:
			flag.Type = STRINGS
		}

		command.Flags[key] = &flag
	}

	debug("\nProcessed:", toString(command))
	return command
}

var shouldPrint = false

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
