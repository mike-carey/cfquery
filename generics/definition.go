// +build generics

package generics

//go:generate genny -in=util/generic-group-by.go -out=foobar-group-by.go -pkg generics gen "Item=Foo"
//go:generate genny -in=util/generic-filter-by.go -out=foobar-filter-by.go -pkg generics gen "Item=Foo"
