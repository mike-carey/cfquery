package commands

import (
	"os"
	"errors"
	"fmt"
	"sync"

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

func (w *Worker) Work() (interface{}, error) {
	logger.Info("Validating Command")
	err := w.Validate()
	if err != nil {
		return nil, err
	}

	logger.Info("Working...")
	i, err := w.Command.Run(w.Options, w.Inquistor)
	if err != nil {
		return nil, err
	}

	return i, nil
}

type WorkerFactory struct {
	Options *Options
}

func (f *WorkerFactory) NewWorker(cmd Command, cnf *cfclient.Config) (*Worker, error) {
	cli, err := cfclient.NewClient(cnf)
	if err != nil {
		return nil, err
	}

	w := &Worker{
		Command: cmd,
		Options: f.Options,
		Inquistor: query.NewInquistor(cli),
	}

	return w, nil
}

func (f *WorkerFactory) Go(cmd Command) error {
	formatter := formatter.NewFormatter(f.Options.Format)

	type result struct {
		Key string
		Value interface{}
	}

	var wg sync.WaitGroup

	size := len(f.Options.Foundations)

	wg.Add(size)

	errCh := make(chan error, size)
	resCh := make(chan result, size)

	errPool := make([]error, 0)
	resPool := make(map[string]interface{}, 0)

	for key, cnf := range f.Options.Foundations {
		go func(key string, cnf *cfclient.Config) {
			logger.Infof("Spawing a worker for %s", key)
			defer wg.Done()

			w, e := f.NewWorker(cmd, cnf)
			if e != nil {
				errCh <- e
				return
			}

			i, e := w.Work()
			if e != nil {
				errCh <- e
			}

			resCh <- result{
				Key: key,
				Value: i,
			}
		}(key, cnf)
	}

	for _, _ = range f.Options.Foundations {
		select {
		case res := <-resCh:
			resPool[res.Key] = res.Value
		case err := <-errCh:
			errPool = append(errPool, err)
		}
	}

	wg.Wait()

	if len(errPool) > 0 {
		return util.StackErrors(errPool)
	}

	buffer, err := formatter.Format(resPool)
	if err != nil {
		return err
	}

	buffer.WriteTo(os.Stdout)

	return nil
}
