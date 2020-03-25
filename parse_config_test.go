package flag

import (
	"strings"
	"testing"

	"github.com/hatchify/simply"
)

func TestConfig_Empty_Parse_Empty_Allow(context *testing.T) {
	// White box input
	input := "gomu"

	// OS-Parsed input format
	args := strings.Split(input, " ")

	// Expected results
	expectedAction := emptyAction
	expectedArgs := emptyArguments
	expectedFlags := emptyFlags
	expectedCommand := Command{
		expectedAction,
		expectedArgs,
		expectedFlags,
	}

	// Execute test with input
	parg := new()
	command, err := parg.validate(args)

	// Ensure no error was received
	test := simply.Target(err, context, "Error should not exist")
	result := test.Assert().Equals(nil)
	test.Validate(result)

	// Ensure Command was received
	test = simply.Target(command, context, "Command should exist")
	result = test.Assert().DoesNotEqual(nil)
	test.Validate(result)

	// Validate empty command struct
	test = simply.Target(command, context, "Command struct should be empty")
	result = test.Equals(expectedCommand)
	test.Validate(result)

	// Validate empty action string
	action := simply.Target(command.Action, context, "Action string should be empty")
	action.Validate(action.Equals(expectedAction))

	// Validate empty arguments slice
	arg := simply.Test(context, "Arguments slice should be empty")
	arg.Validate(arg.Target(command.Arguments).Equals(expectedArgs))

	// Validate empty flags map
	flag := simply.Test(context, "Flags map should be empty")
	flag.Target(command.Flags)
	result = flag.Equals(expectedFlags)
	flag.Validate(result)
}

func TestConfig_Cmd_Parse_Empty_Error(context *testing.T) {
	// White box input without allowed action
	input := "gomu"
	args := strings.Split(input, " ")

	expectedAction := syncAction
	expectedArgs := emptyArguments
	expectedFlags := emptyFlags
	expectedCommand := Command{
		expectedAction,
		expectedArgs,
		expectedFlags,
	}

	// Set allowed actions "sync"
	parg := new()
	parg.AddCommand(expectedCommand)
	command, err := parg.validate(args)

	test := simply.Target(err, context, "Error Should Exists")
	result := test.DoesNotEqual(nil)
	test.Validate(result)

	test = simply.Target(command, context, "Command Should Not Exist")
	result = test.Assert().Equals(nil)
	test.Validate(result)
}

