package util

//go:generate genny -in=../generics/util/generic-group-by.go -out=app-group-by.go -pkg util gen "Item=cfclient.App"
//go:generate genny -in=../generics/util/generic-filter-by.go -out=app-filter-by.go -pkg util gen "Item=cfclient.App"

//go:generate genny -in=../generics/util/generic-group-by.go -out=service-instance-group-by.go -pkg util gen "Item=cfclient.ServiceInstance"
//go:generate genny -in=../generics/util/generic-filter-by.go -out=service-instance-filter-by.go -pkg util gen "Item=cfclient.ServiceInstance"

//go:generate genny -in=../generics/util/generic-group-by.go -out=space-group-by.go -pkg util gen "Item=cfclient.Space"
//go:generate genny -in=../generics/util/generic-filter-by.go -out=space-filter-by.go -pkg util gen "Item=cfclient.Space"

//go:generate genny -in=../generics/util/generic-group-by.go -out=org-group-by.go -pkg util gen "Item=cfclient.Org"
//go:generate genny -in=../generics/util/generic-filter-by.go -out=org-filter-by.go -pkg util gen "Item=cfclient.Org"

//go:generate ../generics/patch.sh -- *-by.go
