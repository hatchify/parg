package flag

import (
	"fmt"
	"strconv"
)

// Argument is used for config and returning parsed flags
type Argument struct {
	// Name of argument
	Name string `json:"name"`
	// Rules for parsing argument values
	Type ArgType `json:"type,omitempty"`
	// Throws error if required and not provided
	Required bool `json:"required,omitempty"`

	// Populated value for argument
	Value interface{} `json:"value,omitempty"`
}

// Parse attempts to set the given value for the given argument. Returns false if it does not meet type criteria
func (arg *Argument) Parse(value string) (err error) {
	switch arg.Type {
	case DEFAULT:
		// String
		arg.Value = value
		return
	case BOOL:
		// Existence is sufficient
		arg.Value = true
		return
	case INT:
		if val, err := strconv.Atoi(value); err == nil {
			arg.Value = val
		} else {
			return fmt.Errorf("Invalid value encountered. Cannot set <" + value + "> for INT type argument <" + arg.Name + ">")
		}
		return
	}

	return fmt.Errorf("Invalid value encountered. Cannot set <" + value + "> for " + string(arg.Type) + " argument <" + arg.Name + ">")
}
