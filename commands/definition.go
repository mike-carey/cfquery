package commands

import (
	"io"

	"github.com/mike-carey/cfquery/config"
	"github.com/mike-carey/cfquery/logger"

	"github.com/jessevdk/go-flags"
)

//go:generate counterfeiter -o fakes/fake_command.go command.go Command

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
