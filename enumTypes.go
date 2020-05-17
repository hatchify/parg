package flag

// ArgType indicates format of flag arguments
type ArgType string

const (
	// DEFAULT expects exactly 1 string argument
	DEFAULT ArgType = ""

	// BOOL expects no flag arguments, just the flag itself
	BOOL = "bool"

	// STRINGS expects at least 1 or more string arguments
	STRINGS = "[]string"

	// INT expects exactly 1 number argument
	INT = "int"

	// INTS expects at least 1 or more number arguments
	INTS = "[]int"
)

// Expects returns a string indicating what the type should parse
func (a *ArgType) Expects() string {
	switch a {
	case DEFAULT:
		return "a single string"
	case STRINGS:
		return "one or more strings"
	case INT:
		return "a single integer"
	case INTS:
		return "one or more integers"
	case BOOL:
		return "no trailing arguments"
	default:
		return "unknown"
	}
}
