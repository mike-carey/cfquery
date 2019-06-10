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

// func (i *Items) ToJson() {
// 	enc := json.NewEncoder(os.Stdin)
// 	enc.SetIndent("", "    ")
// 	if err := enc.Encode(&a); err != nil {
// 		panic(err)
// 	}
// }
