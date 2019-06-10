package query

import (
	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

func (f *Formatter) AsString(i *Item) string {
	return fmt.Sprintf("%v", val)
}

func (f *Formatter) AsJSON(i *Item) string {
	return
}
