package flag

import (
	"fmt"
	"strconv"
)

// Flag represents an allowed -flag param in the structure:
// `./<program> <optional -flags> <command> <optiona larguments> <optional -flags>`
// Lack of command flags or global flags implies no flags are allowed
type Flag struct {
	// Name of flag
	Name string `json:"name"`
	// -flag values to match this flag instance
	Identifiers []string `json:"identifiers"`
	// Rules for parsing flag values
	Type ArgType `json:"type,omitempty"`

	// Populated values for returned flags
	Value interface{} `json:"value,omitempty"`
}

// Parse attempts to set the given value for the given flag. Returns false if it does not meet type criteria
func (flag *Flag) Parse(value string) error {
	switch flag.Type {
	case DEFAULT:
		// String
		if flag.Value != nil {
			return fmt.Errorf("Redundant value encountered. Cannot set <" + value + "> for single STRING flag <" + flag.Name + "> - already contains value: " + flag.Value.(string))
		}
		flag.Value = value
	case BOOL:
		// Existence is sufficient
		flag.Value = true
	case INT:
		if flag.Value != nil {
			return fmt.Errorf("Redundant value encountered. Cannot set <" + value + "> for single INT flag <" + flag.Name + "> - already contains value: " + flag.Value.(string))
		}

		if val, err := strconv.Atoi(value); err == nil {
			// Value is number type
			flag.Value = val
		} else {
			return fmt.Errorf("Invalid value encountered. Cannot set <" + value + "> for INT flag <" + flag.Name + ">")
		}
	case INTS:
		if val, err := strconv.Atoi(value); err == nil {
			// Value is number type
			if slice, ok := flag.Value.([]int); ok {
				flag.Value = append(slice, val)
			} else {
				flag.Value = []int{val}
			}
		} else {
			return fmt.Errorf("Invalid value encountered. Cannot set <" + value + "> for INTS flag <" + flag.Name + ">")
		}
	case STRINGS:
		if slice, ok := flag.Value.([]string); ok {
			flag.Value = append(slice, value)
		} else {
			flag.Value = []string{value}
		}
	default:
		return fmt.Errorf("Invalid type encountered. Cannot set <" + value + "> for unknown type of flag <" + flag.Name + ">")
	}

	return nil
}
