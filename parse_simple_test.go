package flag

import (
	"strings"
	"testing"

	"github.com/hatchify/simply"
)

func TestSimpleParse_Empty(context *testing.T) {
	// White box input
	input := "gomu"

	// OS-Parsed input format
	args := strings.Split(input, " ")

	// Execute test with input
	command := simpleParse(args)

	expectedAction := emptyAction
	expectedArgs := emptyArguments
	expectedFlags := emptyFlags
	expectedCommand := Command{
		Action:    expectedAction,
		Arguments: expectedArgs,
		Flags:     expectedFlags,
	}

	// Run short hand validations
	test := simply.Target(command, context, "Command")
	result := test.Equals(expectedCommand)
	test.Validate(result)

	// Run expanded short hand
	action := simply.Target(command.Action, context, "Action")
	action.Validate(action.Equals(expectedAction))

	// Run long hand
	arg := simply.Test(context, "Arguments")
	arg.Validate(arg.Target(command.Arguments).Equals(expectedArgs))

	// Run expanded long hand
	flag := simply.Test(context, "Flags")
	flag.Target(command.Flags)
	result = flag.Equals(expectedFlags)
	flag.Validate(result)
}

func TestSimpleParse_1Flag(context *testing.T) {
	input := "gomu -i hatchify"

	args := strings.Split(input, " ")
	command := simpleParse(args)

	// Expected values
	expectedAction := emptyAction
	expectedArgs := emptyArguments
	expectedFlags := map[string]*Flag{
		hatchifyIFlag.Name: &hatchifyIFlag,
	}
	expectedCommand := Command{
		Action:    expectedAction,
		Arguments: expectedArgs,
		Flags:     expectedFlags,
	}

	// Run short hand validations
	test := simply.Target(command, context, "Command")
	result := test.Equals(expectedCommand)
	test.Validate(result)

	// Run expanded short hand
	action := simply.Target(command.Action, context, "Action")
	action.Validate(action.Equals(expectedAction))

	// Run long hand
	arg := simply.Test(context, "Arguments")
	arg.Validate(arg.Target(command.Arguments).Equals(expectedArgs))

	// Run expanded long hand
	flag := simply.Test(context, "Flags")
	flag.Target(command.Flags)
	result = flag.Equals(expectedFlags)
	flag.Validate(result)
}

func TestSimpleParse_Cmd(context *testing.T) {
	input := "gomu sync"

	args := strings.Split(input, " ")
	command := simpleParse(args)

	// Expected values
	expectedAction := syncAction
	expectedArgs := emptyArguments
	expectedFlags := emptyFlags
	expectedCommand := Command{
		Action:    expectedAction,
		Arguments: expectedArgs,
		Flags:     expectedFlags,
	}

	// Run short hand validations
	test := simply.Target(command, context, "Command")
	result := test.Equals(expectedCommand)
	test.Validate(result)

	// Run expanded short hand
	action := simply.Target(command.Action, context, "Action")
	action.Validate(action.Equals(expectedAction))

	// Run long hand
	arg := simply.Test(context, "Arguments")
	arg.Validate(arg.Target(command.Arguments).Equals(expectedArgs))

	// Run expanded long hand
	flag := simply.Test(context, "Flags")
	flag.Target(command.Flags)
	result = flag.Equals(expectedFlags)
	flag.Validate(result)
}

