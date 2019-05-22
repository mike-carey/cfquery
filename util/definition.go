package util

//go:generate genny -in=../generics/util/generic-group-by.go -out=app-group-by.go -pkg util gen "Item=cfclient.App"

//go:generate ../generics/patch.sh --ignore definition.go -- *.go
