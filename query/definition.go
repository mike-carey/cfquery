package query

//**
// App
//*
//go:generate genny -in=../generics/query/generic-types.go -out=gen-app-types.go -pkg query gen "Item=cfclient.App"
//go:generate genny -in=../generics/query/generic-group-by.go -out=gen-app-group-by.go -pkg query gen "Item=cfclient.App"
//go:generate genny -in=../generics/query/generic-filter-by.go -out=gen-app-filter-by.go -pkg query gen "Item=cfclient.App"
//go:generate genny -in=../generics/query/generic-cf-service-base.go -out=gen-app-service-base.go -pkg query gen "Item=cfclient.App"
//go:generate genny -in=../generics/query/generic-cf-service-base_test.go -out=gen-app-service-base_test.go -pkg query_test gen "Item=cfclient.App"

//**
// Org
//*
//go:generate genny -in=../generics/query/generic-types.go -out=gen-org-types.go -pkg query gen "Item=cfclient.Org"
//go:generate genny -in=../generics/query/generic-group-by.go -out=gen-org-group-by.go -pkg query gen "Item=cfclient.Org"
//go:generate genny -in=../generics/query/generic-filter-by.go -out=gen-org-filter-by.go -pkg query gen "Item=cfclient.Org"
//go:generate genny -in=../generics/query/generic-cf-service-base.go -out=gen-org-service-base.go -pkg query gen "Item=cfclient.Org"
//go:generate genny -in=../generics/query/generic-cf-service-base_test.go -out=gen-org-service-base_test.go -pkg query_test gen "Item=cfclient.Org"

//**
// Space
//*
//go:generate genny -in=../generics/query/generic-types.go -out=gen-space-types.go -pkg query gen "Item=cfclient.Space"
//go:generate genny -in=../generics/query/generic-group-by.go -out=gen-space-group-by.go -pkg query gen "Item=cfclient.Space"
//go:generate genny -in=../generics/query/generic-filter-by.go -out=gen-space-filter-by.go -pkg query gen "Item=cfclient.Space"
//go:generate genny -in=../generics/query/generic-cf-service-base.go -out=gen-space-service-base.go -pkg query gen "Item=cfclient.Space"
//go:generate genny -in=../generics/query/generic-cf-service-base_test.go -out=gen-space-service-base_test.go -pkg query_test gen "Item=cfclient.Space"

//**
// Service Binding
//*
//go:generate genny -in=../generics/query/generic-types.go -out=gen-service-binding-types.go -pkg query gen "Item=cfclient.ServiceBinding"
//go:generate genny -in=../generics/query/generic-group-by.go -out=gen-service-binding-group-by.go -pkg query gen "Item=cfclient.ServiceBinding"
//go:generate genny -in=../generics/query/generic-filter-by.go -out=gen-service-binding-filter-by.go -pkg query gen "Item=cfclient.ServiceBinding"
//go:generate genny -in=../generics/query/generic-cf-service-base.go -out=gen-service-binding-service-base.go -pkg query gen "Item=cfclient.ServiceBinding"
//go:generate genny -in=../generics/query/generic-cf-service-base_test.go -out=gen-service-binding-service-base_test.go -pkg query_test gen "Item=cfclient.ServiceBinding"

//**
// Service Instance
//*
//go:generate genny -in=../generics/query/generic-types.go -out=gen-service-instance-types.go -pkg query gen "Item=cfclient.ServiceInstance"
//go:generate genny -in=../generics/query/generic-group-by.go -out=gen-service-instance-group-by.go -pkg query gen "Item=cfclient.ServiceInstance"
//go:generate genny -in=../generics/query/generic-filter-by.go -out=gen-service-instance-filter-by.go -pkg query gen "Item=cfclient.ServiceInstance"
//go:generate genny -in=../generics/query/generic-cf-service-base.go -out=gen-service-instance-service-base.go -pkg query gen "Item=cfclient.ServiceInstance"
//go:generate genny -in=../generics/query/generic-cf-service-base_test.go -out=gen-service-instance-service-base_test.go -pkg query_test gen "Item=cfclient.ServiceInstance"

//go:generate ../generics/patch.sh -- gen-*.go
//go:generate ../generics/patch.sh --inject-test-imports -- gen-*_test.go
