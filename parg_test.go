package flag

import (
	"strings"
	"testing"

	"github.com/hatchify/simply"
)

var emptyAction = ""
var emptyArguments = []*Argument{}
var emptyFlags = map[string]*Flag{}

func TestSimpleParse_Empty(context *testing.T) {
	// White box input
	input := "gomo"

	// OS-Parsed input format
	args := strings.Split(input, " ")

	// Execute test with input
	command := simpleParse(args)

	expectedAction := emptyAction
	expectedArgs := emptyArguments
	expectedFlags := emptyFlags

	// Run short hand validations
	test := simply.Target(command, context, "Command")
	test.Validate(test.Equals(Command{expectedAction, expectedArgs, expectedFlags}))

	// Run expanded short hand
	action := simply.Target(command.Action, context, "Action")
	result := action.Equals(expectedAction)
	action.Validate(result)

	// Run long hand
	arg := simply.Test(context, "Arguments")
	result = arg.Target(command.Arguments).Equals(expectedArgs)
	arg.Validate(result)

	// Run expanded long hand
	flag := simply.Test(context, "Flags")
	flag.Target(command.Flags)
	result = flag.Equals(expectedFlags)
	flag.Validate(result)
}

func TestSimpleParse_Cmd_1Arg_1Flag(context *testing.T) {
	input := "gomu sync mod-common -i hatchify"

	args := strings.Split(input, " ")
	command := simpleParse(args)

	// Expected values
	expectedAction := "sync"

	var arg Argument
	arg.Name = "mod-common"
	arg.Value = "mod-common"
	expectedArgs := []*Argument{&arg}

	var flag Flag
	flag.Name = "-i"
	flag.Identifiers = []string{"-i"}
	flag.Value = []string{"hatchify"}
	expectedFlags := map[string]*Flag{"-i": &flag}

	// Run short hand validations
	test := simply.Target(command, context, "Command")
	test.Validate(test.Equals(Command{expectedAction, expectedArgs, expectedFlags}))

	// Run expanded short hand
	action := simply.Target(command.Action, context, "Action")
	result := action.Equals(expectedAction)
	action.Validate(result)

	// Run long hand
	argTest := simply.Test(context, "Arguments")
	result = argTest.Target(command.Arguments).Equals(expectedArgs)
	argTest.Validate(result)

	// Run expanded long hand
	flagTest := simply.Test(context, "Flags")
	flagTest.Target(command.Flags)
	result = flagTest.Equals(expectedFlags)
	flagTest.Validate(result)
}

func TestSimple_1Flag_Cmd_1Arg(context *testing.T) {
	input := "gomu -i hatchify sync mod-common"

	args := strings.Split(input, " ")
	command := simpleParse(args)

	expectedAction := "sync"

	var arg Argument
	arg.Name = "mod-common"
	arg.Value = "mod-common"
	expectedArgs := []*Argument{&arg}

	var flag Flag
	flag.Name = "-i"
	flag.Identifiers = []string{"-i"}
	flag.Value = []string{"hatchify"}
	expectedFlags := map[string]*Flag{"-i": &flag}

	// Run short hand validations
	test := simply.Target(command, context, "Command")
	test.Validate(test.Equals(Command{expectedAction, expectedArgs, expectedFlags}))

	// Run expanded short hand
	action := simply.Target(command.Action, context, "Action")
	result := action.Equals(expectedAction)
	action.Validate(result)

	// Run long hand
	argTest := simply.Test(context, "Arguments")
	result = argTest.Target(command.Arguments).Equals(expectedArgs)
	argTest.Validate(result)

	// Run expanded long hand
	flagTest := simply.Test(context, "Flags")
	flagTest.Target(command.Flags)
	result = flagTest.Equals(expectedFlags)
	flagTest.Validate(result)
}

func TestSimple_1BoolFlag_1Flag_Cmd_1Arg(context *testing.T) {
	input := "gomu -name-only -i hatchify sync mod-common"

	args := strings.Split(input, " ")
	command := simpleParse(args)

	expectedAction := "sync"

	var arg Argument
	arg.Name = "mod-common"
	arg.Value = "mod-common"
	expectedArgs := []*Argument{&arg}

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

	// Run short hand validations
	test := simply.Target(command, context, "Command")
	test.Validate(test.Equals(Command{expectedAction, expectedArgs, expectedFlags}))

	// Run expanded short hand
	action := simply.Target(command.Action, context, "Action")
	result := action.Equals(expectedAction)
	action.Validate(result)

	// Run long hand
	argTest := simply.Test(context, "Arguments")
	result = argTest.Target(command.Arguments).Equals(expectedArgs)
	argTest.Validate(result)

	// Run expanded long hand
	flagTest := simply.Test(context, "Flags")
	flagTest.Target(command.Flags)
	result = flagTest.Equals(expectedFlags)
	flagTest.Validate(result)
}
