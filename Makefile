.PHONY: generate
generate: ynab-generate ;

pkg/ynab/client.go: tools/ynab_codegen_config.yaml tools/ynab_openapi3.yaml pkg/ynab/generate.go
	go generate pkg/ynab/generate.go

.PHONY: ynab-generate
ynab-generate: pkg/ynab/client.go ;
