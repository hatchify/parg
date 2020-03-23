package flag

import (
	"strings"
	"testing"

	"github.com/hatchify/simply"
)

var emptyString = ""
var emptyArguments = []*Argument{}
var emptyFlags = map[string]*Flag{}

func TestSimpleParse_Empty(context *testing.T) {
	// White box input
	input := "gomo"

	// OS-Parsed input format
	args := strings.Split(input, " ")

	// Execute test with input
	command := simpleParse(args)

	// Run validations
	test := simply.TestTarget(command.Action, "Action", context)
	result := test.Equals(emptyString)
	test.Validate(result)

	argTest := simply.Test("Arguments", context)
	argResult := argTest.Target(command.Arguments).Equals(emptyArguments)
	argTest.Validate(argResult)

	flagTest := simply.Test("Flags", context)
	flagResult := flagTest.Target(command.Flags).Equals(emptyFlags)
	flagTest.Validate(flagResult)
}

func TestSimpleParse_Cmd_1Arg_1Flag(context *testing.T) {
	input := "gomu sync mod-common -i hatchify"

	args := strings.Split(input, " ")
	command := simpleParse(args)

	expectedAction := "sync"

	actionTest := simply.Test("Action", context)
	actionResult := actionTest.Target(command.Action).Equals(expectedAction)
	actionTest.Validate(actionResult)

	var arg Argument
	arg.Name = "mod-common"
	arg.Value = "mod-common"
	expectedArguments := []*Argument{&arg}

	argTest := simply.Test("Arguments", context)
	argResult := argTest.Target(command.Arguments).Equals(expectedArguments)
	argTest.Validate(argResult)

	var flag Flag
	flag.Name = "-i"
	flag.Identifiers = []string{"-i"}
	flag.Value = []string{"hatchify"}
	expectedFlags := map[string]*Flag{"-i": &flag}

	flagTest := simply.Test("Flags", context)
	flagResult := flagTest.Target(command.Flags).Equals(expectedFlags)
	flagTest.Validate(flagResult)
}

func TestSimple_1Flag_Cmd_1Arg(context *testing.T) {
	input := "gomu -i hatchify sync mod-common"

	args := strings.Split(input, " ")
	command := simpleParse(args)

	expectedAction := "sync"

	actionTest := simply.Test("Action", context)
	actionResult := actionTest.Target(command.Action).Equals(expectedAction)
	actionTest.Validate(actionResult)

	var arg Argument
	arg.Name = "mod-common"
	arg.Value = "mod-common"
	expectedArguments := []*Argument{&arg}

	argTest := simply.Test("Arguments", context)
	argResult := argTest.Target(command.Arguments).Equals(expectedArguments)
	argTest.Validate(argResult)

	var flag Flag
	flag.Name = "-i"
	flag.Identifiers = []string{"-i"}
	flag.Value = []string{"hatchify"}
	expectedFlags := map[string]*Flag{"-i": &flag}

	flagTest := simply.Test("Flags", context)
	flagResult := flagTest.Target(command.Flags).Equals(expectedFlags)
	flagTest.Validate(flagResult)
}

func TestSimple_1BoolFlag_1Flag_Cmd_1Arg(context *testing.T) {
	input := "gomu -name-only -i hatchify sync mod-common"

	args := strings.Split(input, " ")
	command := simpleParse(args)

	expectedAction := "sync"

	actionTest := simply.Test("Action", context)
	actionResult := actionTest.Target(command.Action).Equals(expectedAction)
	actionTest.Validate(actionResult)

	var arg Argument
	arg.Name = "mod-common"
	arg.Value = "mod-common"
	expectedArguments := []*Argument{&arg}

	argTest := simply.Test("Arguments", context)
	argResult := argTest.Target(command.Arguments).Equals(expectedArguments)
	argTest.Validate(argResult)

	var iFlag Flag
	iFlag.Name = "-i"
	iFlag.Identifiers = []string{"-i"}
	iFlag.Value = []string{"hatchify"}
	var nameFlag Flag
	nameFlag.Name = "-name-only"
	nameFlag.Identifiers = []string{"-name-only"}
	nameFlag.Value = true
	nameFlag.Type = BOOL
	expectedFlags := map[string]*Flag{"-i": &iFlag, "-name-only": &nameFlag}

	flagTest := simply.Test("Flags", context)
	flagResult := flagTest.Target(command.Flags).Equals(expectedFlags)
	flagTest.Validate(flagResult)
}
