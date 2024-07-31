.PHONY: generate
generate: ynab-generate ;

pkg/ynab/client.go: tools/ynab_codegen_config.yaml tools/ynab_openapi3.yaml pkg/ynab/ynab.go
	go generate pkg/ynab/ynab.go

.PHONY: ynab-generate
ynab-generate: pkg/ynab/client.go ;
