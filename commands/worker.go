package commands

import (
	"os"
	"errors"
	"fmt"

	"github.com/mike-carey/cfquery/formatter"
	"github.com/mike-carey/cfquery/logger"
	"github.com/mike-carey/cfquery/query"
	"github.com/mike-carey/cfquery/util"

	cfclient "github.com/cloudfoundry-community/go-cfclient"
)

type Worker struct {
	Command Command
	Options *Options
	Inquistor *query.Inquistor
	Formatter *formatter.Formatter
}

func (w *Worker) Validate() error {
	if w.Options.Target != "" && !util.StringSliceContains(w.Command.TargetOptions(), w.Options.Target) {
		return errors.New(fmt.Sprintf("Unknown target: %s", w.Options.Target))
	}

	if w.Options.SortBy != "" && !util.StringSliceContains(w.Command.SortByOptions(), w.Options.SortBy) {
		return errors.New(fmt.Sprintf("Unknown sort by: %s", w.Options.SortBy))
	}

	if w.Options.GroupBy != "" && !util.StringSliceContains(w.Command.GroupByOptions(), w.Options.GroupBy) {
		return errors.New(fmt.Sprintf("Unknown group by: %s", w.Options.GroupBy))
	}

	return nil
}

func (w *Worker) Work() error {
	logger.Info("Validating Command")
	err := w.Validate()
	if err != nil {
		return err
	}

	logger.Info("Working...")
	i, err := w.Command.Run(w.Options, w.Inquistor)
	if err != nil {
		return err
	}

	buffer, err := w.Formatter.Format(i)
	if err != nil {
		return err
	}

	buffer.WriteTo(os.Stdout)

	return nil
}

type WorkerFactory struct {
	Options *Options
}

func (f *WorkerFactory) NewWorker(cmd Command) (*Worker, error) {
	cnf, err := f.Options.Foundations.Get("pws")
	if err != nil {
		return nil, err
	}

	cli, err := cfclient.NewClient(cnf)
	if err != nil {
		return nil, err
	}

	w := &Worker{
		Command: cmd,
		Options: f.Options,
		Inquistor: query.NewInquistor(cli),
		Formatter: formatter.NewFormatter(f.Options.Format),
	}

	return w, nil
}
