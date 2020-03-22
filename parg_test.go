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
	actionTest, result := simply.Test(context, "Action")
	actionTest.Expects(command.Action).ToEqual(emptyString).AndValidate(*result)

	argTest, result := simply.Test(context, "Arguments")
	argTest.Expects(command.Arguments).ToEqual(emptyArguments).AndValidate(*result)

	flagTest, result := simply.Test(context, "Flags")
	flagTest.Expects(command.Flags).ToEqual(emptyFlags).AndValidate(*result)
}

func TestSimpleParse_Cmd_1Arg_1Flag(context *testing.T) {
	input := "gomu sync mod-common -i hatchify"

	args := strings.Split(input, " ")
	command := simpleParse(args)

	expectedAction := "sync"
	actionTest, result := simply.Test(context, "Action")
	actionTest.Expects(command.Action).ToEqual(expectedAction).AndValidate(*result)

	var arg Argument
	arg.Name = "mod-common"
	arg.Value = "mod-common"
	expectedArguments := []*Argument{&arg}
	argTest, result := simply.Test(context, "Arguments")
	argTest.Expects(command.Arguments).ToEqual(expectedArguments).AndValidate(*result)

	var flag Flag
	flag.Name = "-i"
	flag.Identifiers = []string{"-i"}
	flag.Value = []string{"hatchify"}
	expectedFlags := map[string]*Flag{"-i": &flag}

	flagTest, result := simply.Test(context, "Flags")
	flagTest.Expects(command.Flags).ToEqual(expectedFlags).AndValidate(*result)
}

/*
func TestSimple_1Flag_Cmd_1Arg(t *testing.T) {
	input := "gomu -i hatchify sync mod-common"

	args := strings.Split(input, " ")
	command := simple(args)

	test("Action").compare(command.Action, "sync",
		func(msg string) {
			t.Errorf(msg)
		})
}

func TestSimple_1BoolFlag_1Flag_Cmd_1Arg(t *testing.T) {
	input := "gomu -name-only -i hatchify sync mod-common"

	args := strings.Split(input, " ")
	command := simple(args)

	test("Action").compare(command.Action, "sync",
		func(msg string) {
			t.Errorf(msg)
		})
}
*/
