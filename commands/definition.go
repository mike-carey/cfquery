package commands

//go:generate genny -in=../generics/commands/generic-command.go -out=gen-apps-command.go -pkg commands gen "Item=cfclient.App"
//go:generate genny -in=../generics/commands/generic-command.go -out=gen-service-instances-command.go -pkg commands gen "Item=cfclient.ServiceInstance"
//go:generate genny -in=../generics/commands/generic-command.go -out=gen-spaces-command.go -pkg commands gen "Item=cfclient.Space"
//go:generate ../generics/patch.sh -- gen-*.go

//go:generate counterfeiter -o fakes/fake_command.go command.go Command

import (
	"io"

	"github.com/mike-carey/cfquery/config"
	"github.com/mike-carey/cfquery/logger"

	"github.com/jessevdk/go-flags"
)

var workerFactory *WorkerFactory

func Main(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	opts := &Options{}
	gopts := &GlobalOptions{}
	cmds := &Commands{}

	workerFactory = &WorkerFactory{Options: opts,}

	logger.Info("Parsing global options")
	args, err := flags.NewParser(gopts, flags.IgnoreUnknown).ParseArgs(args)
	if err != nil {
		return err
	}

	if gopts.Verbose == true {
		logger.SetVerbose()
	}

	f, err := config.LoadConfig(gopts.Config)
	if err != nil {
		return err
	}

	logger.Info("Parsing options")
	args, err = flags.ParseArgs(opts, args)
	if err != nil {
		return err
	}

	opts.Foundations = f

	logger.Info("Parsing commands")
	args, err = flags.ParseArgs(cmds, args)

	return err
}
