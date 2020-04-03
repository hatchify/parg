package flag

// Expected test values

// Empty values
var emptyAction = ""
var emptyArguments = []*Argument{}
var emptyFlags = map[string]*Flag{}

// Actions
var syncAction = "sync"
var deployAction = "deploy"

// Args
var pargArg = Argument{
	Name:  "parg",
	Value: "parg",
}
var modcommonArg = Argument{
	Name:  "mod-common",
	Value: "mod-common",
}
var simplyArg = Argument{
	Name:  "simply",
	Value: "simply",
}

// Config Flags
var iConfigFlag = Flag{
	Name:        iFlagName,
	Identifiers: []string{iFlagName},
	Type:        DEFAULT,
}
var includeConfigFlag = Flag{
	Name:        iFlagName,
	Identifiers: []string{iFlagName, "-include"},
	Type:        STRINGS,
}

var nameOnlyConfigFlag = Flag{
	Name:        nameOnlyFlagName,
	Identifiers: []string{nameOnlyFlagName},
	Type:        BOOL,
}

var bConfigFlag = Flag{
	Name:        bFlagName,
	Identifiers: []string{bFlagName},
	Type:        DEFAULT,
}
var branchConfigFlag = Flag{
	Name:        bFlagName,
	Identifiers: []string{bFlagName, "-branch"},
	Type:        DEFAULT,
}

// Expected Flags
// -b
var bFlagName = "-b"
var bFlag = Flag{
	Name:        bFlagName,
	Identifiers: []string{bFlagName},
	Type:        DEFAULT,
	Value:       "JIRA-Ticket",
}
var branchFlag = Flag{
	Name:        bFlagName,
	Identifiers: []string{bFlagName, "-branch"},
	Type:        DEFAULT,
	Value:       "JIRA-Ticket",
}

// -i
var iFlagName = "-i"
var hatchifyIFlag = Flag{
	Name:        iFlagName,
	Identifiers: []string{iFlagName},
	Type:        DEFAULT,
	Value:       "hatchify",
}
var hatchifyvroomyIFlag = Flag{
	Name:        iFlagName,
	Identifiers: []string{iFlagName},
	Type:        STRINGS,
	Value:       []string{"hatchify", "vroomy"},
}
var hatchifyvroomytestorgIFlag = Flag{
	Name:        iFlagName,
	Identifiers: []string{iFlagName},
	Type:        STRINGS,
	Value:       []string{"hatchify", "vroomy", "test-org"},
}
var test1test2hatchifyvroomyIFlag = Flag{
	Name:        iFlagName,
	Identifiers: []string{iFlagName},
	Type:        STRINGS,
	Value:       []string{"test1", "test2", "hatchify", "vroomy"},
}

// -include
var test1test2hatchifyvroomyIncludeFlag = Flag{
	Name:        iFlagName,
	Identifiers: []string{iFlagName, "-include"},
	Type:        STRINGS,
	Value:       []string{"test1", "test2", "hatchify", "vroomy"},
}

// -name-only
var nameOnlyFlagName = "-name-only"
var nameOnlyFlag = Flag{
	Name:        "-name-only",
	Identifiers: []string{"-name-only"},
	Type:        BOOL,
	Value:       true,
}
