package query

//go:generate genny -in=../generics/cf/generic-cf-service-base.go -out=service-instance-service-base.go -pkg query gen "CFObject=cfclient.ServiceInstance"
//go:generate genny -in=../generics/cf/generic-cf-service-base_test.go -out=service-instance-service-base_test.go -pkg query_test gen "CFObject=cfclient.ServiceInstance"

//go:generate genny -in=../generics/cf/generic-cf-service-base.go -out=service-binding-service-base.go -pkg query gen "CFObject=cfclient.ServiceBinding"
//go:generate genny -in=../generics/cf/generic-cf-service-base_test.go -out=service-binding-service-base_test.go -pkg query_test gen "CFObject=cfclient.ServiceBinding"

//go:generate genny -in=../generics/cf/generic-cf-service-base.go -out=app-service-base.go -pkg query gen "CFObject=cfclient.App"
//go:generate genny -in=../generics/cf/generic-cf-service-base_test.go -out=app-service-base_test.go -pkg query_test gen "CFObject=cfclient.App"

//go:generate genny -in=../generics/cf/generic-cf-service-base.go -out=space-service-base.go -pkg query gen "CFObject=cfclient.Space"
//go:generate genny -in=../generics/cf/generic-cf-service-base_test.go -out=space-service-base_test.go -pkg query_test gen "CFObject=cfclient.Space"

//go:generate genny -in=../generics/cf/generic-cf-service-base.go -out=org-service-base.go -pkg query gen "CFObject=cfclient.Org"
//go:generate genny -in=../generics/cf/generic-cf-service-base_test.go -out=org-service-base_test.go -pkg query_test gen "CFObject=cfclient.Org"

//go:generate ../generics/patch.sh -- *-service-base.go
//go:generate ../generics/patch.sh --inject-test-imports -- *-service-base_test.go
