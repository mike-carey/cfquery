// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package generics

import (
	"sync"

	"github.com/mike-carey/cfquery/logger"
)

func ForEachFooToBar(these []Foo, do func(Foo) (Bar, error)) ([]Bar, []error) {
	pool := make([]Bar, 0)
	errs := make([]error, 0)

	if len(these) > 0 {
		var wg sync.WaitGroup

		wg.Add(len(these))

		logger.Infof("Asynchronously running for each input")

		poolCh := make(chan Bar, len(these))
		errsCh := make(chan error, len(these))
		for _, this := range these {
			go func(this Foo) {
				defer wg.Done()

				t, e := do(this)
				if e != nil {
					errsCh <- e
				} else {
					poolCh <- t
				}
			}(this)
		}

		wg.Wait()
		logger.Infof("Collecting all outputs")

		for _ = range these {
			select {
			case this := <-poolCh:
				pool = append(pool, this)
			case err := <-errsCh:
				errs = append(errs, err)
			}
		}
	}

	return pool, errs
}
