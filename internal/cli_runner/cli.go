package clirunner

import (
	"github.com/alecthomas/kong"
)

var CLI struct {
	Validate ValidateCmd `cmd:"" help:"Help validate a json file"`
}

func Run() error {
	context := kong.Parse(&CLI,
		kong.Name("json validator"),
		kong.Description("validate json"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: true,
		}))

	return context.Run()
}
