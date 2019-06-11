package query

import (
	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

type Items []Item
type ItemMap map[string]Item
type ItemGroup map[string]Items
type MappedItemMap map[string]ItemMap
type MappedItemGroup map[string]ItemGroup
