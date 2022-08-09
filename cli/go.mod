module github.com/vmware-samples/sbom-composer-go/cli

go 1.16

require (
	github.com/spf13/cobra v1.5.0
	sbom-composer-go/parser v0.0.0-00010101000000-000000000000
// github.com/vmware-samples/sbom-composer/parser
)

replace sbom-composer-go/parser => ../parser
