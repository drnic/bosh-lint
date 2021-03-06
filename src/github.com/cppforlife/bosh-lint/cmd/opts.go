package cmd

import (
	boshcmd "github.com/cloudfoundry/bosh-cli/cmd"
)

type BoshOpts struct {
	VersionOpt func() error `long:"version" short:"v" description:"Show CLI version"`

	JSONOpt    bool `long:"json"     description:"Output as JSON"`
	TTYOpt     bool `long:"tty"      description:"Force TTY-like output"`
	NoColorOpt bool `long:"no-color" description:"Toggle colorized output"`

	Help HelpOpts `command:"help" description:"Show this help message"`

	LintRelease LintReleaseOpts `command:"lint-release"  description:"Lint release"`

	LintCPIConfig     LintCPIConfigOpts     `command:"lint-cpi-config"     hidden:"true" description:"Lint CPI config"`
	LintRuntimeConfig LintRuntimeConfigOpts `command:"lint-runtime-config" hidden:"true" description:"Lint runtime config"`
	LintCloudConfig   LintCloudConfigOpts   `command:"lint-cloud-config"   hidden:"true" description:"Lint cloud config"`
	LintManifest      LintManifestOpts      `command:"lint-manifest"                     description:"Lint deployment manifest"`
}

type HelpOpts struct {
	cmd
}

type LintReleaseOpts struct {
	Directory boshcmd.DirOrCWDArg `long:"dir" description:"Release directory path if not current working directory" default:"."`
	Verbose   int                 `long:"verbose" value-name:"LEVEL" description:"Show all checks [0,1,2]"`
	cmd
}

type LintCPIConfigOpts struct {
	Args    FileArgs `positional-args:"true" required:"true"`
	Verbose int      `long:"verbose" value-name:"LEVEL" description:"Show more details [0,1,2]"`
	cmd
}

type LintCloudConfigOpts struct {
	Args    FileArgs `positional-args:"true" required:"true"`
	Verbose int      `long:"verbose" value-name:"LEVEL" description:"Show more details [0,1,2]"`
	cmd
}

type LintRuntimeConfigOpts struct {
	Args    FileArgs `positional-args:"true" required:"true"`
	Verbose int      `long:"verbose" value-name:"LEVEL" description:"Show more details [0,1,2]"`
	cmd
}

type LintManifestOpts struct {
	Args    FileArgs `positional-args:"true" required:"true"`
	Verbose int      `long:"verbose" value-name:"LEVEL" description:"Show more details [0,1,2]"`
	cmd
}

type FileArgs struct {
	File boshcmd.FileBytesArg `positional-arg-name:"PATH" description:"Path to a file"`
}

// MessageOpts is used for version and help flags
type MessageOpts struct {
	Message string
}

type cmd struct{}

// Execute is necessary for each command to be goflags.Commander
func (c cmd) Execute(_ []string) error {
	panic("Unreachable")
}
