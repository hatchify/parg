package flag

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
}

// NewCommand returns a new instance of an empty command
func NewCommand() (cmd *Command) {
	var command Command
	command.Arguments = []*Argument{}
	command.Flags = map[string]*Flag{}

	cmd = &command
	return
}