func TestSimpleParse_1Flag_Cmd(context *testing.T) {
	input := "gomu -i hatchify sync"

	args := strings.Split(input, " ")
	command := simpleParse(args)

	// Expected values
	expectedAction := syncAction
	expectedArgs := emptyArguments
	expectedFlags := map[string]*Flag{
		hatchifyIFlag.Name: &hatchifyIFlag,
	}
	expectedCommand := Command{
		Action:    expectedAction,
		Arguments: expectedArgs,
		Flags:     expectedFlags,
	}

	// Run short hand validations
	test := simply.Target(command, context, "Command")
	result := test.Equals(expectedCommand)
	test.Validate(result)

	// Run expanded short hand
	action := simply.Target(command.Action, context, "Action")
	action.Validate(action.Equals(expectedAction))

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

func TestSimpleParse_Cmd_1Flag(context *testing.T) {
	input := "gomu sync -i hatchify"

	args := strings.Split(input, " ")
	command := simpleParse(args)

	// Expected values
	expectedAction := syncAction
	expectedArgs := emptyArguments
	expectedFlags := map[string]*Flag{
		hatchifyIFlag.Name: &hatchifyIFlag,
	}
	expectedCommand := Command{
		Action:    expectedAction,
		Arguments: expectedArgs,
		Flags:     expectedFlags,
	}

	// Run short hand validations
	test := simply.Target(command, context, "Command")
	result := test.Equals(expectedCommand)
	test.Validate(result)

	// Run expanded short hand
	action := simply.Target(command.Action, context, "Action")
	action.Validate(action.Equals(expectedAction))

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

func TestSimpleParse_Cmd_1FlagArray(context *testing.T) {
	input := "gomu sync -i hatchify vroomy"

	args := strings.Split(input, " ")
	command := simpleParse(args)

	// Expected values
	expectedAction := syncAction
	expectedArgs := emptyArguments
	expectedFlags := map[string]*Flag{
		iFlagName: &hatchifyvroomyIFlag,
	}
	expectedCommand := Command{
		Action:    expectedAction,
		Arguments: expectedArgs,
		Flags:     expectedFlags,
	}

	// Run short hand validations
	test := simply.Target(command, context, "Command")
	result := test.Equals(expectedCommand)
	test.Validate(result)

	// Run expanded short hand
	action := simply.Target(command.Action, context, "Action")
	action.Validate(action.Equals(expectedAction))

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

func TestSimpleParse_1Flag_Cmd_1FlagMatch(context *testing.T) {
	input := "gomu -i hatchify sync -i vroomy"

	args := strings.Split(input, " ")
	command := simpleParse(args)

	// Expected values
	expectedAction := syncAction
	expectedArgs := emptyArguments
	expectedFlags := map[string]*Flag{
		iFlagName: &hatchifyvroomyIFlag,
	}
	expectedCommand := Command{
		Action:    expectedAction,
		Arguments: expectedArgs,
		Flags:     expectedFlags,
	}

	// Run short hand validations
	test := simply.Target(command, context, "Command")
	result := test.Equals(expectedCommand)
	test.Validate(result)

	// Run expanded short hand
	action := simply.Target(command.Action, context, "Action")
	action.Validate(action.Equals(expectedAction))

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

func TestSimpleParse_Cmd_1Arg(context *testing.T) {
	input := "gomu sync parg"

	args := strings.Split(input, " ")
	command := simpleParse(args)

	// Expected values
	expectedAction := syncAction
	expectedArgs := []*Argument{
		&pargArg,
	}
	expectedFlags := emptyFlags
	expectedCommand := Command{
		Action:    expectedAction,
		Arguments: expectedArgs,
		Flags:     expectedFlags,
	}

	// Run short hand validations
	test := simply.Target(command, context, "Command")
	result := test.Equals(expectedCommand)
	test.Validate(result)

	// Run expanded short hand
	action := simply.Target(command.Action, context, "Action")
	action.Validate(action.Equals(expectedAction))

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

func TestSimpleParse_Cmd_2Arg(context *testing.T) {
	input := "gomu sync parg simply"

	args := strings.Split(input, " ")
	command := simpleParse(args)

	// Expected values
	expectedAction := syncAction
	expectedArgs := []*Argument{
		&pargArg,
		&simplyArg,
	}
	expectedFlags := emptyFlags
	expectedCommand := Command{
		Action:    expectedAction,
		Arguments: expectedArgs,
		Flags:     expectedFlags,
	}

	// Run short hand validations
	test := simply.Target(command, context, "Command")
	result := test.Equals(expectedCommand)
	test.Validate(result)

	// Run expanded short hand
	action := simply.Target(command.Action, context, "Action")
	action.Validate(action.Equals(expectedAction))

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

func TestSimpleParse_Cmd_1Arg_1Flag(context *testing.T) {
	input := "gomu sync mod-common -i hatchify"

	args := strings.Split(input, " ")
	command := simpleParse(args)

	// Expected values
	expectedAction := syncAction
	expectedArgs := []*Argument{
		&modcommonArg,
	}
	expectedFlags := map[string]*Flag{
		hatchifyIFlag.Name: &hatchifyIFlag,
	}
	expectedCommand := Command{
		Action:    expectedAction,
		Arguments: expectedArgs,
		Flags:     expectedFlags,
	}

	// Run short hand validations
	test := simply.Target(command, context, "Command")
	result := test.Equals(expectedCommand)
	test.Validate(result)

	// Run expanded short hand
	action := simply.Target(command.Action, context, "Action")
	action.Validate(action.Equals(expectedAction))

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

	expectedAction := syncAction
	expectedArgs := []*Argument{
		&modcommonArg,
	}
	expectedFlags := map[string]*Flag{
		hatchifyIFlag.Name: &hatchifyIFlag,
	}
	expectedCommand := Command{
		Action:    expectedAction,
		Arguments: expectedArgs,
		Flags:     expectedFlags,
	}

	// Run short hand validations
	test := simply.Target(command, context, "Command")
	result := test.Equals(expectedCommand)
	test.Validate(result)

	// Run expanded short hand
	action := simply.Target(command.Action, context, "Action")
	action.Validate(action.Equals(expectedAction))

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

	expectedAction := syncAction
	expectedArgs := []*Argument{
		&modcommonArg,
	}
	expectedFlags := map[string]*Flag{
		hatchifyIFlag.Name: &hatchifyIFlag,
		nameOnlyFlagName:   &nameOnlyFlag,
	}
	expectedCommand := Command{
		Action:    expectedAction,
		Arguments: expectedArgs,
		Flags:     expectedFlags,
	}

	// Run short hand validations
	test := simply.Target(command, context, "Command")
	result := test.Equals(expectedCommand)
	test.Validate(result)

	// Run expanded short hand
	action := simply.Target(command.Action, context, "Action")
	action.Validate(action.Equals(expectedAction))

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

func TestSimple_1BoolFlag_1Flag_Cmd_2Arg_1_Flag_1FlagMatch(context *testing.T) {
	input := "gomu -name-only -i hatchify deploy mod-common simply -b JIRA-Ticket -i vroomy"

	args := strings.Split(input, " ")
	command := simpleParse(args)

	expectedAction := deployAction
	expectedArgs := []*Argument{
		&modcommonArg,
		&simplyArg,
	}
	expectedFlags := map[string]*Flag{
		iFlagName:        &hatchifyvroomyIFlag,
		nameOnlyFlagName: &nameOnlyFlag,
		bFlagName:        &bFlag,
	}
	expectedCommand := Command{
		Action:    expectedAction,
		Arguments: expectedArgs,
		Flags:     expectedFlags,
	}

	// Run short hand validations
	test := simply.Target(command, context, "Command")
	result := test.Equals(expectedCommand)
	test.Validate(result)

	// Run expanded short hand
	action := simply.Target(command.Action, context, "Action")
	action.Validate(action.Equals(expectedAction))

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

func TestSimple_1BoolFlag_1Flag_Cmd_2Arg_1_Flag_1FlagArrayMatch(context *testing.T) {
	input := "gomu -name-only -i hatchify deploy mod-common simply -b JIRA-Ticket -i vroomy test-org"

	args := strings.Split(input, " ")
	command := simpleParse(args)

	expectedAction := deployAction
	expectedArgs := []*Argument{
		&modcommonArg,
		&simplyArg,
	}
	expectedFlags := map[string]*Flag{
		iFlagName:        &hatchifyvroomytestorgIFlag,
		nameOnlyFlagName: &nameOnlyFlag,
		bFlagName:        &bFlag,
	}
	expectedCommand := Command{
		Action:    expectedAction,
		Arguments: expectedArgs,
		Flags:     expectedFlags,
	}

	// Run short hand validations
	test := simply.Target(command, context, "Command")
	result := test.Equals(expectedCommand)
	test.Validate(result)

	// Run expanded short hand
	action := simply.Target(command.Action, context, "Action")
	action.Validate(action.Equals(expectedAction))

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

// This test cannot pass with default parse rules
//   1) bool flags immediately preceding command names
//   2_ array flags before command is set
// Both justify custom config for flag parsing
/*
func TestSimple_1BoolFlag_Cmd_2Arg_1_Flag_1FlagArray(context *testing.T) {
	input := "gomu -name-only deploy mod-common simply -b JIRA-Ticket -i vroomy hatchify"

	args := strings.Split(input, " ")
	command := simpleParse(args)

	expectedAction := "deploy"

	var arg Argument
	arg.Name = "mod-common"
	arg.Value = "mod-common"
	var arg2 Argument
	arg2.Name = "simply"
	arg2.Value = "simply"
	expectedArgs := []*Argument{&arg, &arg2}

	var iFlag Flag
	iFlag.Name = "-i"
	iFlag.Identifiers = []string{"-i"}
	iFlag.Value = []string{"vroomy", "hatchify"}
	iFlag.Type = STRINGS
	var bFlag Flag
	bFlag.Name = "-b"
	bFlag.Identifiers = []string{"-b"}
	bFlag.Value = "JIRA-Ticket"
	bFlag.Type = DEFAULT
	var nameFlag Flag
	nameFlag.Name = "-name-only"
	nameFlag.Identifiers = []string{"-name-only"}
	nameFlag.Value = true
	nameFlag.Type = BOOL
	expectedFlags := map[string]*Flag{"-i": &iFlag, "-name-only": &nameFlag, "-b": &bFlag}

	// Run short hand validations
	test := simply.Target(command, context, "Command")
	test.Validate(test.Equals(Command{Action: expectedAction, Arguments: expectedArgs, expectedFlags}))

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

	// Deep comparison
	flagValTest := simply.Test(context, "FlagValues")
	flagValTest.Target(command.Flags["-i"].Value)
	result = flagValTest.Equals(iFlag.Value)
	flagValTest.Validate(result)
}
*/