func TestConfig_Cmd_Parse_Cmd_Allow(context *testing.T) {
	input := "gomu sync"
	args := strings.Split(input, " ")

	expectedAction := syncAction
	expectedArgs := emptyArguments
	expectedFlags := emptyFlags
	expectedCommand := Command{
		expectedAction,
		expectedArgs,
		expectedFlags,
	}

	// Set allowed actions "sync"
	parg := new()
	parg.AddCommand(expectedCommand)
	command, err := parg.validate(args)

	test := simply.Target(err, context, "Error should not exist")
	result := test.Assert().Equals(nil)
	test.Validate(result)

	test = simply.Target(command, context, "Command should exist")
	result = test.Assert().DoesNotEqual(nil)
	test.Validate(result)

	// Run short hand validations
	test = simply.Target(command, context, "Command")
	result = test.Equals(expectedCommand)
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

func TestConfig_Empty_Parse_Cmd_Error(context *testing.T) {
	// White box input unallowed action
	input := "gomu sync"
	args := strings.Split(input, " ")

	// Execute test with input
	parg := new()
	command, err := parg.validate(args)

	test := simply.Target(err, context, "Error Should Exists")
	result := test.DoesNotEqual(nil)
	test.Validate(result)

	test = simply.Target(command, context, "Command Should Not Exist")
	result = test.Assert().Equals(nil)
	test.Validate(result)
}

func TestConfig_Cmd_Parse_Cmd2_Error(context *testing.T) {
	// White box input without allowed action
	input := "gomu deploy"
	// OS-Parsed input format
	args := strings.Split(input, " ")

	expectedAction := syncAction
	expectedArgs := emptyArguments
	expectedFlags := emptyFlags
	expectedCommand := Command{
		expectedAction,
		expectedArgs,
		expectedFlags,
	}

	parg := new()
	// Set allowed actions "sync"
	parg.AddCommand(expectedCommand)

	// Execute test with input
	command, err := parg.validate(args)

	test := simply.Target(err, context, "Error Should Exists")
	result := test.DoesNotEqual(nil)
	test.Validate(result)

	test = simply.Target(command, context, "Command Should Not Exist")
	result = test.Assert().Equals(nil)
	test.Validate(result)
}

func TestConfigParse_1FlagAllowed(context *testing.T) {
	input := "gomu -i hatchify"
	args := strings.Split(input, " ")

	// Expected values
	expectedAction := emptyAction
	expectedArgs := emptyArguments
	expectedFlags := map[string]*Flag{
		hatchifyIFlag.Name: &hatchifyIFlag,
	}
	expectedCommand := Command{
		expectedAction,
		expectedArgs,
		expectedFlags,
	}

	parg := new()
	parg.AddGlobalFlag(hatchifyIFlag)

	command, err := parg.validate(args)

	test := simply.Target(err, context, "Error should not exist")
	result := test.Assert().Equals(nil)
	test.Validate(result)

	test = simply.Target(command, context, "Command should exist")
	result = test.Assert().DoesNotEqual(nil)
	test.Validate(result)

	// Run short hand validations
	test = simply.Target(command, context, "Command should contain iFlag")
	result = test.Equals(expectedCommand)
	test.Validate(result)

	// Run expanded short hand
	action := simply.Target(command.Action, context, "Action should be empty")
	action.Validate(action.Equals(expectedAction))

	// Run long hand
	arg := simply.Test(context, "Arguments should be empty")
	arg.Validate(arg.Target(command.Arguments).Equals(expectedArgs))

	// Run expanded long hand
	flag := simply.Test(context, "Flags should contain iFlag")
	flag.Target(command.Flags)
	result = flag.Equals(expectedFlags)
	flag.Validate(result)
}

func TestConfigParse_BoolFlagAllowed(context *testing.T) {
	input := "gomu -name-only"
	args := strings.Split(input, " ")

	// Expected values
	expectedAction := emptyAction
	expectedArgs := emptyArguments
	expectedFlags := map[string]*Flag{
		nameOnlyFlagName: &nameOnlyFlag,
	}
	expectedCommand := Command{
		expectedAction,
		expectedArgs,
		expectedFlags,
	}

	parg := new()
	parg.AddGlobalFlag(nameOnlyFlag)

	command, err := parg.validate(args)

	test := simply.Target(err, context, "Error should not exist")
	result := simply.Assert(test).Equals(nil)
	test.Validate(result)

	test = simply.Target(command, context, "Command should exist")
	result = simply.Assert(test).DoesNotEqual(nil)
	test.Validate(result)

	// Run short hand validations
	test = simply.Target(command, context, "Command should contain nameOnlyFlag")
	result = test.Equals(expectedCommand)
	test.Validate(result)

	// Run expanded short hand
	action := simply.Target(command.Action, context, "Action should be empty")
	action.Validate(action.Equals(expectedAction))

	// Run long hand
	arg := simply.Test(context, "Arguments should be empty")
	arg.Validate(arg.Target(command.Arguments).Equals(expectedArgs))

	// Run expanded long hand
	flag := simply.Test(context, "Flags should contain iFlag")
	flag.Target(command.Flags)
	result = flag.Equals(expectedFlags)
	flag.Validate(result)
}

func TestConfigParse_1FlagError(context *testing.T) {
	parg := new()
	parg.AddGlobalFlag(hatchifyIFlag)

	input := "gomu -n hatchify"

	args := strings.Split(input, " ")
	command, err := parg.validate(args)

	test := simply.Target(err, context, "Error should exist")
	result := test.DoesNotEqual(nil)
	test.Validate(result)

	test = simply.Target(command, context, "Command should not exist")
	result = simply.Assert(test).Equals(nil)
	test.Validate(result)
}

func TestConfigParse_FlagTypeError(context *testing.T) {
	parg := new()
	parg.AddGlobalFlag(hatchifyIFlag)

	input := "gomu -n hatchify"

	args := strings.Split(input, " ")
	command, err := parg.validate(args)

	test := simply.Target(err, context, "Error should exist")
	result := test.DoesNotEqual(nil)
	test.Validate(result)

	test = simply.Target(command, context, "Command should not exist")
	result = simply.Assert(test).Equals(nil)
	test.Validate(result)
}

func TestConfigParse_1Flag_Cmd(context *testing.T) {
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
		expectedAction,
		expectedArgs,
		expectedFlags,
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

func TestConfigParse_Cmd_1Flag(context *testing.T) {
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
		expectedAction,
		expectedArgs,
		expectedFlags,
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

func TestConfigParse_Cmd_1FlagArray(context *testing.T) {
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
		expectedAction,
		expectedArgs,
		expectedFlags,
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

func TestConfigParse_1Flag_Cmd_1FlagMatch(context *testing.T) {
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
		expectedAction,
		expectedArgs,
		expectedFlags,
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

func TestConfigParse_Cmd_1Arg(context *testing.T) {
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
		expectedAction,
		expectedArgs,
		expectedFlags,
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

func TestConfigParse_Cmd_2Arg(context *testing.T) {
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
		expectedAction,
		expectedArgs,
		expectedFlags,
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

func TestConfigParse_Cmd_1Arg_1Flag(context *testing.T) {
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
		expectedAction,
		expectedArgs,
		expectedFlags,
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

func TestConfig_1Flag_Cmd_1Arg(context *testing.T) {
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
		expectedAction,
		expectedArgs,
		expectedFlags,
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

func TestConfig_1BoolFlag_1Flag_Cmd_1Arg(context *testing.T) {
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
		expectedAction,
		expectedArgs,
		expectedFlags,
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

func TestConfig_1BoolFlag_1Flag_Cmd_2Arg_1_Flag_1FlagMatch(context *testing.T) {
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
		expectedAction,
		expectedArgs,
		expectedFlags,
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

func TestConfig_1BoolFlag_1Flag_Cmd_2Arg_1_Flag_1FlagArrayMatch(context *testing.T) {
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
		expectedAction,
		expectedArgs,
		expectedFlags,
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
func TestConfig_1BoolFlag_Cmd_2Arg_1_Flag_1FlagArray(context *testing.T) {
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

	// Deep comparison
	flagValTest := simply.Test(context, "FlagValues")
	flagValTest.Target(command.Flags["-i"].Value)
	result = flagValTest.Equals(iFlag.Value)
	flagValTest.Validate(result)
}
*/
